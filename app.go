package main

import (
	"context"
	"fmt"
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
	a.ctx = ctx

	err := storage.InitializeAppDataDir() // Corrected function call
	if err != nil {
		fmt.Println("Error initializing app data directory:", err)
		return
	}

	err = storage.StoreProfileData("C:/Users/armyo/OneDrive/Desktop/character.json")
	if err != nil {
		fmt.Println("Error storing profile data:", err)
		return
	}

	fmt.Println("Data stored successfully!")
	// TODO: Online check
	// TODO: Database check
	// TODO: If either fails -> Offline mode
}

// default greet thing
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// TODO: Database interactions
