package main

import (
	"github.com/golang-migrate/migrate"
)

func main() {

	m, err := migrate.New("", "")
	if err != nil {

	}

	m.Up()

}
