package session

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var currentDriver IDriver
var lock sync.Mutex
var writer http.ResponseWriter
var sessionName string = "applicationGoSession"

// NewSession cria uma nova instancia para sessao de seguintes tipos:
// memory, filesystem
func NewSession(driver string, settings DriverMapSetting) {
	currentDriver = NewDriver(driver, settings)
}

// SetNameSession defini um nome para sessao atual
func SetNameSession(name string) {
	sessionName = name
}

// Start inicia sessao
func Start() {
	lock.Lock()

	sessID := ID()

	if sessID == "" {
		sessID = createSessionID()
	}

	currentDriver.start(sessID)

	lock.Unlock()
}

// Destroy destroi a sessao atual
func Destroy() {
	lock.Lock()
	currentDriver.destroy()
	lock.Unlock()
}

// ID recupera id da sessao
func ID() string {
	return currentDriver.id()
}

// Push adiciona um item na chave expecificada
func Push(key interface{}, value interface{}) {
	lock.Lock()
	currentDriver.add(key, value)
	lock.Unlock()
}

// All recupera todos os itens guardados na sessao
func All() []ManagerKeys {
	return currentDriver.all()
}

// Get recupera um item expecífico da sessao
func Get(key interface{}) ManagerKeys {
	return currentDriver.get(key)
}

// Delete um item da sessao
func Delete(key interface{}) {
	lock.Lock()
	currentDriver.destroy()
	lock.Unlock()
}

func createSessionID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	start := make([]string, 20)
	end := make([]string, 20)

	// generate inicio do id
	for n := range start {
		if n%2 == 0 {
			strStart := string(65 + r.Intn(25))
			strEnd := string(65 + r.Intn(25))

			start[n] = string(strStart)
			end[n] = string(strEnd)
		} else {
			start[n] = strconv.Itoa(r.Intn(10))
			end[n] = strconv.Itoa(r.Intn(10))
		}
	}

	id := strings.Join(start[:], "") + strconv.Itoa(int(r.Int63())) + strings.Join(end[:], "")

	return id
}

// SetWriter precisa ser chamado em todo inicio de request
// que quiser utilizar de sessao, caso contrário não será
// setado no cookie
func SetWriter(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  sessionName,
		Value: ID(),
	})
}
