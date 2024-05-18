package assets

import (
	"embed"
	"io/fs"
	"path/filepath"
)

var (
	// go:embed static/*
	assets embed.FS

	// go:embed logo.png
	LogoData []byte
)

func Assets() fs.FS {
	return assets
}

func Asset(name string) ([]byte, error) {
	return assets.ReadFile(name)
}

func AssetNames() []string {
	names := make([]string, 0)
	err := fs.WalkDir(assets, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			names = append(names, path)
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return names
}

func AssetFile(name string) (fs.File, error) {
	return assets.Open(name)
}

func AssetDir(name string) ([]string, error) {
	dir := filepath.Dir(name)
	names := make([]string, 0)
	err := fs.WalkDir(assets, dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			names = append(names, path)
		}
		return nil
	})

	return names, err
}

func AssetExists(name string) bool {
	_, err := Asset(name)
	return err == nil
}
