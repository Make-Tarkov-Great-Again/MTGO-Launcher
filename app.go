package main

import (
	"context"
	"fmt"
	launcher "mtgolauncher/backend/Launcher"
	log "mtgolauncher/backend/Logging"
	storage "mtgolauncher/backend/Storage"
	"mtgolauncher/backend/Storage/config"
)

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

// Initilize app.
func (a *App) startup(ctx context.Context) {
	log.Init()
	a.ctx = ctx

	err := storage.InitializeAppDataDir()
	if err != nil {
		fmt.Println("Error initializing app data directory:", err)
		return
	}

	log.Info("MTGO-Launcher version 0.0.1. This application falls under MIT licence. If you paid money for this, you got scammed. | https://github.com/Make-Tarkov-Great-Again/MTGO-Launcher")
	//Online check
	if launcher.NewOnline().Check() {
		fmt.Println("Connected to the internet... Starting in online mode.")
		//I guess i really dont have to do anything here...
	} else {
		fmt.Println("No internet connection... Starting in offline mode. Check your firewall if this is a error.")
		//TODO: Offline mode
	}
	//Jarvis, Initilize my subsystems pls. kthx
	config.Init()

	// TODO: Database check
	// TODO: If either fails -> Offline mode
}

// default greet thing
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// TODO: Database interactions
