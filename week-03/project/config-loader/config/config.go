package config

// ConfigLoader is the main interface for loading configuration
type ConfigLoader interface {
	// Load loads the configuration from a source
	Load() error
	
	// Get retrieves a configuration value by key
	Get(key string) (interface{}, bool)
	
	// GetString retrieves a string configuration value
	GetString(key string) (string, bool)
	
	// GetInt retrieves an integer configuration value
	GetInt(key string) (int, bool)
	
	// GetBool retrieves a boolean configuration value
	GetBool(key string) (bool, bool)
}

// BaseConfig implements common functionality for all config types
type BaseConfig struct {
	// Data stores the configuration values
	Data map[string]interface{}
}

// NewBaseConfig creates a new base configuration
func NewBaseConfig() *BaseConfig {
	return &BaseConfig{
		Data: make(map[string]interface{}),
	}
}

// Get retrieves a value from the configuration
func (c *BaseConfig) Get(key string) (interface{}, bool) {
	value, exists := c.Data[key]
	return value, exists
}

// GetString retrieves a string value from the configuration
func (c *BaseConfig) GetString(key string) (string, bool) {
	value, exists := c.Get(key)
	if !exists {
		return "", false
	}
	
	strValue, ok := value.(string)
	return strValue, ok
}

// GetInt retrieves an integer value from the configuration
func (c *BaseConfig) GetInt(key string) (int, bool) {
	// TODO: Implement integer retrieval with type assertion
	return 0, false
}

// GetBool retrieves a boolean value from the configuration
func (c *BaseConfig) GetBool(key string) (bool, bool) {
	// TODO: Implement boolean retrieval with type assertion
	return false, false
}

// EnvFileConfig implements ConfigLoader for .env files
type EnvFileConfig struct {
	*BaseConfig
	FilePath string
}

// NewEnvFileConfig creates a new .env file configuration loader
func NewEnvFileConfig(filePath string) *EnvFileConfig {
	return &EnvFileConfig{
		BaseConfig: NewBaseConfig(),
		FilePath:   filePath,
	}
}

// Load loads configuration from a .env file
func (e *EnvFileConfig) Load() error {
	// TODO: Implement .env file loading
	// 1. Open the file
	// 2. Read line by line
	// 3. Parse each line into key/value pairs
	// 4. Store in e.Data map
	return nil
}

// JSONConfig implements ConfigLoader for JSON files
type JSONConfig struct {
	*BaseConfig
	FilePath string
}

// NewJSONConfig creates a new JSON configuration loader
func NewJSONConfig(filePath string) *JSONConfig {
	return &JSONConfig{
		BaseConfig: NewBaseConfig(),
		FilePath:   filePath,
	}
}

// Load loads configuration from a JSON file
func (j *JSONConfig) Load() error {
	// TODO: Implement JSON file loading
	return nil
}

// YAMLConfig implements ConfigLoader for YAML files
type YAMLConfig struct {
	*BaseConfig
	FilePath string
}

// NewYAMLConfig creates a new YAML configuration loader
func NewYAMLConfig(filePath string) *YAMLConfig {
	return &YAMLConfig{
		BaseConfig: NewBaseConfig(),
		FilePath:   filePath,
	}
}

// Load loads configuration from a YAML file
func (y *YAMLConfig) Load() error {
	// TODO: Implement YAML file loading
	return nil
}
