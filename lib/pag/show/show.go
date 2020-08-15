package show

import (
	"fmt"
	"github.com/hachi-n/pag/lib/db"
)

func Apply(projectType string) error {
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

	fmt.Println(tree)
	return nil
}
