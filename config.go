package main

type Config struct {
	JsonFile       string `yaml:"json-file"`
	TemplateFile   string `yaml:"template-file"`
	DefaultType    string `yaml:"default-type"`
	ClassName      string `yaml:"class-name"`
	FileOutputFlag bool   `yaml:"file-output-flag"`
	PauseFlag      bool   `yaml:"pause-flag"`
}

// Initialize a Config struct
func newConfig() *Config {
	return &Config{
		JsonFile:       "sample.json",
		TemplateFile:   "template.txt",
		DefaultType:    "String",
		ClassName:      "MyClass",
		FileOutputFlag: false,
		PauseFlag:      true,
	}
}
