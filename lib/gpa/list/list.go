package list

import (
	"fmt"
	"github.com/hachi-n/gpa/lib/db"
)

func Apply() error {
	database, err := db.LoadDatabase()
	if err != nil {
		return err
	}
	list, err := database.List()
	if err != nil {
		return err
	}

	for _, v := range list {
		fmt.Println(v)
	}
	return nil
}
