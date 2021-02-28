package main

import (
	"fmt"
	"github.com/leoython/text-to-video/internal"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Speech struct {
	Folder   string
	FileName string
	Language string
}

func (speech *Speech) GenerateAudio(text string) error {

	var err error
	fileName := speech.Folder + "/" + speech.FileName + ".mp3"

	if err = internal.CreateFolderIfNotExists(speech.Folder); err != nil {
		return err
	}
	if err = speech.downloadIfNotExists(fileName, text); err != nil {
		return err
	}

	return nil
}

func (speech *Speech) downloadIfNotExists(fileName string, text string) error {
	f, err := os.Open(fileName)
	if err != nil {
		response, err := http.Get(fmt.Sprintf("http://translate.google.com/translate_tts?ie=UTF-8&total=1&idx=0&textlen=32&client=tw-ob&q=%s&tl=%s", url.QueryEscape(text), speech.Language))
		if err != nil {
			return err
		}
		defer response.Body.Close()

		output, err := os.Create(fileName)
		if err != nil {
			return err
		}

		_, err = io.Copy(output, response.Body)
		return err
	}

	_ = f.Close()
	return nil
}

