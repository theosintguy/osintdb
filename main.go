package main

import (
	_ "embed"
	"github.com/wailsapp/wails"
	"os"
	"osintdb/internal/db"
	"osintdb/internal/stores"
	"osintdb/internal/utils"
	"path/filepath"
)

func quit() {
	os.Exit(0)
}

//go:embed frontend/public/build/bundle.js
var js string

//go:embed frontend/public/build/bundle.css
var css string


func main() {
	bootstrap()
	database := db.New()
	defer database.Close()

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "osintDB",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	dataStores := stores.New()
	database.RegisterStores(dataStores)
	app.Bind(dataStores.Table)

	app.Bind(database)
	app.Bind(quit)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}

func bootstrap() {
	// find path to osintdb config and store directory
	osintdbDir := os.Getenv("OSINTDB_DIR")
	if osintdbDir == "" {
		homeDir := os.Getenv("HOME")
		if _, err := os.Stat(homeDir); err != nil {
			panic("Could not find the environment variable 'HOME'")
		}
		osintdbDir = filepath.Join(homeDir, ".local", "osintdb")
	}

	// create and bootstrap directory if it does not exist already
	err := os.MkdirAll(osintdbDir, 0755)
	if err != nil {
		panic(err)
	}
	if _, err = os.Stat(filepath.Join(osintdbDir, "objectbox")); err != nil {
		err = utils.CopyFile("/opt/osintdb/tools.min.json", filepath.Join(osintdbDir, "tools.min.json"))
		if err != nil {
			panic(err)
		}
	}
}
