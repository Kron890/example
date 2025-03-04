package config

// копируем данные в структуру
type Config struct {
	// Fill struct from parse file, never put here config file directly
	SomeField string `yaml:"some-field"`
}

func GetConfig(path string) (*Config, error) {
	//file, err := os.ReadFile(path)
	//if err != nil {
	//	return nil, err
	//}

	parsedConfig := string("")

	return &Config{
		SomeField: parsedConfig,
	}, nil
}
