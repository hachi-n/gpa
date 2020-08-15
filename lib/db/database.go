package db

import (
	"encoding/json"
	"fmt"
	"github.com/hachi-n/pag/lib/file"
	"os"
	"path/filepath"
	"strings"
)

const (
	databaseFile = ".config/pag/database.json"
)

type Database struct {
	Path    string
	Records []Record
}

type Record struct {
	ProjectType string
	Tree        interface{}
}

func NewDatabase() *Database {
	databaseFilePath := databaseFilePath()
	return &Database{
		Path: databaseFilePath,
	}
}

func LoadDatabase() (*Database, error) {
	database := NewDatabase()

	var err error
	database.Records, err = loadRecords(database.Path)
	if err != nil {
		return nil, err
	}

	return database, nil
}

func loadRecords(configPath string) ([]Record, error) {
	b, err := file.Read(configPath)
	if err != nil {
		return nil, err
	}

	var configJson interface{}
	if err := json.Unmarshal(b, &configJson); err != nil {
		return nil, err
	}

	configJsonMap, ok := configJson.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("config type assertion error")
	}

	var records []Record
	for projectType, tree := range configJsonMap {
		var configRecord Record
		configRecord.ProjectType, configRecord.Tree = projectType, tree
		records = append(records, configRecord)
	}
	return records, nil
}

func databaseFilePath() string {
	return filepath.Join(os.Getenv("HOME"), databaseFile)
}

func generateJsonBytes(i interface{}) ([]byte, error) {
	repeatSpaceNum := 4
	return json.MarshalIndent(i, "", strings.Repeat(" ", repeatSpaceNum))
}

func (d *Database) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	for _, record := range d.Records {
		m[record.ProjectType] = record.Tree
	}
	return json.Marshal(m)

}

func (d *Database) Update(configPath string) error {
	records, err := loadRecords(configPath)
	if err != nil {
		return err
	}

	d.Records = append(d.Records, records...)

	if err := d.writeDatabase(); err != nil {
		return err
	}

	return nil
}

func (d *Database) writeDatabase() error {
	databaseDir := filepath.Dir(d.Path)
	if _, err := os.Stat(databaseDir); err != nil {
		if err := os.MkdirAll(databaseDir, 0755); err != nil {
			return err
		}
	}

	b, err := generateJsonBytes(d)
	if err != nil {
		return err
	}

	if err := file.Create(d.Path, b); err != nil {
		return err
	}
	return nil
}

func (d *Database) List() ([]string, error) {
	records, err := loadRecords(d.Path)
	if err != nil {
		return nil, err
	}

	var projectTypes []string
	for _, record := range records {
		projectTypes = append(projectTypes, record.ProjectType)
	}

	return projectTypes, nil
}

func (d *Database) Show(projectType string) (string, error) {
	records, err := loadRecords(d.Path)
	if err != nil {
		return "", err
	}

	var jsonBytes []byte
	for _, record := range records {
		if record.ProjectType != projectType {
			continue
		}

		jsonBytes, err = generateJsonBytes(record.Tree)
		if err != nil {
			return "", err
		}
	}

	return string(jsonBytes), nil
}
