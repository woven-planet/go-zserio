package main

import (
	"fmt"
	"log"

	zserio "github.com/woven-planet/go-zserio"
	"myproject.home/zserio-example/contacts"
)

//go:generate go run github.com/woven-planet/go-zserio/cmd/go-zserio generate --rootpackage myproject.home/zserio-example --out . ./schema

func main() {
	address := contacts.Address{Street: "Mainstreet"}
	bytes, err := zserio.Marshal(&address)
	if err != nil {
		log.Fatalf("error serializing address: %w", err)
	}

	fmt.Print(bytes)
}
