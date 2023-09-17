package main

import (
	"context"
	"embed"
	"log"

	launcher "mtgolauncher/backend/Launcher"
	websocket "mtgolauncher/backend/Websocket"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	online := launcher.NewOnline()
	UI := launcher.NewUI()
	Storage := launcher.NewStorage()
	Config := launcher.NewConfig()
	Mod := launcher.NewMod()
	App := launcher.NewApp()
	AKI := launcher.NewAKI()
	MTGA := launcher.NewMTGA()
	Download := launcher.NewDownload()
	websocketManager := &websocket.WebSocketManager{}
	websocketManager.InitWebSocket()

	go websocketManager.StartServer()

	// Create application with options
	err := wails.Run(&options.App{
		Title:         "MTGO Launcher",
		Width:         1020,
		Height:        662,
		DisableResize: true,
		Frameless:     true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			App.Startup(ctx)
			AKI.Startup(ctx)

		},
		OnShutdown: func(ctx context.Context) {
			websocketManager.ShutdownServer()
		},
		Bind: []interface{}{
			app,
			online,
			UI,
			Storage,
			Config,
			Mod,
			App,
			AKI,
			MTGA,
			websocketManager,
			Download,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:              true,
			WindowIsTranslucent:               true,
			BackdropType:                      windows.Tabbed,
			DisableWindowIcon:                 false,
			DisableFramelessWindowDecorations: false,
			IsZoomControlEnabled:              false,
			WebviewUserDataPath:               "",
			WebviewBrowserPath:                "",
			Theme:                             windows.SystemDefault,
			CustomTheme: &windows.ThemeSettings{
				DarkModeTitleBar:   windows.RGB(20, 20, 20),
				DarkModeTitleText:  windows.RGB(200, 200, 200),
				DarkModeBorder:     windows.RGB(20, 0, 20),
				LightModeTitleBar:  windows.RGB(200, 200, 200),
				LightModeTitleText: windows.RGB(20, 20, 20),
				LightModeBorder:    windows.RGB(200, 200, 200),
			},
			Messages:  &windows.Messages{},
			OnSuspend: func() {},
			OnResume:  func() {},
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "MTGO Launcher",
				Message: "© 2021 Make Tarkov Great Again",
			},
		},
		Linux: &linux.Options{
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
