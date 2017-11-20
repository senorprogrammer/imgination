package tui

/*
* Code graciously stolen from https://github.com/olivere/iterm2-imagetools/blob/master/cmd/imgcat/imgcat.go
 */

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jroimartin/gocui"
)

func InlineImage(view *gocui.View, path string) string {
	if path == "" {
		return "?"
	}

	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading image %s. Error: %s", path, err.Error())
	}

	width := 120
	height := 80
	b64FileName := base64.StdEncoding.EncodeToString([]byte(path))
	b64FileContents := base64.StdEncoding.EncodeToString(body)

	// str := fmt.Sprintf("\\033]1337;File=name=%s;inline=1;width=%d;height=%d;preserveAspectRatio=0:%s\a\n", b64FileName, width, height, b64FileContents)
	str := fmt.Sprintf("\x1b]1337;File=name=%s;inline=1;width=%d;height=%d;preserveAspectRatio=1:%s\a\n", b64FileName, width, height, b64FileContents)
	return str
}
