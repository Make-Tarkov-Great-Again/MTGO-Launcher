/*
Package storage provides functionalities to manage data storage for the MTGO-Launcher.

This package offers methods to handle data storage, retrieval, and management within the application. It includes functions to work with various types of data such as profiles, configuration, and more.

Usage:

	import "mtgolauncher/backend/Storage"

Store profile data using the `StoreProfileData` function to save profile information to storage.

	profilePath := "C://AKI/user"
	err := storage.StoreProfileData(profilePath, false)
	if err != nil {
		fmt.Println("Error storing profile data:", err)
	}

# Valid commands for storage

# App Data Directory Initialization

  - [InitializeAppDataDir]: Initializes the app data directory for MTGO-Launcher. This function creates the necessary directory structure. Is called when applacation starts, so you probally dont need this.

# Storage

  - [GetAppDataDir]: Returns the path to the app data directory.

  - [CreateFolderIfNotExists]: Creates a folder at the specified path if it does not exist.

  - [WriteDataToFile]: Writes data to a file at the specified path.

  - [StoreData]: Stores data in a specified subdirectory and filename within the app data directory.

# Profile Data Storage

- [StoreProfileData]: Copies profile data to storage at "%appdata%/MT-GO/profiles". See function doc for more info.

# File Operations

- [CopyFile]: Copies a file from the source path to the destination path.

# UNFINISHED/UNSTARTED Functions

- [StoreMod]: (UNFINISHED/UNSTARTED) Stores mod data in storage.

- [StoreModList]: (UNFINISHED/UNSTARTED) Stores mod list data in storage.

- [StoreJSON]: (UNFINISHED/UNSTARTED) Stores JSON data in storage.

- [StoreBundles]: (UNFINISHED/UNSTARTED) Stores bundle data in storage.

- [StoreMisc]: (UNFINISHED/UNSTARTED) Stores miscellaneous data in storage.

Note: The storage package provides essential functionalities for handling data storage within the MTGO-Launcher application. The "UNFINISHED/UNSTARTED" functions listed above are placeholders for future development.

See the function documentation for detailed information and usage examples.
*/
package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

const appSubdir = "MT-GO"

/*
	Returns the AppDataDir for MTGO.

%appdata%/MT-GO
*/
func GetAppDataDir() (string, error) {
	appDataDir := os.Getenv("APPDATA")
	if appDataDir == "" {
		return "", fmt.Errorf("APPDATA environment variable not set")
	}
	return filepath.Join(appDataDir, appSubdir), nil
}

// QoD function
func CreateFolderIfNotExists(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return os.MkdirAll(folderPath, 0755)
	}
	return nil
}

// QoD function
func WriteDataToFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

// QoD function
func StoreData(subdir, filename string, data []byte) error {
	appDataDir, err := GetAppDataDir()
	if err != nil {
		return err
	}

	subdirPath := filepath.Join(appDataDir, subdir)
	err = CreateFolderIfNotExists(subdirPath)
	if err != nil {
		return err
	}

	filePath := filepath.Join(subdirPath, filename)
	return WriteDataToFile(filePath, data)
}

// Initilizes the appdata folder, only needs to be ran once, obviously.
func InitializeAppDataDir() error {
	appDataDir, err := GetAppDataDir()
	if err != nil {
		return err
	}

	return CreateFolderIfNotExists(appDataDir)
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

// Dont call. Call other functions insted. Like StoreProfileData
func storeDataWithType(dataType dataType, identifier, filename string, data interface{}) error {
	appDataDir, err := GetAppDataDir()
	if err != nil {
		return err
	}

	subdir := filepath.Join(appDataDir, dataType.folderName(identifier))
	err = CreateFolderIfNotExists(subdir)
	if err != nil {
		return err
	}

	filePath := filepath.Join(subdir, filename)

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		return WriteDataToFile(filePath, jsonData)
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

/*
StoreProfileData copies profile data to storage, @ "%appdata%/MT-GO"

Parameters:

- filePath: The absolute path to the source file containing the profile data.

- copyAll: A boolean flag indicating whether to copy the entire profile within the src folder (true) or just the character.json file (false).

Returns:
An error if there are any issues encountered during the storage process, otherwise nil.

Additional Info:

This function reads the character.json file to obtain profile information, including the profile's nickname, level, and faction. It uses the profile nickname to name the folder where the profile will be stored within "%appdata%/MT-GO/profiles". Additionally, a profile-details.json file is created within the profile folder, containing basic profile information in JSON format so we can quickly get basic info about the profile without reprasing the possibly long character.json.

Example usage:

err := storage.StoreProfileData("C://AKI/user", false)

	if err != nil {
	    fmt.Println("Error storing profile data:", err)
	}
*/
func StoreProfileData(filePath string, copyAll bool) error {
	appDataDir, err := GetAppDataDir()
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
	err = CreateFolderIfNotExists(profilesDir)
	if err != nil {
		return err
	}

	targetDir := filepath.Join(profilesDir, profiles.folderName(targetDirectory))
	err = CreateFolderIfNotExists(targetDir)
	if err != nil {
		return err
	}

	profileDataJSON, err := json.Marshal(profileData)
	if err != nil {
		return err
	}

	profileDetailsPath := filepath.Join(targetDir, "profile-details.json")
	err = WriteDataToFile(profileDetailsPath, profileDataJSON)
	if err != nil {
		return err
	}

	if copyAll {
		files, err := os.ReadDir(filepath.Dir(filePath))
		if err != nil {
			return err
		}

		for _, file := range files {
			srcFilePath := filepath.Join(filepath.Dir(filePath), file.Name())
			destFilePath := filepath.Join(targetDir, file.Name())
			err = CopyFile(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		}
	} else {
		newFilePath := filepath.Join(targetDir, "character.json")
		err = CopyFile(filePath, newFilePath)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
	CopyFile function copies a file from the source path to the destination path.

Parameters:
src - The source file path to be copied.

dest - The destination file path where the source file will be copied to.

Returns:
An error if the copying process encounters any issues, otherwise nil.

Example usage:

err := CopyFile("source.txt", "destination.txt")

	if err != nil {
	    fmt.Println("Error copying file:", err)
	}
*/
func CopyFile(src, dest string) error {
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

func StartListener() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	dirToWatch, err := GetAppDataDir()

	err = watchDirRecursive(watcher, dirToWatch)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Printf("New file created: %s\n", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("Error: %v\n", err)
			}
		}
	}()

	select {}
}

func watchDirRecursive(watcher *fsnotify.Watcher, dir string) error {
	err := watcher.Add(dir)
	if err != nil {
		return err
	}

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
}
