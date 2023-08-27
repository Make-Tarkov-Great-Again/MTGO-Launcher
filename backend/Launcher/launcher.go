package launcher

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	wails "github.com/wailsapp/wails/v2/pkg/runtime"
)

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

var program *Launcher

func init() {
	program = NewLauncher()
}
type Storage struct {
}

func NewStorage() Storage {
	return Storage{}
}

func (s Storage) AddModEntry(mod string) {
	//what the fuck was i going to use this for??????????? i cannot fucking remember
}

func (s Storage) Check(neededSpace int64) {
	// TODO: Implement Storage space Check
}

func (s Storage) Clear() {

}
type Config struct {
}

func NewConfig() Config {
	return Config{}
}

//func (c Config) Update() {
//	// I dont think this is how i wanna do this specificly. kekw
//}

func (c Config) ClearIconCache() {
	// TODO: Implement ClearIconCache
	return
}

type Download struct {
}

func NewDownload() Download {
	return Download{}
}

func (d Download) Mod(modID string) {
	return
}
type Online struct {
}

func NewOnline() *Online {
	return &Online{}
}
// Checks if the app has a internet connection
//
// Returns:
//
//	bool: true if "online"
//	bool: false if !"online"
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
type Mod struct {
}

func NewMod() Mod {
	return Mod{}
}

// Throws conflict warning. Lets you pick to disable one of the conflicts, or contuine.
func (m Mod) ThrowConflict() {
	// TODO: Implement the ThrowConflict method
}

// Send missing mod popup. Cancel launch on "Cancel" and contuine on "I know what im doing!".
func (m Mod) ProfileThrowMissing() {
	// TODO: Implement the ProfileThrowMissing method
}
type UI struct {
	ctx context.Context
}

func NewUI() UI {
	// Initialize and return a UI instance
	return UI{}
}

// Send Panic popup message to app and closes on button press
func (u UI) Panic() {
	// TODO: Implement Panic popup
	//wails.Quit(u.ctx)
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

type App struct {
	ctx context.Context
}

func NewApp() App {
	// Initialize and return an App instance
	return App{}
}

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
type AKI struct {
}

func NewAKI() AKI {
	return AKI{}
}

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

	// lol no server, and no hoes
	if exePath == "" {
		title := "Starting AKI Server Failed"
		message := fmt.Sprintf("No server was found in AKI server path: %s. Is this the root folder of your AKI installation?", serverPath)
		program.UI.Error(title, message)                 // Call the UI.Error() method on the UI instance
		return nil, fmt.Errorf("%s: %s", title, message) // Return an error
	}

	cmd := exec.Command(exePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	process := cmd.Process
	childProcesses = append(childProcesses, process)

	return process, nil
}

type MTGA struct {
}

func NewMTGA() MTGA {
	// Initialize and return an MTGA instance
	return MTGA{}
}

func (m MTGA) StartServer() {
	// TODO: Implement the StartServer method
}
