package futils

import (
	"io"
	"net/http"
)

func (*Service) GetFileFromURL(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	fileContent, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}
