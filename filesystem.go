package session

import "os"

// FileSystem é um driver para controle de ssesão
// Através de arquivos locais

// FileSystem um struct que recebe settings do tipo DriverMapSetting
type FileSystem struct {
	settings DriverMapSetting
}

var filesystem *FileSystem

func NewFileSystem(settings DriverMapSetting) *FileSystem {
	filesystem := &FileSystem{
		settings: settings,
	}
	return filesystem
}

func(f* FileSystem) start() {
	// ...
}

func(f* FileSystem) createSession() {
	_, err := os.Stat(filesystem.settings["path"].(string))

	if err != nil {
		// criar arquivo ...
	}
}