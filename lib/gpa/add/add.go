package add

import "github.com/hachi-n/gpa/lib/db"

func Apply(configPath string) error {
	database, err := db.LoadDatabase()
	if err != nil {
		return err
	}
	if err := database.Update(configPath); err != nil {
		return err
	}
	return nil
}
