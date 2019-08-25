package repository

import "fmt"

type NotFoundErr struct {
	Msg string
}

func (nf NotFoundErr) Error() string {
	return fmt.Sprintf("could not found: %s", nf.Msg)
}
