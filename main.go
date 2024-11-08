package main

import (
	"AutoCode/openAPI"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	api := openAPI.NewOpenAPI()

	// Create application with options
	err := wails.Run(&options.App{
		Title: "AutoCode",
		//Width:  1024,
		//Height: 768,
		Width:  1600,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		BackgroundColour: &options.RGBA{R: 235, G: 235, B: 235, A: 1},
		OnStartup:        app.startup,
		DisableResize:    true,
		Bind: []interface{}{
			app,
			api,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
