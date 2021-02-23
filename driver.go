package session

// Driver interface usada para todas as implementações
// de drivers feita
type Driver interface {
	start()
	destroy()
	add(key interface{}, value interface{})
	all()
	get(key interface{})
}

// Manager um struct que representa o gerenciamento de sessão
// dentro dos drivers
type ResourceManager struct {
	settings DriverMapSetting
	sessionID string
	keys ManagerKeys
}

// DriverMapSetting tipo responsável por receber configurações
// fornecidade pelo desenvolvedor nas implementações dos driver no formato chave e valor
//
// OBS.: Esse tipo será implementado nos drivers com suas respectivas structs na propriedade "settings", exemplo:
//
//	type FileSystem struct {
//		settings DriverMapSetting
//	}
//
type DriverMapSetting map[string]interface{}

type ManagerKeys map[interface{}]interface{}

// NewDriver factory responsável por fornecer a instancia de driver desejada
func NewDriver(driver string, settings DriverMapSetting) *ResourceManager {

	switch driver {
	//case "filesystem":
	//	return NewFileSystem(settings)
	case "memory":
		new := newMemory(settings)
		return &new
	}

	panic("driver not found")

	return nil
}