package web

import (
	"embed"
	"io/fs"
)

//go:embed templates/* static/*
var assets embed.FS

func Templates() fs.FS {
	return assets
}

func Static() (fs.FS, error) {
	return fs.Sub(assets, "static")
}
