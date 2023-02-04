package usecase

import (
	"compress/gzip"
	"errors"
	"io"
	"strings"
)

func (uc UseCase) UseGzip(body io.Reader, contentType string) (data []byte, err error) {
	if strings.Contains(contentType, "gzip") {
		data, err = uc.decompressGzip(body)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = io.ReadAll(body)
		if err != nil {
			return nil, err
		}
	}

	if len(string(data)) < 3 {
		return nil, errors.New("wrong data")
	}

	return data, nil
}

func (uc UseCase) decompressGzip(body io.Reader) ([]byte, error) {
	gz, err := gzip.NewReader(body)
	if err != nil {
		return nil, err
	}

	defer gz.Close()

	data, err := io.ReadAll(gz)
	if err != nil {
		return nil, err
	}

	return data, nil
}
