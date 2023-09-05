package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"text/template"

	"gopkg.in/yaml.v3"
)

// 类定义结构体
type Class struct {
	Name       string
	Properties map[string]string
}

var config *Config
var wg sync.WaitGroup
var outputDir string

func main() {
	getConfig()

	outputDir = config.FileOutputDirectory
	if config.FileOutputFlag && len(outputDir) > 0 {
		if f, err := os.Stat(outputDir); os.IsNotExist(err) || !f.IsDir() {
			err := os.Mkdir(outputDir, 0700)
			checkError(err)
		}
		outputDir += "/"
	}

	// Read the JSON file
	data, err := getYamlData()
	checkError(err)
	wg.Add(1)
	go generateClass(data, config.RootClassName)

	wg.Wait()
	if config.PauseFlag {
		fmt.Println("请按回车键结束程序。")
		fmt.Scanln() // wait for Enter Key
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal("Failed to generate class definition:", err)
	}
}

func generateClass(data map[string]interface{}, name string) {
	defer wg.Done()
	properties, err := getPropertyMap(data)
	if err != nil {
		fmt.Println(name + "类生成失败。错误信息：" + err.Error())
		return
	}

	class := Class{
		Name:       name,
		Properties: properties,
	}

	err = excuteOutput(class)
	if err != nil {
		return
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
		outputFile := outputDir + class.Name + ".java"
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
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, class)
		if err != nil {
			return err
		}
		fmt.Println(buf.String())
	}

	return nil
}

func getYamlData() (map[string]interface{}, error) {
	// Get JSON string from JSON file
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

	return data, nil
}

// Read json file and get property map (json key -> json type)
func getPropertyMap(data map[string]interface{}) (map[string]string, error) {
	// Get the keys and types
	result := make(map[string]string)
	for key, value := range data {
		javaKey := snakeToCamel(key)
		valueType := fmt.Sprintf("%T", value)
		javaType := getJavaType(valueType, javaKey)
		result[javaKey] = javaType

		if valueType == "map[string]interface {}" {
			wg.Add(1)
			go generateClass(value.(map[string]interface{}), capitalizeFirst(javaKey))
		}

		if valueType == "[]interface {}" && len(value.([]interface{})) > 0 {
			arrValue := value.([]interface{})
			childType := fmt.Sprintf("%T", arrValue[0])
			if childType == "map[string]interface {}" {
				wg.Add(1)
				go generateClass(arrValue[0].(map[string]interface{}), capitalizeFirst(javaKey))
			}
		}
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
