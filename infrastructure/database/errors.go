package database

import "fmt"

type DBerror struct {
	message string
	cause   error
}

func (d *DBerror) Error() string {
	return fmt.Sprintf("DATABASE ERROR : %v | %v", d.message, d.cause)
}


// The Unwrap method is used in ideomatic go when a custom error type wraps a existing error 
// The Unwrap method used to identify the underlying error desipte of being wrapped by a custom error. 
func (d *DBerror) Unwrap() error {
	return d.cause
}
