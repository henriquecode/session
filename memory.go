package session

// MemoryManager é um struct que representa gerenciamento por
// memoria
type MemoryManager struct {
	settings  DriverMapSetting
	sessionID string
	Keys      []ManagerKeys
}

func newMemory(settings DriverMapSetting) *MemoryManager {
	memory := MemoryManager{
		settings: settings,
	}
	return &memory
}

func (r *MemoryManager) start(sessionID string) {
	r.sessionID = sessionID
}

func (r *MemoryManager) destroy() {
	r.Keys = make([]ManagerKeys, 0)
	r.sessionID = ""
}

func (r *MemoryManager) add(key interface{}, value interface{}) {
	r.Keys = append(r.Keys, ManagerKeys{
		Key:   key,
		Value: value,
	})
}

func (r *MemoryManager) all() []ManagerKeys {
	return r.Keys
}

func (r *MemoryManager) get(key interface{}) ManagerKeys {

	for k, v := range r.Keys {
		if k == key {
			return v
		}
	}

	return ManagerKeys{}
}

func (r *MemoryManager) id() string {
	return r.sessionID
}
