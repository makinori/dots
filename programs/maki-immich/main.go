package main

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
)

func selectAlbumKdialog() (Album, error) {
	albums, err := getAlbums()
	if err != nil {
		return Album{}, err
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
		return Album{}, err
	}

	albumId := strings.TrimSpace(string(output))

	for _, album := range albums {
		if album.Id == albumId {
			return album, nil
		}
	}

	return Album{}, errors.New("failed to match album id")
}

/*
func selectAlbumFyne() (Album, error) {
	albums, err := getAlbums()
	if err != nil {
		return Album{}, err
	}

	var albumNames []string

	for _, album := range albums {
		albumNames = append(albumNames, album.AlbumName)
	}

	i := radioListDialog(albumNames)
	if i < 0 {
		return Album{}, errors.New("no album selected")
	}

	album := albums[i]

	return album, nil
}
*/

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

	// album, err := selectAlbumFyne()
	album, err := selectAlbumKdialog()

	if err != nil {
		// didnt select
		if err.Error() == "exit status 1" || err.Error() == "no album selected" {
			os.Exit(0)
		}

		dialog(err.Error(), true)
		os.Exit(1)
	}

	// upload files

	var completed []string
	var failed []string

	for _, filePath := range filePaths {
		err := uploadFile(filePath, album.Id)
		filename := path.Base(filePath)

		if err == ErrDuplicate {
			completed = append(completed, filename+" (duplicate)")
		} else if err == nil {
			completed = append(completed, filename)
		} else {
			failed = append(failed, filename)
		}
	}

	finalMsg := fmt.Sprintf("Added %d to: %s", len(completed), album.AlbumName)

	if len(completed) > 0 {
		finalMsg += "\n" + strings.Join(completed, "\n")
	}

	if len(failed) > 0 {
		finalMsg += "\n\nFailed:\n" + strings.Join(failed, "\n")
	}

	dialog(finalMsg, false)
}
