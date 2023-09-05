package main

type Config struct {
	JsonFile            string `yaml:"json-file"`
	TemplateFile        string `yaml:"template-file"`
	DefaultType         string `yaml:"default-type"`
	RootClassName       string `yaml:"root-class-name"`
	FileOutputFlag      bool   `yaml:"file-output-flag"`
	FileOutputDirectory string `yaml:"file-output-directory"`
	PauseFlag           bool   `yaml:"pause-flag"`
}

// Initialize a Config struct
func newConfig() *Config {
	return &Config{
		JsonFile:            "sample.json",
		TemplateFile:        "template.txt",
		DefaultType:         "String",
		RootClassName:       "MyClass",
		FileOutputFlag:      false,
		FileOutputDirectory: "output",
		PauseFlag:           true,
	}
}
