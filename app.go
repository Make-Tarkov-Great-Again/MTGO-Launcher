package main

import (
	"context"
	"fmt"
	launcher "mtgolauncher/backend/Launcher"
	log "mtgolauncher/backend/Logging"
	storage "mtgolauncher/backend/Storage"
)

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

// @param ctx context.Context
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
	//Initialization for logging.
	log.LogInit()
	log.Info("MTGO-Launcher version 0.0.1. This application falls under MIT licence. If you paid money for this, you got scammed. | https://github.com/Make-Tarkov-Great-Again/MTGO-Launcher")
	log.Info("Yeah so basiclly you wont see this unless your log surfing. In that case. Fuck you. Get off my lawn.", true)
	//Online check
	if app.Online.Check() {
		fmt.Println("Connected to the internet... Starting in online mode.")
		//I guess i really dont have to do anything here...
	} else {
		fmt.Println("No internet connection... Starting in offline mode. Check your firewall if this is a error.")
		//TODO: Offline mode
	}

	// Example mod information
	modID := 1
	modName := "AKI.AnotherAnimeModelMod"
	modAuthor := "YourMom"
	modDescription := "Imanage most of your mods being anime models."

	// Call CreateConfigIfNotExist to create the mod configuration
	err = app.Config.CreateConfigIfNotExist(modID, modName, modAuthor, modDescription)
	if err != nil {
		fmt.Println("Error creating mod configuration:", err)
		return
	}

	fmt.Println("Mod configuration created successfully!")

	// TODO: Database check
	// TODO: If either fails -> Offline mode
}

// default greet thing
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// TODO: Database interactions
