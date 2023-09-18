package config

import (
	"encoding/json"
	"fmt"
	logging "mtgolauncher/backend/Logging"
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
	// Initialize the app directory
	appDir, err := storage.GetAppDataDir()
	if err != nil {
		logging.Error("Failed to get app data directory: %v", err)
	}

	// Create the "Runtime" directory if it doesn't exist
	runtimeDir := path.Join(appDir, "Runtime")
	if _, err := os.Stat(runtimeDir); os.IsNotExist(err) {
		if err := os.MkdirAll(runtimeDir, 0755); err != nil {
			logging.Error("Failed to create Runtime directory: %v", err)
		}
	}

	// Define the path to the config file
	configPath := path.Join(runtimeDir, "config.json")

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create the config file with default values
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

	// Read and parse the config file
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

	return runtimeConfig.UserSettings.Server.AkiServerPath, runtimeConfig.UserSettings.Server.AkiServerAddress, runtimeConfig.UserSettings.LastProfile, runtimeConfig.UserSettings.ClientPath, nil
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
