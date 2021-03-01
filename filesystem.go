package session

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Filesystem é um struct que representa gerenciamento
// atravces de arquivos
type Filesystem struct {
	Settings  DriverMapSetting
	SessionID string
	Keys      []ManagerKeys
}

var fileSession file
var filesystem Filesystem

func newFilesystem(settings DriverMapSetting) *Filesystem {
	filesystem = Filesystem{
		Settings: settings,
	}
	return &filesystem
}

func (r *Filesystem) start(sessionID string) {

	r.SessionID = sessionID

	fileSession = file{
		name:     "session___" + sessionID,
		path:     "files_session/",
		fullpath: "files_session/" + "session___" + sessionID,
	}

	fileSession.create()
	fileSession.store()
}

func (r *Filesystem) destroy() {

	fileSession.delete()

	r.Keys = make([]ManagerKeys, 0)
	r.SessionID = ""
}

func (r *Filesystem) add(key interface{}, value interface{}) {
	r.Keys = append(r.Keys, ManagerKeys{
		Key:   key,
		Value: value,
	})

	fileSession.store()
}

func (r *Filesystem) all() []ManagerKeys {

	fileSession.read()

	return r.Keys
}

func (r *Filesystem) get(key interface{}) ManagerKeys {

	fileSession.read()

	for k, v := range r.Keys {
		if r.Keys[k].Key == key {
			return v
		}
	}

	return ManagerKeys{}
}

func (r *Filesystem) id() string {
	return r.SessionID
}

// File criado para gerenciar arquivos da sessao
type file struct {
	name     string
	path     string
	fullpath string
}

func (f *file) create() {

	if _, err := os.Stat(f.path); err != nil {

		if err := os.Mkdir(f.path, 0755); err != nil {
			panic(err)
		}
	}

	err := ioutil.WriteFile(f.fullpath, []byte(""), 0644)

	if err != nil {
		log.Fatal(err)
	}
}

func (f *file) read() {
	var dataSession = Filesystem{}

	data, err := ioutil.ReadFile(f.fullpath)

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &dataSession); err != nil {
		log.Fatal(err)
	}

	filesystem = dataSession
}

func (f *file) store() {
	dataJSON, err := json.Marshal(filesystem)

	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(f.fullpath, dataJSON, 0644); err != nil {
		log.Fatal(err)
	}
}

func (f *file) delete() {
	_ = os.Remove(f.fullpath)
}
