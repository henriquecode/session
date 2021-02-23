package session

// Memory é um driver para controle de sessão
// Através da memória

var memory ResourceManager

func newMemory(settings DriverMapSetting) ResourceManager {
	memory = ResourceManager{
		settings: settings,
	}
	return memory
}

func(f* ResourceManager) start() {
	// ..
}

func(f* ResourceManager) destroy() {
	f.keys = make(ManagerKeys)
	f.sessionID = ""
}

func(f* ResourceManager) add(key interface{}, value interface{}) {
	f.keys[key] = value
}

func(f* ResourceManager) all() {
	// ..
}

func(f* ResourceManager) get(key interface{}) interface{} {
	return f.keys[key]
}