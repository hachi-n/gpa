package add

import "github.com/hachi-n/pag/lib/db"

func Apply() error {
	database := db.NewDatabase()
	return database.Destroy()
}
