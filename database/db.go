package database

type dbHelper struct {
	db *sql.database
}

func New () IDatabase {
	return &dbHelper{}
}