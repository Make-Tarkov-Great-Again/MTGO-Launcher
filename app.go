package main

import (
	"context"
	"fmt"
	launcher "mtgolauncher/backend/Launcher"
	log "mtgolauncher/backend/Logging"
	notifications "mtgolauncher/backend/Notifications"
	profile "mtgolauncher/backend/Profile"
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
	go log.Init()
	a.ctx = ctx
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
	go func() {
		err := storage.InitializeAppDataDir()
		if err != nil {
			fmt.Println("Error initializing app data directory:", err)
		}
	}()
	go config.Init()
	go profile.ParseAndSaveProfile("E:\\SPT-AKI\\user\\profiles\\beb27d576cc4f1b8b8fae52d.json")
	go storage.StartListener()
	notifications.UpdateAvaiable()

	// TODO: Database check
	// TODO: If either fails -> Offline mode
}

// default greet thing
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// TODO: Database interactions
