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
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows"

	appData "mtgolauncher/backend/Storage"

	config "mtgolauncher/backend/Storage/config"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	wails "github.com/wailsapp/wails/v2/pkg/runtime"
)

var program *Launcher

func init() {
	program = NewLauncher()
}

//#region Storage

func (s *Storage) AddModEntry(mod string) {
	//what the fuck was i going to use this for??????????? i cannot fucking remember
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
func (d *Download) Mod(modID string) {
	return
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
	selection, err := wails.MessageDialog(u.ctx, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Message:       "Whoops, something went very wrong, and we can't continue! See error below!\n" + message,
		Buttons:       []string{"Ok"},
		DefaultButton: "Ok",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	if selection != "" {
		wails.Quit(u.ctx)
	}
	return
}

// Send error popup message to app
func (u *UI) Error(title string, message string) {
	selection, err := wails.MessageDialog(u.ctx, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Title:         title,
		Message:       "Whoops, something went wrong, but we can continue! See error below!\n" + message,
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

// Backup error function incase above function doesnt work, due to CTX
func (u *UI) Errorctx(title string, message string, ctx context.Context) {
	selection, err := wails.MessageDialog(u.ctx, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Title:         title,
		Message:       "Whoops, something went wrong, but we can continue! See error below!\n" + message,
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
//func (u *UI) TestError(title string, message string, ctx context.Context) {
//	runtime.MessageDialog(u.ctx, runtime.MessageDialogOptions{
//		Type:  runtime.InfoDialog,
//		Title: "Works",
//	})
//}

// Send info popup message to app
func (u *UI) Info() {
	// TODO: Implement Info popup
	return
}

// Reloads frontend.
func (u *UI) Reload() {
	wails.WindowReloadApp(u.ctx)
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
AKI.StartServer starts an AKI server from the specified serverPath.

Parameters:
- serverPath: The path to the AKI server executable.

Returns:
- *os.Process: A pointer to the started process if successful.
- error: An error if the server couldn't be started.

Example usage:
process, err := aki.StartServer("C:/path/to/aki-server")

	if err != nil {
	    fmt.Println("Error starting AKI server:", err)
	} else {

	    fmt.Println("AKI server started with process ID:", process.Pid)
	}
*/
func (a *AKI) StartServer(serverPath string) (*os.Process, error) {
	files, err := os.ReadDir(serverPath)
	if err != nil {
		return nil, err
	}

	var exePath string
	for _, file := range files {
		if strings.HasPrefix(filepath.Base(file.Name()), "Aki.Server") {
		if strings.HasSuffix(file.Name(), ".exe") {
			exePath = filepath.Join(serverPath, file.Name())
			break
			} else {
				program.UI.Errorctx("Unknown file", "The file you have selected for your AKI Path doesn't seem to be an AKI Server... Make sure it's named \"AKI.Server\" and is a .exe file!", a.ctx)
			}
		}
	}

	if exePath == "" {
		title := "Starting AKI Server Failed"
		message := fmt.Sprintf("No server was found in AKI server path: %s. Is this the root folder of your AKI installation?", serverPath)
		program.UI.Errorctx(title, message, a.ctx)
		return nil, fmt.Errorf("%s: %s", title, message)
	}

	// Define the outputCallback function
	outputCallback := func(line string) {
		fmt.Println("Server Output:", line)
	}

	// Create the command
	cmd := exec.Command(exePath)

	// Set up pipes for capturing standard output and standard error
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	// Start the command
	err = cmd.Start()
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

// Starts the MTGA server via @mtga-path
func (m *MTGA) StartServer() {
	// TODO: Implement the StartServer method
}

//#endregion EMU

// #region structs
type Launcher struct {
	Storage  *Storage
	Config   *Config
	Download *Download
	Online   *Online
	Mod      *Mod
	UI       *UI
	App      *App
	AKI      *AKI
	MTGA     *MTGA
}

type Storage struct {
	AppDataDir string
}
type Config struct {
	AppDataDir string
}
type Download struct {
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
}

type MTGA struct {
}

//#endregion

//#region === Component Initialization ===

// NewLauncher creates a new Launcher instance with initialized components.
func NewLauncher() *Launcher {
	return &Launcher{
		Storage:  NewStorage(),
		Config:   NewConfig(),
		Download: NewDownload(),
		Online:   NewOnline(),
		Mod:      NewMod(),
		UI:       NewUI(),
		App:      NewApp(),
		AKI:      NewAKI(),
		MTGA:     NewMTGA(),
	}
}

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
	return &Download{}
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
