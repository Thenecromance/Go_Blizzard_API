package utils

import "sync"

type RequestMethod struct {
	Methods  string
	Path     string
	CnRegion bool
}

type CNRequestMethod struct {
	Name      string
	RealmSlug string
}

// =================================================================================
// Global API Configuration Registry (全局 API 配置注册表)
// =================================================================================

var (
	// apiConfigRegistry internal map to store API metadata globally.
	apiConfigRegistry = make(map[string]RequestMethod)
	registryLock      sync.RWMutex
)

// LoadApiConfigs loads a map of API configurations into the global registry.
// Call this inside your main.go or init() with your complete JSON data loaded into the map.
func LoadApiConfigs(configs map[string]RequestMethod) {
	registryLock.Lock()
	defer registryLock.Unlock()
	for k, v := range configs {
		apiConfigRegistry[k] = v
	}
}

// GetApiConfig retrieves the RequestMethod configuration for a given API name.
// It returns a copy of the struct to ensure thread safety.
// If the config is not found, it returns an empty struct (Path == "").
func GetApiConfig(name string) RequestMethod {
	registryLock.RLock()
	defer registryLock.RUnlock()
	return apiConfigRegistry[name]
}
