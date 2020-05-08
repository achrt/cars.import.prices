package cloud_crutch

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type FileData struct {
	Data io.ReadCloser
}

type UrlHandler struct {
}

func GetFileHttp(url string) (*FileData, error) {
	respFile, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	if respFile.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Answer from %s not OK: [%d]", url, respFile.StatusCode))
	}
	return &FileData{respFile.Body}, nil

}

func (uh *UrlHandler) GetReaderFromUrl(path string) (io.Reader, error) {
	fd, err := GetFileHttp(path)

	if err != nil {
		return nil, err
	}
	return fd.Data, nil
}
