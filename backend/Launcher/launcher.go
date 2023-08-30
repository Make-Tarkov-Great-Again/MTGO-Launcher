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
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"

	appData "mtgolauncher/backend/Storage"

	wails "github.com/wailsapp/wails/v2/pkg/runtime"
)

var program *Launcher

func init() {
	program = NewLauncher()
}

//#region Storage

func (s Storage) AddModEntry(mod string) {
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
func (s Storage) Check(neededSpace int64) bool {
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

func (s Storage) Clear() {
	err := os.RemoveAll(s.AppDataDir)
	if err != nil {
		fmt.Printf("[Launcher.Storage] Failed to removed all files from %s. \n %s", s.AppDataDir, err)
	}
}
func (s Storage) ClearIconCache() {
	// TODO: Implement ClearIconCache
	return
}

//#endregion Storage

//#region Config
//func (c Config) Update() {
//	// I dont think this is how i wanna do this specificly. kekw
//}

func (c Config) ClearIconCache() {
	// TODO: Implement ClearIconCache
	return
}

//#endregion Config

// #region Download
func (d Download) Mod(modID string) {
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
func (m Mod) ThrowConflict() {
	// TODO: Implement the ThrowConflict method
}

// Send missing mod popup. Cancel launch on "Cancel" and contuine on "I know what im doing!".
func (m Mod) ProfileThrowMissing() {
	// TODO: Implement the ProfileThrowMissing method
}

//#endregion Mod

// #region UI
// Send Panic popup message to app and closes on button press
func (u UI) Panic(title string, message string) {
	selection, err := wails.MessageDialog(u.ctx, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Message:       "Whoops, something went very wrong, and we can't continue! See error below!\n" + message,
		Buttons:       []string{"Ok"},
		DefaultButton: "Ok",
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	if selection == "Ok" {
		wails.Quit(u.ctx)
	}
	return
}

// Send error popup message to app
func (u UI) Error(title string, message string) {
	wails.MessageDialog(u.ctx, wails.MessageDialogOptions{
		Type:          wails.ErrorDialog,
		Title:         title,
		Message:       message,
		DefaultButton: "Ok",
	})
}

// Send info popup message to app
func (u UI) Info() {
	// TODO: Implement Info popup
	return
}

// Reloads frontend.
func (u UI) Reload() {
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

//endregion App

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
func (a AKI) StartServer(serverPath string) (*os.Process, error) {
	files, err := os.ReadDir(serverPath)
	if err != nil {
		return nil, err
	}

	var exePath string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".exe") {
			exePath = filepath.Join(serverPath, file.Name())
			break
		}
	}

	if exePath == "" {
		title := "Starting AKI Server Failed"
		message := fmt.Sprintf("No server was found in AKI server path: %s. Is this the root folder of your AKI installation?", serverPath)
		program.UI.Error(title, message)
		return nil, fmt.Errorf("%s: %s", title, message)
	}

	// Define the outputCallback function
	outputCallback := func(line string) {
		fmt.Println("Server Output:", line)
		// You can perform additional actions here as needed
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
func (m MTGA) StartServer() {
	// TODO: Implement the StartServer method
}

//#endregion EMU

// #region structs
type Launcher struct {
	Storage  Storage
	Config   Config
	Download Download
	Online   *Online
	Mod      Mod
	UI       UI
	App      App
	AKI      AKI
	MTGA     MTGA
}

type Storage struct {
	AppDataDir string
}
type Config struct {
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
type App struct {
	ctx context.Context
}
type AKI struct {
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
func NewStorage() Storage {
	appDataDir, err := appData.GetAppDataDir()
	if err != nil {
		fmt.Println("Error:", err)
		return Storage{}
	}

	return Storage{
		AppDataDir: appDataDir,
	}
}

// NewConfig creates and returns a new Config instance.
func NewConfig() Config {
	return Config{}
}

// NewDownload creates and returns a new Download instance.
func NewDownload() Download {
	return Download{}
}

// NewOnline creates and returns a new Online instance.
func NewOnline() *Online {
	return &Online{}
}

// NewMod creates and returns a new Mod instance.
func NewMod() Mod {
	return Mod{}
}

// NewUI creates and returns a new UI instance.
func NewUI() UI {
	return UI{}
}

// NewApp creates and returns a new App instance.
func NewApp() App {
	return App{}
}

// NewAKI creates and returns a new AKI instance.
func NewAKI() AKI {
	return AKI{}
}

// NewMTGA creates and returns a new MTGA instance.
func NewMTGA() MTGA {
	return MTGA{}
}

//#endregion === End of Component Initialization ===
