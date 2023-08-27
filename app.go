package main

import (
	"context"
	"fmt"
	launcher "mtgolauncher/backend/Launcher"
	storage "mtgolauncher/backend/Storage"
)

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

// Functions to run on startup
func (a *App) startup(ctx context.Context) {
	app := launcher.NewLauncher()
	a.ctx = ctx

	err := storage.InitializeAppDataDir() // Corrected function call
	if err != nil {
		fmt.Println("Error initializing app data directory:", err)
		return
	}
	fmt.Println("Data stored successfully!")

	//Online check
	if app.Online.Check() {
		fmt.Println("Connected to the internet... Starting in online mode.")
		//I guess i really dont have to do anything here...
	} else {
		fmt.Println("No internet connection... Starting in offline mode. Check your firewall if this is a error.")
		//TODO: Offline mode
	}

	// TODO: Database check
	// TODO: If either fails -> Offline mode
}

// default greet thing
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// TODO: Database interactions
