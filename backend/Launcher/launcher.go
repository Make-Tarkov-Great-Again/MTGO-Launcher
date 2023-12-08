/*
This package functions as the main wrapper package for commands used in runtime for the MTGO-Launcher.
Launcher employs method receivers/receiver functions to organize these commands for ease of use.

It is recommended to import Launcher and then alias it with "app" or "a" for ease of use.

# Usage

	app := launcher.NewLauncher()

	if app.Online.Check() {
		// Do something if there is a connection.
	} else {
		// Do something if there isn't a connection.
	}

# Valid commands for launcher

# Storage

  - [Storage.AddModEntry]
  - [Storage.Clear]
  - [Storage.Check]

# Config

  - [Config.ClearIconCache]

# Download

  - [Download.Mod]

# Online

  - [Online.Check]
  - [Online.Heartbeat]

# Mod

  - [Mod.ProfileThrowMissing]
  - [Mod.ThrowConflict]

# UI

  - [UI.Error]
  - [UI.Info]
  - [UI.Panic]
  - [UI.Reload]

# App

  - [App.Close]
  - [App.CloseServers]
  - [App.Hide]
  - [App.Minimize]
  - [App.Show]

# AKI

  - [AKI.StartServer]

# MTGA

  - [MTGA.StartServer]
*/
package launcher

/*
	Storage:  NewStorage(),
	Config:   NewConfig(),
	Download: NewDownload(),
	Online:   NewOnline(),
	Mod:      NewMod(),
	UI:       NewUI(),
	App:      NewApp(),
	AKI:      NewAKI(),
	MTGA:     NewMTGA(),
*/
import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mholt/archiver"
	wails "github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/sys/windows"

	flog "mtgolauncher/backend/Logging"
	appData "mtgolauncher/backend/Storage"
	config "mtgolauncher/backend/Storage/config"
)

var WailsContext context.Context

//#region Storage

func (s *Storage) AddModEntry(mod string) {
	//what the fuck was i going to use this for??????????? i cannot fucking remember
}

func (s *Storage) moveContents(srcDir, destDir string) error {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		src := filepath.Join(srcDir, file.Name())
		dest := filepath.Join(destDir, file.Name())

		srcFile, err := os.Open(src)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer destFile.Close()

		if _, err := io.Copy(destFile, srcFile); err != nil {
			return err
		}

		if err := os.Remove(src); err != nil {
			return err
		}
	}

	return nil
}

// This function uses code from the "github.com/mholt/archiver" library,
// which is licensed under the MIT License.
func (s *Storage) ExtractArchive(archivePath, destinationDir string) error {
	err := archiver.Unarchive(archivePath, destinationDir)
	if err != nil {
		return err
	}
	return nil
}

/*
Package Launcher

# Storage.Check

Storage.Check checks the available disk space on the system drive to determine if the required amount of space is available.

Parameters:

- neededSpace: The amount of free space needed, in bytes.

Returns:

- true if the system has at least the specified amount of free space.

- false if the system does not have enough free space or if an error occurs while checking.

Example usage:

hasEnoughSpace := storage.Check(1024 * 1024 * 1024) // 1 GB

	if hasEnoughSpace {
	    fmt.Println("Sufficient disk space is available.")
	} else {

	    fmt.Println("Insufficient disk space.")
	}
*/
func (s *Storage) Check(neededSpace int64) bool {
	var freeBytesAvailable uint64
	var totalNumberOfBytes uint64
	var totalNumberOfFreeBytes uint64
	err := windows.GetDiskFreeSpaceEx(windows.StringToUTF16Ptr("C:"),
		&freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return freeBytesAvailable >= uint64(neededSpace)
}

func (s *Storage) Clear() {
	err := os.RemoveAll(s.AppDataDir)
	if err != nil {
		fmt.Printf("[Launcher.Storage] Failed to removed all files from %s. \n %s", s.AppDataDir, err)
	}
}

func (s *Storage) ClearIconCache() {
	// TODO: Implement ClearIconCache
	return
}

//#endregion Storage

// #region Config

// GetConfig retrieves the global configuration.
func (c *Config) GetConfig() *Config {
	// Implement logic to retrieve the global configuration.
	return c
}

type ModDetails struct {
	ModInfo struct {
		Name         string     `json:"Name"`
		Version      ModVersion `json:"Version"`
		ID           int        `json:"ID"`
		Author       string     `json:"Author"`
		Description  string     `json:"Description"`
		DownloadedAt time.Time  `json:"DownloadedAt"`
	}
	ModConfig map[string]interface{} `json:"ModConfig"`
}

type ModVersion struct {
	Emu        string `json:"Emu"`
	EmuVersion string `json:"EmuVersion"`
}

func (c *Config) CreateConfigIfNotExist(modID int, modName string, modAuthor string, modDescription string) error {
	appDataDir := c.AppDataDir

	err := appData.CreateFolderIfNotExists(filepath.Join(appDataDir, "Mods"))
	if err != nil {
		return err
	}

	modDir := filepath.Join(appDataDir, "Mods", fmt.Sprintf("%d", modID))
	modConfigFile := filepath.Join(modDir, "mod-details.json")

	if _, err := os.Stat(modConfigFile); os.IsNotExist(err) {
		err = appData.CreateFolderIfNotExists(modDir)
		if err != nil {
			return err
		}

		// Read package.json file to determine emulator type and version
		packageJSONPath := filepath.Join(modDir, "package.json")
		packageJSON, err := os.ReadFile(packageJSONPath)
		if err != nil {
			return err
		}

		var packageInfo struct {
			Name        string `json:"name"`
			AkiVersion  string `json:"AkiVersion"`
			MtgaVersion string `json:"MtgaVersion"`
			Author      string `json:"author"`
			Description string `json:"description"`
		}

		err = json.Unmarshal(packageJSON, &packageInfo)
		if err != nil {
			return err
		}

		emulatorType := "Unknown"
		emulatorVersion := "Unknown"

		if packageInfo.AkiVersion != "" {
			emulatorType = "AKI"
			emulatorVersion = packageInfo.AkiVersion
		} else if packageInfo.MtgaVersion != "" {
			emulatorType = "MTGA"
			emulatorVersion = packageInfo.MtgaVersion
		}

		// Find and read config.json files
		configFiles, err := config.FindConfigFiles(modDir)
		if err != nil {
			return err
		}

		// Create init structure
		modDetails := ModDetails{
			ModInfo: struct {
				Name         string     `json:"Name"`
				Version      ModVersion `json:"Version"`
				ID           int        `json:"ID"`
				Author       string     `json:"Author"`
				Description  string     `json:"Description"`
				DownloadedAt time.Time  `json:"DownloadedAt"`
			}{
				Name: modName,
				Version: ModVersion{
					Emu:        emulatorType,
					EmuVersion: emulatorVersion,
				},
				ID:           modID,
				Author:       packageInfo.Author,
				Description:  packageInfo.Description,
				DownloadedAt: time.Now(),
			},
			ModConfig: make(map[string]interface{}),
		}

		// Unmarshal and add config.json files to ModConfig
		for key, filePath := range configFiles {
			configJSON, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}
			var configData interface{}
			err = json.Unmarshal(configJSON, &configData)
			if err != nil {
				return err
			}
			modDetails.ModConfig[key] = configData
		}

		// Marshal
		modDetailsJSON, err := json.MarshalIndent(modDetails, "", "  ")
		if err != nil {
			return err
		}

		// Write
		err = appData.StoreData("Mods", fmt.Sprintf("/%d/mod-details.json", modID), modDetailsJSON)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateConfig updates the global configuration.
func (c *Config) UpdateConfig(updatedConfig *Config) error {
	// Implement logic to update the global configuration.
	return fmt.Errorf("not implemented")
}

// GetConfigByModID retrieves configuration for a specific mod by its ID.
func (c *Config) GetConfigByModID(modID string) (*Config, error) {
	// Implement logic to retrieve configuration by mod ID.
	return nil, fmt.Errorf("not implemented")
}

// GetConfigByModName retrieves configuration for a specific mod by its name.
func (c *Config) GetConfigByModName(modName string) (*Config, error) {
	// Implement logic to retrieve configuration by mod name.
	return nil, fmt.Errorf("not implemented")
}

//#endregion Config

// #region Download

func sendErrorMessage(wsConn *websocket.Conn, err error) {
	errorMessage := struct {
		Step  string `json:"step"`
		Error string `json:"error"`
	}{Step: "Error", Error: err.Error()}

	sendMessage(wsConn, errorMessage)
}

func sendMessage(wsConn *websocket.Conn, message interface{}) {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		NewUI().Error("JSON Marshaling Error", err.Error())
		return
	}
	wsConn.WriteMessage(websocket.TextMessage, jsonMessage)
}

func (d *Download) EnqueueDownload(modID int, fileURL string, wsConn *websocket.Conn, modName string, modAuthor string) {
	d.queueMutex.Lock()
	defer d.queueMutex.Unlock()

	request := &DownloadRequest{
		ModID:     modID,
		FileURL:   fileURL,
		WsConn:    wsConn,
		ModName:   modName,
		ModAuthor: modAuthor,
	}

	d.queue = append(d.queue, request)
}

func (d *Download) DequeueDownload() *DownloadRequest {
	d.queueMutex.Lock()
	defer d.queueMutex.Unlock()

	if len(d.queue) == 0 {
		return nil
	}

	request := d.queue[0]
	d.queue = d.queue[1:]
	return request
}

// Starts the download listener
func (d *Download) StartDownloading() {
	go func() {
		for {
			request := d.DequeueDownload()
			if request == nil {
				// No bitches for me to interact with, im sleeping
				time.Sleep(time.Second)
				continue
			}

			err := d.Mod(request.ModID, request.FileURL, request.WsConn, request.ModName, request.ModAuthor)
			if err != nil {
				fmt.Printf("Failed to download mod: %s\n", err)
			}
		}
	}()
}

// This function will be greatly changed for how it gets mods, you can already see hints of this, via "MOD-ID",
// Mods will be gotten in a secure and centerailzed way in not too long,
func (d *Download) Mod(modID int, fileURL string, wsConn *websocket.Conn, modName string, modAuthor string) error {
	// TODO: Get download URL from Database via ModID
	//database.request.mod(modID)

	// Start the download
	resp, err := http.Get(fileURL)
	if err != nil {
		NewUI().Error("Download Error", "Failed to download mod data: "+err.Error())
		sendErrorMessage(wsConn, err)
		return err
	}
	defer resp.Body.Close()

	// File name getter thing through header and backup for if the header doesn't exist
	fileName := resp.Header.Get("Content-Disposition")
	if fileName == "" {
		u, err := url.Parse(fileURL)
		if err != nil {
			NewUI().Error("Download Error", err.Error())
			sendErrorMessage(wsConn, err)
			return err
		}
		fileName = path.Base(u.Path)
	} else {
		_, params, err := mime.ParseMediaType(fileName)
		if err == nil {
			fileName = params["filename"]
		}
	}

	// Convert mod id to string because haha funny number can't be a string ever
	modIDst := strconv.Itoa(modID)
	downloadDir := path.Join(d.ModsFolder, modIDst)
	if err := appData.CreateFolderIfNotExists(downloadDir); err != nil {
		NewUI().Error("Download Error", err.Error())
		sendErrorMessage(wsConn, err)
		return err
	}

	newfilename := modIDst + strings.ToLower(filepath.Ext(fileName))
	localFilePath := path.Join(downloadDir, newfilename)

	localFile, err := os.Create(localFilePath)
	if err != nil {
		NewUI().Error(fmt.Sprintf("Failed to download %s", modName), fmt.Sprintf("%s", err))
		sendErrorMessage(wsConn, err)
		return err
	}
	defer localFile.Close()

	fileSize := resp.ContentLength
	var downloaded int64
	formattedFileSize := strconv.FormatInt(fileSize, 10)

	flog.OnlineLog(fmt.Sprintf("Downloading %s, %s bytes, to %s", fileName, formattedFileSize, localFilePath))
	buffer := make([]byte, 1024)

	var name string = ""

	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break // End of file
			}
			NewUI().Error("Download Error", "Failed to download mod data: "+err.Error())
			sendErrorMessage(wsConn, err)
			return err
		}

		_, err = localFile.Write(buffer[:n])
		if err != nil {
			NewUI().Error("Download Error", "Failed to download mod data: "+err.Error())
			sendErrorMessage(wsConn, err)
			return err
		}

		downloaded += int64(n)

		percentage := int(float64(downloaded) / float64(fileSize) * 100)

		if modName == "" {
			modName = fileName
		}

		progress := struct {
			FileName   string `json:"name"`
			FileSize   int64  `json:"filesize"`
			Percentage int    `json:"percentage"`
			Modname    string `json:"modName"`
			QueueSize  int    `json:"queueSize"`
		}{FileName: name, FileSize: fileSize, Percentage: percentage, Modname: modName, QueueSize: len(d.queue)}

		jsonProgress, err := json.Marshal(progress)
		if err != nil {
			sendErrorMessage(wsConn, err)
			NewUI().Error("Download Error", ""+err.Error())
			return err
		}

		wsConn.WriteMessage(websocket.TextMessage, jsonProgress)
	}

	progress := struct {
		FileName   string `json:"name"`
		FileSize   int64  `json:"filesize"`
		Percentage int    `json:"percentage"`
		Modname    string `json:"modName"`
		QueueSize  int    `json:"queueSize"`
	}{FileName: name, FileSize: fileSize, Percentage: 100, Modname: modName, QueueSize: len(d.queue)}

	jsonProgress, err := json.Marshal(progress)
	if err != nil {
		sendErrorMessage(wsConn, err)
		NewUI().Error("Download Error", ""+err.Error())

		return err
	}

	wsConn.WriteMessage(websocket.TextMessage, jsonProgress)

	ext := strings.ToLower(filepath.Ext(fileName))
	if ext == ".zip" || ext == ".rar" || ext == ".7z" || ext == ".tar" {
		extractionDir := downloadDir
		if err := appData.CreateFolderIfNotExists(extractionDir); err != nil {
			NewUI().Error("Extraction Error", "Failed to extract mod: "+err.Error())
			sendErrorMessage(wsConn, err)
			return err
		}

		startExtractionMessage := struct {
			Step      string `json:"step"`
			QueueSize int    `json:"queueSize"`
		}{Step: "Extracting", QueueSize: len(d.queue)}

		jsonStartExtraction, err := json.Marshal(startExtractionMessage)
		if err != nil {
			NewUI().Error("Download Error", ""+err.Error())
			sendErrorMessage(wsConn, err)
			return err
		}
		wsConn.WriteMessage(websocket.TextMessage, jsonStartExtraction)

		if err := NewStorage().ExtractArchive(localFilePath, extractionDir); err != nil {
			defer localFile.Close()
			NewUI().Error("Download Error", ""+err.Error())
			sendErrorMessage(wsConn, err)
			return err
		}

		// List the contents of the extractionDir
		contents, err := os.ReadDir(extractionDir)
		if err != nil {
			NewUI().Error("Download Error", ""+err.Error())
			sendErrorMessage(wsConn, err)
			return err
		}

		//i hate my life. hour 13 of trying to get this to work as i want it.
		for _, item := range contents {
			//is it a folder?
			if item.IsDir() {
				//find package.json, if it exists, move everything within the same folder as it to the root of the mods folder (eg, %appdata%/MT-GO/Mods/modID) and pray nothing goes wrong
				packageJSONPath := filepath.Join(extractionDir, item.Name(), "package.json")
				if _, err := os.Stat(packageJSONPath); err == nil {
					files, err := os.ReadDir(filepath.Join(extractionDir, item.Name()))
					if err != nil {
						NewUI().Error("Download Error", ""+err.Error())
						sendErrorMessage(wsConn, err)
						return err
					}
					for _, file := range files {
						src := filepath.Join(extractionDir, item.Name(), file.Name())
						dest := filepath.Join(downloadDir, file.Name())
						if err := os.Rename(src, dest); err != nil {
							NewUI().Error("Move Error", "Failed to move mod "+modName+" to local storage: "+err.Error())
							sendErrorMessage(wsConn, err)
							return err
						}
					}
					// Remove the empty shit.
					if err := os.Remove(filepath.Join(extractionDir, item.Name())); err != nil {
						sendErrorMessage(wsConn, err)
						return err
					}
				}
			}

		}
	}
	fmt.Printf("should make config now lol")
	err = NewConfig().CreateConfigIfNotExist(modID, modName, modAuthor, "damn")
	defer os.Remove(localFilePath)
	if err != nil {
		NewUI().Error("Configuration Error", "Failed to create mod configuration: "+err.Error())
		sendErrorMessage(wsConn, err)
		return err
	}
	//she done pt2
	extractionDoneMessage := struct {
		Step string `json:"step"`
	}{Step: "Done"}

	jsonExtractionDone, err := json.Marshal(extractionDoneMessage)
	if err != nil {
		sendErrorMessage(wsConn, err)
		return err
	}

	wsConn.WriteMessage(websocket.TextMessage, jsonExtractionDone)

	return nil
}

//#endregion Config

//#region Online
/*
Package launcher

# Online.Check

Check checks if the app has an internet connection.

Returns:
- true if the app is online (able to reach "http://www.google.com").
- false if the app is offline or encounters an error while checking the connection.

Example usage:

isOnline := online.Check()

	if isOnline {
	    fmt.Println("App is connected to the internet.")
	} else {

	    fmt.Println("App is offline.")
	}
*/
func (o *Online) Check() bool {
	_, err := http.Get("http://www.google.com")
	if err != nil {
		return false
	}
	return true
}

// Returns { alive } if successful probally
func (o *Online) Heartbeat() {
	//Database heartbeat. no database yet lol
	return
}

//#endregion Online

//#region Mod

// Throws conflict warning. Lets you pick to disable one of the conflicts, or contuine.
func (m *Mod) ThrowConflict() {
	// TODO: Implement the ThrowConflict method
}

// Send missing mod popup. Cancel launch on "Cancel" and contuine on "I know what im doing!".
func (m *Mod) ProfileThrowMissing() {
	// TODO: Implement the ProfileThrowMissing method
}

func (m *Mod) ActivateMod() {
	//TODO: Implment Mod activation.
}

func (m *Mod) DisableMods() {
	//TODO: Implment Mod disabling.
	// -> Config.go
}

//#endregion Mod

// #region UI

// Send Panic popup message to app and closes on button press
func (u *UI) Panic(title string, message string) {
	selection, err := wails.MessageDialog(WailsContext, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Message:       "Whoops, something went very wrong, and we can't continue! See error below!\n" + message,
		Buttons:       []string{"Ok"},
		DefaultButton: "Ok",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	if selection != "" {
		wails.Quit(WailsContext)
	}
	return
}

func (u *UI) OpenFileSelector(title string, filters []struct{ displayName, pattern string }) string {
	selection, err := wails.OpenDirectoryDialog(WailsContext, wails.OpenDialogOptions{
		DefaultDirectory:     "C:\\",
		Title:                title,
		ShowHiddenFiles:      true,
		CanCreateDirectories: true,
		ResolvesAliases:      true,
	})
	if err != nil {
		return err.Error()
	}
	if selection == "" {
		flog.Error("Selection came back empty")
		return "Error"
	}
	return selection
}

// Send Panic popup message to app and closes on button press
func (u *UI) PanicStatement(title string, message string) {
	selection, err := wails.MessageDialog(WailsContext, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Message:       message,
		Buttons:       []string{"Ok"},
		DefaultButton: "Ok",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	if selection != "" {
		wails.Quit(WailsContext)
	}
	return
}

// Send error popup message to app
func (u *UI) Error(title string, message string) {
	fmt.Println("UI Error ctx:")
	fmt.Println(u.ctx)
	selection, err := wails.MessageDialog(WailsContext, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Title:         title,
		Message:       message,
		Buttons:       []string{"Continue", "Exit"},
		DefaultButton: "Continue",
	})
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
	}
	switch selection {
	case "Exit":
		wails.Quit(u.ctx)
	default:
		fmt.Printf("selection: %v\n", selection)

	}
}

func (u *UI) ErrorStatement(title string, message string) {
	fmt.Println("UI Error ctx:")
	fmt.Println(u.ctx)
	selection, err := wails.MessageDialog(WailsContext, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Title:         title,
		Message:       "Whoops, something went wrong, but we can continue! See error below!\n\n" + message,
		Buttons:       []string{"Continue", "Exit"},
		DefaultButton: "Continue",
	})
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
	}
	switch selection {
	case "Exit":
		wails.Quit(u.ctx)
	default:
		fmt.Printf("selection: %v\n", selection)

	}
}

func (u *UI) Question(title, message string) bool {
	selection, err := wails.MessageDialog(WailsContext, wails.MessageDialogOptions{
		Type:    wails.QuestionDialog,
		Title:   title,
		Message: message,
	})
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
	}
	switch selection {
	case "No":
		return false
	default:
		fmt.Printf("selection: %v\n", selection)
		return true
	}
}

//func (u *UI) TestError(title string, message string, ctx context.Context) {
//	runtime.MessageDialog(u.ctx, runtime.MessageDialogOptions{
//		Type:  runtime.InfoDialog,
//		Title: "Works",
//	})
//}

// Reloads frontend.
func (u *UI) Reload() {
	wails.WindowReloadApp(WailsContext)
}

//#endregion UI

// #region App
// Minimizes the launcher
func (a App) Minimize() {
	wails.WindowMinimise(a.ctx)
}

// Closes the launcher
func (a App) Close() {
	wails.Quit(a.ctx)
}

// Hides the launcher from view entirely
func (a App) Hide() {
	wails.WindowHide(a.ctx)
}

// Shows the launcher if its hidden.
func (a App) Show() {
	wails.WindowShow(a.ctx)
}

var childProcesses []*os.Process

// Kills all child processes, So closes the servers.
func (a App) CloseServers() {
	for _, process := range childProcesses {
		err := process.Signal(syscall.SIGTERM)
		if err != nil {
			fmt.Println("Error killing process:", err)
		}
	}
}

//#endregion App

//#region EMU

/*
MTGA.StartServer starts an MTGA server from the specified serverPath.

Parameters:
- serverPath: The path to the MTGA server executable.

Returns:
- *os.Process: A pointer to the started process if successful.
- error: An error if the server couldn't be started.

Example usage:
process, err := mtga.StartServer("C:/path/to/aki-server")

	if err != nil {
	    fmt.Println("Error starting MTGA server:", err)
	} else {

	    fmt.Println("mtga server started with process ID:", process.Pid)
	}
*/
// Starts the MTGA server via @mtga-path
func (m *MTGA) StartServer(serverPath string) (*os.Process, error) {
	cmd := exec.Command("go", "run", "backend.go")
	cmd.Dir = serverPath

	// Define the outputCallback function
	outputCallback := func(line string) {
		flog.AKIServerOutput(line)
		fmt.Println("Server Output:", line)
	}

	// Set up pipes for capturing standard output and standard error
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	// Start the command
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	process := cmd.Process
	childProcesses = append(childProcesses, process)

	// Goroutine to capture standard output
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			output := scanner.Text()
			outputCallback(output) // Call the provided callback with the captured output
		}
	}()

	// Goroutine to capture standard error
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			output := scanner.Text()
			outputCallback(output) // Call the provided callback with the captured output
		}
	}()

	return process, nil
}

//#endregion EMU

// #region structs

type Storage struct {
	AppDataDir string
}
type Config struct {
	AppDataDir string
}
type Download struct {
	AppDataDir string
	ModsFolder string
	queue      []*DownloadRequest
	queueMutex sync.Mutex
}

type DownloadRequest struct {
	ModID     int
	FileURL   string
	WsConn    *websocket.Conn
	ModName   string
	ModAuthor string
}
type Online struct {
}
type Mod struct {
}
type UI struct {
	ctx context.Context
}

func (u *UI) Startup(ctx context.Context) {
	u.ctx = ctx
}

type App struct {
	ctx context.Context
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

type AKI struct {
	ctx context.Context
}

func (a *AKI) Startup(ctx context.Context) {
	a.ctx = ctx
	WailsContext = ctx
}

type MTGA struct {
}

//#endregion

//#region === Component Initialization ===

// NewLauncher creates a new Launcher instance with initialized components.

// NewStorage creates and returns a new Storage instance.
func NewStorage() *Storage {
	appDataDir, err := appData.GetAppDataDir()
	if err != nil {
		fmt.Println("Error:", err)
		return &Storage{}
	}

	return &Storage{
		AppDataDir: appDataDir,
	}
}

// NewConfig creates and returns a new Config instance.
func NewConfig() *Config {
	appDataDir, err := appData.GetAppDataDir()
	if err != nil {
		fmt.Println("Error:", err)
		return &Config{}
	}

	return &Config{
		AppDataDir: appDataDir,
	}
}

// NewDownload creates and returns a new Download instance.
func NewDownload() *Download {

	appDataDir, err := appData.GetAppDataDir()
	ModsFolder := path.Join(appDataDir, "Mods")
	if err != nil {
		fmt.Println("Error:", err)
		return &Download{}
	}

	return &Download{
		AppDataDir: appDataDir,
		ModsFolder: ModsFolder,
	}
}

// NewOnline creates and returns a new Online instance.
func NewOnline() *Online {
	return &Online{}
}

// NewMod creates and returns a new Mod instance.
func NewMod() *Mod {
	return &Mod{}
}

// NewUI creates and returns a new UI instance.
func NewUI() *UI {
	return &UI{}
}

// NewApp creates and returns a new App instance.
func NewApp() *App {
	return &App{}
}

// NewAKI creates and returns a new AKI instance.
func NewAKI() *AKI {
	return &AKI{}
}

// NewMTGA creates and returns a new MTGA instance.
func NewMTGA() *MTGA {
	return &MTGA{}
}

//#endregion === End of Component Initialization ===
