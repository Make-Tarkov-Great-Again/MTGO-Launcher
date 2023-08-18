package main

import (
	"context"
	"fmt"
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
	//TODO: Online check
	//TODO: Database check
	//TODO: If either fail -> Offline mode
}

// default greet thing
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

//TODO: Database interactions
