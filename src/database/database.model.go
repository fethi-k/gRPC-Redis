package database

type Database interface {
	RecordDB(key string, value string) (string, error)
	GetRecordDB(key string) (string, error)
}

func Factory(databaseName string) (Database, error) {
	switch databaseName {
	case "redis":
		return createRedisDatabase()
	default:
		return nil, &NotImplementedDatabaseError{databaseName}
	}
}
