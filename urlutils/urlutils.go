package urlutils

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"jaytaylor.com/html2text"
)

func IsURL(text string) bool {
	_, err := url.ParseRequestURI(text)
	return err == nil
}

func ExtractContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println("Got content, extracting text...")
	text, err := html2text.FromString(string(body), html2text.Options{PrettyTables: true})
	if err != nil {
		return "", err
	}
	log.Println("Text extracted, processing...")
	log.Println(text)
	return text, nil
}
