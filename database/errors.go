package database

type OperationError struct {
	operation string
}

func (err *OperationError) Error() string {
	return "Could not perform the " + err.operation + " operation"
}

type DownError struct{}

func (dbe *DownError) Error() string {
	return "Database is down"
}

type CreateDatabaseError struct{}

func (err *CreateDatabaseError) Error() string {
	return "Could not create database"
}

type NotImplementedError struct {
	database string
}

func (err *NotImplementedError) Error() string {
	return err.database + "is not implememted"
}
