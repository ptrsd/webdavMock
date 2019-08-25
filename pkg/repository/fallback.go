package repository

import (
	"io/ioutil"
	"os"
	"strings"
)

type FallbackConfig struct {
	Path string `yaml:path`
}

type Fallback struct {
	Config        FallbackConfig
	FallbackStore map[string][]byte
}

func CreateFallback(config FallbackConfig) (fallback Fallback, err error) {
	var fallbackFiles []os.FileInfo
	fallback.FallbackStore = map[string][]byte{}

	if fallbackFiles, err = ioutil.ReadDir(config.Path); err != nil {
		return Fallback{}, err
	}

	for _, fallbackFile := range fallbackFiles {
		if !fallbackFile.IsDir() {
			var fileBytes []byte

			splittedFileName := strings.Split(fallbackFile.Name(), ".")
			if fileBytes, err = ioutil.ReadFile(config.Path + "/" + fallbackFile.Name()); err != nil {
				return Fallback{}, err
			}

			fallback.FallbackStore[splittedFileName[1]] = fileBytes
		}
	}

	return fallback, nil
}

func (f Fallback) Find(path string) (fileContent []byte, _ error) {
	var exists bool

	if fileContent, exists = f.FallbackStore[path]; !exists {
		return nil, NotFoundErr{path}
	}

	return fileContent, nil
}
