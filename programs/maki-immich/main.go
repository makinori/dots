package main

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

var (
	//go:embed _server.txt
	IMMICH_SERVER string
	//go:embed _key.txt
	IMMICH_API_KEY string

	ErrDuplicate = errors.New("duplicate asset")
)

type Album struct {
	AlbumName         string    `json:"albumName"`
	Id                string    `json:"id"`
	AssetCount        int       `json:"assetCount"`
	LastModifiedAsset time.Time `json:"lastModifiedAssetTimestamp"`
}

type Action struct {
	Id     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

func getAlbums() ([]Album, error) {
	req, err := http.NewRequest("GET", IMMICH_SERVER+"/api/albums", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-api-key", IMMICH_API_KEY)

	res, err := new(http.Client).Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var albums []Album

	err = json.Unmarshal(bytes, &albums)
	if err != nil {
		return nil, err
	}

	slices.SortFunc(albums, func(a, b Album) int {
		return b.LastModifiedAsset.Compare(a.LastModifiedAsset)
	})

	return albums, nil
}

func selectAlbum() (string, string, error) {
	albums, err := getAlbums()
	if err != nil {
		return "", "", err
	}

	kdialogArgs := []string{
		"--radiolist", "Select an album to upload to",
	}

	for _, album := range albums {
		kdialogArgs = append(kdialogArgs, album.Id, album.AlbumName, "off")
	}

	cmd := exec.Command("kdialog", kdialogArgs...)

	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	albumId := strings.TrimSpace(string(output))

	for _, album := range albums {
		if album.Id == albumId {
			return albumId, album.AlbumName, nil
		}
	}

	return albumId, "", nil
}

/*
// doesnt handle all cases. this isn't very reliable

func stripExifDate(data []byte) ([]byte, error) {
	// find more using: exiftool -G -a -s -time:all image.jpg

	toStrip := []string{
		"AllDates",
		"CreateDate",
		"DateCreated",
		// "DateTimeCreated", // not writeable
		"DateTimeDigitized",
		"DateTimeOriginal",
		"GPSDateStamp",
		"GPSDateTime",
		"GPSTimeStamp",
		"ModifyDate",
		"TimeCreated",
		"DigitalCreationDate",
		"DigitalCreationTime",
		// "DigitalCreationDateTime", // not writeable
	}

	var args []string

	for _, tag := range toStrip {
		args = append(args, "-"+tag+"=")
	}

	args = append(args, "-")

	output := new(bytes.Buffer)

	cmd := exec.Command("exiftool", args...)
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stdout = output
	cmd.Stderr = io.Discard

	err := cmd.Run()

	if err != nil {
		return []byte{}, err
	}

	return output.Bytes(), nil
}
*/

func getRandom() (string, error) {
	bytes := make([]byte, 8)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func uploadAsset(data []byte, filename string, dateStr string) (string, error) {

	mpBuf := new(bytes.Buffer)
	mp := multipart.NewWriter(mpBuf)

	mpFile, err := mp.CreateFormFile("assetData", filename)
	if err != nil {
		return "", err
	}

	mpFile.Write(data)

	deviceAssetId, err := getRandom()
	if err != nil {
		return "", err
	}

	mp.WriteField("deviceAssetId", deviceAssetId)
	mp.WriteField("deviceId", "GO") // WEB

	// doesn't actually set the date
	mp.WriteField("fileCreatedAt", dateStr)
	mp.WriteField("fileModifiedAt", dateStr)

	mp.Close()

	req, err := http.NewRequest("POST", IMMICH_SERVER+"/api/assets", mpBuf)

	if err != nil {
		return "", err
	}

	req.Header.Add("x-api-key", IMMICH_API_KEY)
	req.Header.Add("Content-Type", mp.FormDataContentType())

	res, err := new(http.Client).Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var action Action

	err = json.Unmarshal(bytes, &action)
	if err != nil {
		return "", err
	}

	if action.Error != "" {
		return "", errors.New(action.Error)
	}

	return action.Id, nil
}

func updateAssetDate(assetId string, dateStr string) error {
	data, err := json.Marshal(map[string]any{
		"ids":              []string{assetId},
		"dateTimeOriginal": dateStr,
	})

	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	buffer.Write(data)

	req, err := http.NewRequest(
		"PUT", IMMICH_SERVER+"/api/assets", buffer,
	)

	if err != nil {
		return err
	}

	req.Header.Add("x-api-key", IMMICH_API_KEY)
	req.Header.Add("Content-Type", "application/json")

	res, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		return fmt.Errorf("status code: %d", res.StatusCode)
	}

	return nil
}

func addToAlbum(albumId string, assetId string) error {
	data, err := json.Marshal(map[string][]string{
		"ids": {assetId},
	})

	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	buffer.Write(data)

	req, err := http.NewRequest(
		"PUT", IMMICH_SERVER+"/api/albums/"+albumId+"/assets", buffer,
	)

	if err != nil {
		return err
	}

	req.Header.Add("x-api-key", IMMICH_API_KEY)
	req.Header.Add("Content-Type", "application/json")

	res, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var actions []Action

	err = json.Unmarshal(bytes, &actions)
	if err != nil {
		return err
	}

	if len(actions) == 0 {
		return errors.New("bad response")
	}

	action := actions[0]

	if action.Error == "duplicate" {
		return ErrDuplicate
	}

	if action.Error != "" {
		return errors.New(action.Error)
	}

	return nil
}

func uploadFile(pathToFile string, albumId string) error {
	fmt.Println(pathToFile)

	fileData, err := os.ReadFile(pathToFile)
	if err != nil {
		return errors.New("failed to read file")
	}

	// try to strip exif
	// ignore error
	// {
	// 	strippedData, err := stripExifDate(fileData)
	// 	if err == nil {
	// 		fileData = strippedData
	// 	}
	// }

	fileDateStr := time.Now().Format(time.RFC3339Nano)

	// upload file which also handles deduplication

	assetId, err := uploadAsset(fileData, filepath.Base(pathToFile), fileDateStr)
	if err != nil {
		return err
	}

	// update file date. ignore error i suppose

	updateAssetDate(assetId, fileDateStr)

	// add to album

	err = addToAlbum(albumId, assetId)

	if err != nil {
		return err
	}

	return nil
}

func dialog(message string, isError bool) {
	var args []string

	if isError {
		args = []string{"--error", message}
	} else {
		args = []string{"--msgbox", message}
	}

	cmd := exec.Command("kdialog", args...)
	cmd.Run()
}

func main() {
	IMMICH_SERVER = strings.TrimSpace(IMMICH_SERVER)
	IMMICH_API_KEY = strings.TrimSpace(IMMICH_API_KEY)

	// get files

	if len(os.Args) <= 1 {
		dialog("Please select one or more files", true)
		os.Exit(1)
	}

	filePaths := os.Args[1:]

	_, usingNautilus := os.LookupEnv("NAUTILUS")
	if usingNautilus {
		fileUrlPaths := strings.Split(filePaths[0], " ")
		filePaths = []string{}

		for _, fileUrlStr := range fileUrlPaths {
			fileUrl, _ := url.Parse(fileUrlStr)
			filePaths = append([]string{fileUrl.Path}, filePaths...)
		}
	}

	// get album id

	albumId, albumName, err := selectAlbum()

	if err != nil {
		// didnt select
		if err.Error() == "exit status 1" {
			os.Exit(0)
		}

		dialog(err.Error(), true)
		os.Exit(1)
	}

	// upload files

	var completed []string
	var failed []string

	for _, filePath := range filePaths {
		err := uploadFile(filePath, albumId)
		filename := path.Base(filePath)

		if err == ErrDuplicate {
			completed = append(completed, filename+" (duplicate)")
		} else if err == nil {
			completed = append(completed, filename)
		} else {
			failed = append(failed, filename)
		}
	}

	finalMsg := fmt.Sprintf("Added %d to: %s", len(completed), albumName)

	if len(completed) > 0 {
		finalMsg += "\n" + strings.Join(completed, "\n")
	}

	if len(failed) > 0 {
		finalMsg += "\n\nFailed:\n" + strings.Join(failed, "\n")
	}

	dialog(finalMsg, false)
}
