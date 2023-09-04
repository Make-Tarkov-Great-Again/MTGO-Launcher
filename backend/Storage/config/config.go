package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PackageJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Licence     string `json:"license"`
	AkiVersion  string `json:"akiVersion"`
	MtgaVersion string `json:"mtgaVersion"`
}

func FindConfigFiles(dirPath string) (map[string]string, error) {
	configFiles := make(map[string]string)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			subDir := filepath.Join(dirPath, file.Name())
			subConfigFiles, err := FindConfigFiles(subDir)
			if err != nil {
				return nil, err
			}

			for key, value := range subConfigFiles {
				configFiles[key] = value
			}
		} else if strings.Contains(strings.ToLower(file.Name()), "config") {
			fmt.Printf("Found config %s \n", file.Name())
			configFiles[file.Name()] = filepath.Join(dirPath, file.Name())
		}
	}

	return configFiles, nil
}
