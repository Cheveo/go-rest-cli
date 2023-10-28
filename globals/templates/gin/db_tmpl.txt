package db

type DB interface {
	CreateDB(databaseUrl string) any
}
