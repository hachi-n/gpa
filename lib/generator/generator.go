package generator

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/pag/lib/file"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GenerateProjectArchitecture(projectName, output, tree string) error {
	var treeJson interface{}
	if err := json.Unmarshal([]byte(tree), &treeJson); err != nil {
		return err
	}
	projectGenerateFunc := projectArchitectureGenerator(projectName, output)

	if err := os.MkdirAll(filepath.Join(output, projectName), 0755); err != nil {
		return err
	}
	projectGenerateFunc("", treeJson)

	return nil
}

func projectArchitectureGenerator(projectName, output string) func(string, interface{}) {
	distination := filepath.Join(output, projectName)

	variableMap := defaultVariableMap(projectName)

	fileProcessFunc := func(dist, src string, currentDir string) error {
		fileName := filepath.Join(distination, currentDir, dist)

		var buf []byte
		// Src file exist check.
		if _, err := os.Stat(src); err == nil {
			buf, err = file.Read(src)
			if err != nil {
				return err
			}
			fmt.Println("fileLoad:", src)
		}

		err := file.Create(fileName, buf)
		if err != nil {
			return err
		}
		fmt.Printf("createFile: %s\n", fileName)

		return nil
	}

	directoryProcessFunc := func(dist string, currentDir string) (string, error) {
		dirName := filepath.Join(distination, currentDir, dist)
		if err := os.Mkdir(dirName, 0755); err != nil {
			return currentDir, err
		}
		fmt.Printf("createDirectory: %s\n", dirName)

		currentDir = filepath.Join(currentDir, dist)
		return currentDir, nil
	}

	var projectGenerateFunc func(currentDir string, treeJson interface{})
	projectGenerateFunc = func(currentDir string, treeJson interface{}) {
		if treeJson == nil {
			return
		}
		treeMap := treeJson.(map[string]interface{})
		for key, value := range treeMap {
			key = convertVariableName(key, variableMap)
			// file process.
			if filePath, ok := value.(string); ok {
				err := fileProcessFunc(key, filePath, currentDir)
				if err != nil {
					fmt.Println(err)
				}
				continue
			}

			// directory process
			currentDir, err := directoryProcessFunc(key, currentDir)
			if err != nil {
				fmt.Println(err)
				continue
			}
			projectGenerateFunc(currentDir, value)
		}
	}
	return projectGenerateFunc
}

func defaultVariableMap(projectName string) map[string]string {
	variableMap := make(map[string]string)
	variableMap["ProjectName"] = projectName
	return variableMap
}

var variableRegexp = regexp.MustCompile(`\$\{.*?\}`)

func convertVariableName(baseName string, variableMap map[string]string) string {
	if !variableRegexp.MatchString(baseName) {
		return baseName
	}

	var replacedName string
	for key, value := range variableMap {
		variableString := fmt.Sprintf("${%s}", key)
		replacedName = strings.Replace(baseName, variableString, value, -1)

	}
	return replacedName
}
