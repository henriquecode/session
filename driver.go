package session

// IDriver interface usada para todas as implementações
// de drivers feita
type IDriver interface {
	start(sessionID string)
	destroy()
	add(key interface{}, value interface{})
	all() []ManagerKeys
	get(key interface{}) ManagerKeys
	id() string
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

// ManagerKeys representa um tipo para chaves dentro da sessao
type ManagerKeys struct {
	Key   interface{}
	Value interface{}
}

// NewDriver factory responsável por fornecer a instancia de driver desejada
func NewDriver(driver string, settings DriverMapSetting) IDriver {

	switch driver {
	case "filesystem":
		new := newFilesystem(settings)
		return new
	case "memory":
		new := newMemory(settings)
		return new
	default:
		panic("driver not found")
	}
}
