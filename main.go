package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

// 类定义结构体
type Class struct {
	Name       string
	Properties map[string]string
}

var config *Config

func main() {
	getConfig()

	// Read the JSON file
	properties, err := getPropertyMap()
	checkError(err)

	// Define a struct for parsing the template
	class := Class{
		Name:       config.ClassName,
		Properties: properties,
	}

	err = excuteOutput(class)
	checkError(err)

	if config.PauseFlag {
		fmt.Println("\n请按回车键结束程序。")
		fmt.Scanln() // wait for Enter Key
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Failed to generate class definition:", err)
	}
}

func excuteOutput(class Class) error {
	// Parse the template
	tmpl, err := template.ParseFiles(config.TemplateFile)
	if err != nil {
		return err
	}

	// 根据config.FileOutputFlag决定要将生成的文本输出到文件还是命令行
	if config.FileOutputFlag {
		outputFile := config.ClassName + ".java"
		file, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer file.Close()

		err = tmpl.Execute(file, class)
		if err != nil {
			return err
		}
		fmt.Println("已将生成内容输出到" + outputFile + "。")
	} else {
		err = tmpl.Execute(os.Stdout, class)
		if err != nil {
			return err
		}
	}

	return nil
}

// Read json file and get property map (json key -> json type)
func getPropertyMap() (map[string]string, error) {
	// Define the JSON string
	jsonString, err := os.ReadFile(config.JsonFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	// Unmarshal the JSON data into a map
	var data map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	// Get the keys and types
	result := make(map[string]string)
	for key, value := range data {
		javaKey := snakeToCamel(key)
		valueType := fmt.Sprintf("%T", value)
		javaType := getJavaType(valueType, javaKey)
		result[javaKey] = javaType
	}

	return result, nil
}

// Get the Java type for a given GO type
func getJavaType(goType string, javaKey string) string {
	switch goType {
	case "string":
		return "String"
	case "float64":
		return "double"
	case "bool":
		return "boolean"
	case "[]interface {}":
		return fmt.Sprintf("List<%s>", capitalizeFirst(javaKey))
	case "map[string]interface {}":
		return capitalizeFirst(javaKey)
	default:
		return config.DefaultType
	}
}

// Read config.yml
func getConfig() {
	config = newConfig()

	// Read the YAML file
	yamlFile, err := os.ReadFile("config.yml")
	if err != nil {
		log.Println("Failed to read config.yml:", err)
	} else {
		// Parse the YAML data into a Config struct
		err = yaml.Unmarshal(yamlFile, config)
		if err != nil {
			log.Println("Failed to parse YAML data:", err)
		}
	}
}
