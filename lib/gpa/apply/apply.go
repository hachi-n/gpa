package apply

import (
	"fmt"
	"github.com/hachi-n/gpa/lib/db"
	"github.com/hachi-n/gpa/lib/generator"
)

func Apply(projectName, output, projectType string) error {
	database, err := db.LoadDatabase()
	if err != nil {
		return err
	}

	tree, err := database.Show(projectType)
	if err != nil {
		return err
	}
	if tree == "" {
		return fmt.Errorf("Record Not Found.")
	}

	if err := generator.GenerateProjectArchitecture(projectName, output, tree); err != nil {
		return err
	}

	return nil
}
