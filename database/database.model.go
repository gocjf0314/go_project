package database

type Database interface {
	Set(key string, value string) (string, error)
	Get(key string) (string, error)
	Delete(key string) (string, error)
}

func Factory(databasename string) (Database, error) {
	switch databasename {
	case "redis":
		return InitializeRedis()
	default:
		return nil, &NotImplementedError{databasename}
	}
}
