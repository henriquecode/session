package session

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var currentDriver *ResourceManager
var lock sync.Mutex
var writer http.ResponseWriter
var sessionName string = "applicationGoSession"

func NewSession(driver string, settings DriverMapSetting) *ResourceManager {

	currentDriver = NewDriver(driver, settings)
	return currentDriver
}

func SetNameSession(name string) {
	sessionName = name
}

func Start() {
	lock.Lock()

	currentDriver.keys = make(ManagerKeys)

	if Id() == "" {
		currentDriver.sessionID = createSessionID()
	}

	currentDriver.start()

	lock.Unlock()
}

func Destroy() {
	lock.Lock()
	currentDriver.destroy()
	lock.Unlock()
}

func Id() string {
	return currentDriver.sessionID
}

func Push(key interface{}, value interface{}) {
	lock.Lock()
	currentDriver.add(key, value)
	lock.Unlock()
}

func All() {
	lock.Lock()
	currentDriver.all()
	lock.Unlock()
}

func Get(key interface{}) interface{} {
	return currentDriver.get(key)
}

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
		Value: Id(),
	})
}
