package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Chapter struct {
	Seconds int
	Title   string
}

var timestampRegExp = regexp.MustCompile("^([0-9]+):([0-9]+) (.+)$")

func parseTimestamps(filePath string) ([]Chapter, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return []Chapter{}, err
	}

	linesIter := strings.SplitSeq(string(bytes), "\n")

	var chapters []Chapter

	for line := range linesIter {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		matches := timestampRegExp.FindStringSubmatch(line)
		if matches == nil {
			return []Chapter{}, errors.New("failed to find matches for: " + line)
		}

		minutes, _ := strconv.Atoi(matches[1])
		seconds, _ := strconv.Atoi(matches[2])
		seconds += minutes * 60

		title := strings.TrimSpace(matches[3])

		// remove dash if it starts with a dash
		if strings.HasPrefix(title, "-") {
			title = strings.TrimSpace(title[1:])
		}

		chapters = append(chapters, Chapter{
			Seconds: seconds,
			Title:   title,
		})
	}

	return chapters, nil
}

func makeFfmpegChapters(chapters []Chapter, title string) string {
	lines := []string{
		";FFMETADATA1",
	}

	title = strings.TrimSpace(title)
	if title != "" {
		lines = append(lines, "title="+title)
	}

	lines = append(lines, "")

	for _, chapter := range chapters {
		lines = append(lines,
			"[CHAPTER]",
			"TIMEBASE=1/10",
			fmt.Sprintf("START=%d", chapter.Seconds*10),
			"END=",
			fmt.Sprintf("title=%s", chapter.Title),
			"",
		)
	}

	return strings.TrimSpace(strings.Join(lines, "\n"))
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(strings.TrimSpace(`
usage: <txt in> <optional title> > metadata.txt

txt should be in format:
0:00 first song
1:24 - can have dashes
...

then run:
ffmpeg -i input.ogg -i metadata.txt -map_metadata 1 -c copy output.ogg
`))
		fmt.Println("")
		os.Exit(0)
	}

	timestampsPath := os.Args[1]

	title := ""
	if len(os.Args) > 2 {
		title = os.Args[2]
	}

	chapters, err := parseTimestamps(timestampsPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(
		makeFfmpegChapters(chapters, title),
	)
}
