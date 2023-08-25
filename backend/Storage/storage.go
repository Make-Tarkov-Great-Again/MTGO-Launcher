package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const appSubdir = "MT-GO"

func getAppDataDir() (string, error) {
	appDataDir := os.Getenv("APPDATA")
	if appDataDir == "" {
		return "", fmt.Errorf("APPDATA environment variable not set")
	}
	return filepath.Join(appDataDir, appSubdir), nil
}

func createFolderIfNotExists(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return os.MkdirAll(folderPath, 0755)
	}
	return nil
}

func writeDataToFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func StoreData(subdir, filename string, data []byte) error {
	appDataDir, err := getAppDataDir()
	if err != nil {
		return err
	}

	subdirPath := filepath.Join(appDataDir, subdir)
	err = createFolderIfNotExists(subdirPath)
	if err != nil {
		return err
	}

	filePath := filepath.Join(subdirPath, filename)
	return writeDataToFile(filePath, data)
}

func InitializeAppDataDir() error {
	appDataDir, err := getAppDataDir()
	if err != nil {
		return err
	}

	return createFolderIfNotExists(appDataDir)
}

type dataType int

const (
	profiles dataType = iota
	modLists
	mods
	configs
	logs
	misc
)

func (dt dataType) folderName(identifier string) string {
	switch dt {
	case profiles, modLists, mods, configs, logs, misc:
		return identifier
	default:
		return ""
	}
}

func storeDataWithType(dataType dataType, identifier, filename string, data interface{}) error {
	appDataDir, err := getAppDataDir()
	if err != nil {
		return err
	}

	subdir := filepath.Join(appDataDir, dataType.folderName(identifier))
	err = createFolderIfNotExists(subdir)
	if err != nil {
		return err
	}

	filePath := filepath.Join(subdir, filename)

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return writeDataToFile(filePath, jsonData)
	}

	return nil
}

type ProfileData struct {
	ID   string `json:"_id"`
	Info struct {
		Nickname string `json:"Nickname"`
		Level    int    `json:"Level"`
		Side     string `json:"Side"`
	} `json:"Info"`
}

func StoreProfileData(filePath string) error {
	appDataDir, err := getAppDataDir()
	if err != nil {
		return err
	}

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var profileData ProfileData
	err = json.Unmarshal(jsonData, &profileData)
	if err != nil {
		return err
	}

	targetDirectory := profileData.Info.Nickname

	profilesDir := filepath.Join(appDataDir, "profiles")
	err = createFolderIfNotExists(profilesDir)
	if err != nil {
		return err
	}

	targetDir := filepath.Join(profilesDir, profiles.folderName(targetDirectory))
	err = createFolderIfNotExists(targetDir)
	if err != nil {
		return err
	}

	profileDataJSON, err := json.Marshal(profileData)
	if err != nil {
		return err
	}

	profileDetailsPath := filepath.Join(targetDir, "profile-details.json")
	err = writeDataToFile(profileDetailsPath, profileDataJSON)
	if err != nil {
		return err
	}

	newFilePath := filepath.Join(targetDir, "character.json")
	err = copyFile(filePath, newFilePath)
	if err != nil {
		return err
	}

	return nil
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}
