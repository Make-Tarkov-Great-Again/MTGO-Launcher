package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	logging "mtgolauncher/backend/Logging"
	profile "mtgolauncher/backend/Profile"
	storage "mtgolauncher/backend/Storage"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ConfigRunT struct {
}

func NewConfig() *ConfigRunT {
	return &ConfigRunT{}
}

var settingsMap = map[string]*string{
	"AkiPath":     &runtimeConfig.UserSettings.Server.AkiServerPath,
	"AkiAddress":  &runtimeConfig.UserSettings.Server.AkiServerAddress,
	"MtgaPath":    &runtimeConfig.UserSettings.Server.MtgaServerPath,
	"MtgaAddress": &runtimeConfig.UserSettings.Server.MtgaServerAddress,
	"ClientPath":  &runtimeConfig.UserSettings.ClientPath,
	"Language":    &runtimeConfig.UserSettings.Language,
	"Theme":       &runtimeConfig.UserSettings.Theme,
}

type Config struct {
	AppInfo struct {
		Version string `json:"version"`
		Build   string `json:"build"`
	} `json:"AppInfo"`
	UserSettings struct {
		Protections bool   `json:"protections"`
		ClientPath  string `json:"clientPath"`
		Server      struct {
			AkiServerPath     string `json:"akiServerPath"`
			MtgaServerPath    string `json:"mtgaServerPath"`
			MtgaServerAddress string `json:"mtgaServerAddress"`
			AkiServerAddress  string `json:"akiServerAddress"`
		} `json:"server"`
		Language    string `json:"language"`
		Theme       string `json:"theme"`
		LastProfile string `json:"lastProfile"`
	} `json:"UserSettings"`
}

type PackageJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Licence     string `json:"license"`
	AkiVersion  string `json:"akiVersion"`
	MtgaVersion string `json:"mtgaVersion"`
}

var count int

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

var runtimeConfig Config

func Init() {
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		logging.Error("Failed to get app data directory: %v", err)
	}

	runtimeDir := path.Join(appDir, "Runtime")
	if _, err := os.Stat(runtimeDir); os.IsNotExist(err) {
		if err := os.MkdirAll(runtimeDir, 0755); err != nil {
			logging.Error("Failed to create Runtime directory: %v", err)
		}
	}

	configPath := path.Join(runtimeDir, "config.json")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := Config{
			AppInfo: struct {
				Version string `json:"version"`
				Build   string `json:"build"`
			}{
				Version: "pre=0.5.1",
				Build:   "1",
			},
			UserSettings: struct {
				Protections bool   `json:"protections"`
				ClientPath  string `json:"clientPath"`
				Server      struct {
					AkiServerPath     string `json:"akiServerPath"`
					MtgaServerPath    string `json:"mtgaServerPath"`
					MtgaServerAddress string `json:"mtgaServerAddress"`
					AkiServerAddress  string `json:"akiServerAddress"`
				} `json:"server"`
				Language    string `json:"language"`
				Theme       string `json:"theme"`
				LastProfile string `json:"lastProfile"`
			}{
				Protections: true,
				Server: struct {
					AkiServerPath     string `json:"akiServerPath"`
					MtgaServerPath    string `json:"mtgaServerPath"`
					MtgaServerAddress string `json:"mtgaServerAddress"`
					AkiServerAddress  string `json:"akiServerAddress"`
				}{
					AkiServerPath:     "",
					MtgaServerPath:    "",
					MtgaServerAddress: "127.0.0.1:8080",
					AkiServerAddress:  "127.0.0.1:6969",
				},
				Language:    "enUS",
				Theme:       "defaultDark",
				LastProfile: "",
			},
		}

		defaultJSON, err := json.MarshalIndent(defaultConfig, "", "  ")
		if err != nil {
			logging.Error("Failed to marshal default config to JSON: %v", err)
		}

		if err := os.WriteFile(configPath, defaultJSON, 0644); err != nil {
			logging.Error("Failed to create config file: %v", err)
		}
	}

	configData, err := os.ReadFile(configPath)
	if err != nil {
		count++
	}

	if err := json.Unmarshal(configData, &runtimeConfig); err != nil {
		count++
	}

	if runtimeConfig.UserSettings.Server.AkiServerPath == "" {
		count++
	}

	if runtimeConfig.UserSettings.Server.MtgaServerPath == "" {
		count++
	}

	if runtimeConfig.UserSettings.ClientPath == "" {
		count++
	}

	if runtimeConfig.UserSettings.Server.MtgaServerAddress == "" {
		count++
	}

	if runtimeConfig.UserSettings.Server.AkiServerAddress == "" {
		count++
	}
	if count > 0 {
		logging.Error(fmt.Sprintf("%d nesscary config values are left unfulfilled! Please edit your config!", count))
	}

}

func (c ConfigRunT) GetRuntimeConfig() (Config, error) {
	appDir, err := storage.GetAppDataDir()

	configPath := path.Join(appDir, "Runtime/config.json")
	if err != nil {
		logging.Error("Failed to get runtime config")
		return Config{}, err
	}

	configData, err := os.ReadFile(configPath)
	if err != nil {
		logging.Error("Failed to read runtime config")
		return Config{}, err
	}

	var runtimeConfig Config
	if err := json.Unmarshal(configData, &runtimeConfig); err != nil {
		logging.Error("Failed to unmarshal runtime config")
		return Config{}, err
	}

	return runtimeConfig, nil
}

func GetAKIServerInfo() (string, string, string, string, error) {
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		return "", "", "", "", err
	}

	configPath := path.Join(appDir, "Runtime/config.json")

	configData, err := os.ReadFile(configPath)
	if err != nil {
		return "", "", "", "", err
	}

	var runtimeConfig Config
	if err := json.Unmarshal(configData, &runtimeConfig); err != nil {
		return "", "", "", "", err
	}

	// Check if the AKI Server path is not empty
	if runtimeConfig.UserSettings.Server.AkiServerPath == "" {
		return "", "", "", "", errors.New("AKI Server path is empty")
	}

	return runtimeConfig.UserSettings.Server.AkiServerPath, runtimeConfig.UserSettings.Server.AkiServerAddress, runtimeConfig.UserSettings.LastProfile, runtimeConfig.UserSettings.ClientPath, nil
}

func (c ConfigRunT) CopyProfiles() error {
	appDir, err := storage.GetAppDataDir()
	sourcePath := path.Join(runtimeConfig.UserSettings.Server.AkiServerPath, "user/profiles")
	destinationPath := path.Join(appDir, "/profiles/AKI")
	// Ensure the destination directory exists
	if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil {
		return err
	}

	// List profile files in the source directory
	profileFiles, err := os.ReadDir(sourcePath)
	if err != nil {
		return err
	}

	// Iterate over profile files and copy each one
	for _, profileFile := range profileFiles {
		if profileFile.IsDir() {
			continue
		}

		srcFilePath := filepath.Join(sourcePath, profileFile.Name())
		destFilePath := filepath.Join(destinationPath, profileFile.Name())

		srcFile, err := os.Open(srcFilePath)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(destFilePath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, srcFile); err != nil {
			return err
		}

		// Parse and save the profile after copying
		if err := profile.ParseAndSaveProfile(destFilePath); err != nil {
			return err
		}
	}

	return nil
}

// Returns Client path, mtga server path, mtga server address, and last profile.
func GetMTGAServerInfo() (string, string, string, string, error) {
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		logging.Error("Failed to get runtime config")
		return "", "", "", "", err
	}

	configPath := path.Join(appDir, "Runtime/config.json")

	configData, err := os.ReadFile(configPath)
	if err != nil {
		logging.Error("Failed to read runtime config")
		return "", "", "", "", err
	}

	var runtimeConfig Config
	if err := json.Unmarshal(configData, &runtimeConfig); err != nil {
		logging.Error("Failed to unmarshal runtime config")
		return "", "", "", "", err
	}

	return runtimeConfig.UserSettings.Server.MtgaServerPath, runtimeConfig.UserSettings.Server.MtgaServerAddress, runtimeConfig.UserSettings.LastProfile, runtimeConfig.UserSettings.ClientPath, nil
}

func GetTheme() (string, error) {
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		logging.Error("Failed to get runtime config")
		return "", err
	}

	configPath := path.Join(appDir, "Runtime/config.json")

	configData, err := os.ReadFile(configPath)
	if err != nil {
		logging.Error("Failed to read runtime config")
		return "", err
	}

	var runtimeConfig Config
	if err := json.Unmarshal(configData, &runtimeConfig); err != nil {
		logging.Error("Failed to unmarshal runtime config")
		return "", err
	}

	return runtimeConfig.UserSettings.Theme, nil
}

func GetLanguage() (string, error) {
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		logging.Error("Failed to get runtime config")
		return "", err
	}

	configPath := path.Join(appDir, "Runtime/config.json")

	configData, err := os.ReadFile(configPath)
	if err != nil {
		logging.Error("Failed to read runtime config")
		return "", err
	}

	var runtimeConfig Config
	if err := json.Unmarshal(configData, &runtimeConfig); err != nil {
		logging.Error("Failed to unmarshal runtime config")
		return "", err
	}

	return runtimeConfig.UserSettings.Language, nil
}

func (c ConfigRunT) SetConfigVariable(key, value string) error {
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		logging.Error("Failed to get runtime config directory")
		return err
	}

	configPath := path.Join(appDir, "Runtime/config.json")

	configData, err := os.ReadFile(configPath)
	if err != nil {
		logging.Error("Failed to read runtime config")
		return err
	}

	var runtimeConfig Config
	if err := json.Unmarshal(configData, &runtimeConfig); err != nil {
		logging.Error("Failed to unmarshal runtime config")
		return err
	}

	//if strPtr, ok := settingsMap[key]; ok {
	//	*strPtr = value
	//} else {
	//	fmt.Printf("Unsupported configuration key: %s\n", key)
	//	return fmt.Errorf("Unsupported configuration Key %s", key)
	//}
	// Doesnt like to set for some reason...

	switch key {
	case "AkiPath":
		runtimeConfig.UserSettings.Server.AkiServerPath = value
	case "AkiAddress":
		runtimeConfig.UserSettings.Server.AkiServerAddress = value
	case "MtgaPath":
		runtimeConfig.UserSettings.Server.MtgaServerPath = value
	case "MtgaAddress":
		runtimeConfig.UserSettings.Server.MtgaServerAddress = value
	case "ClientPath":
		runtimeConfig.UserSettings.ClientPath = value
	case "Language":
		runtimeConfig.UserSettings.Language = value
	case "Theme":
		runtimeConfig.UserSettings.Theme = value
	default:
		return fmt.Errorf("Unsupported configuration key: %s", key)
	}

	// Write that hoe
	updatedConfigData, err := json.MarshalIndent(runtimeConfig, "", "    ")
	if err != nil {
		logging.Error("Failed to marshal updated runtime config")
		return err

	}
	defer NewConfig().CopyProfiles()

	if err := os.WriteFile(configPath, updatedConfigData, 0644); err != nil {
		logging.Error("Failed to write updated runtime config")
		return err
	}

	return nil
}

func (c ConfigRunT) GetConfigVariable(key string) (string, error) {
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		logging.Error("Failed to get runtime config directory")
		return "", err
	}

	configPath := path.Join(appDir, "Runtime/config.json")

	configData, err := os.ReadFile(configPath)
	if err != nil {
		logging.Error("Failed to read runtime config")
		return "", err
	}

	var runtimeConfig Config
	if err := json.Unmarshal(configData, &runtimeConfig); err != nil {
		logging.Error("Failed to unmarshal runtime config")
		return "", err
	}

	switch key {
	case "AkiPath":
		return runtimeConfig.UserSettings.Server.AkiServerPath, nil
	case "AkiAddress":
		return runtimeConfig.UserSettings.Server.AkiServerAddress, nil
	case "MtgaPath":
		return runtimeConfig.UserSettings.Server.MtgaServerPath, nil
	case "MtgaAddress":
		return runtimeConfig.UserSettings.Server.MtgaServerAddress, nil
	case "ClientPath":
		return runtimeConfig.UserSettings.ClientPath, nil
	case "Language":
		return runtimeConfig.UserSettings.Language, nil
	case "Theme":
		return runtimeConfig.UserSettings.Theme, nil
	default:
		logging.Error(fmt.Sprintf("Unsupported configuration key: %s", key))
		return "", fmt.Errorf("Unsupported configuration key: %s", key)
	}
}
