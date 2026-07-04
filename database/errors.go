package database

import "fmt"

type DBerror struct {
	message string
	cause   error
}

func (d *DBerror) Error() string {
	return fmt.Sprintf("DATABASE ERROR : %v | %v", d.message, d.cause)
}

func (d *DBerror) Unwrap() error {
	return d.cause
}
