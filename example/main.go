package main

import (
	"fmt"
	"os"

	"github.com/icza/bitio"
	"myproject.home/zserio-example/contacts"
)

//go:generate go run github.com/woven-planet/go-zserio/cmd/go-zserio generate --rootpackage myproject.home/zserio-example --out . ./schema

func main() {
	address := contacts.Address{Street: "Mainstreet"}
	writer := bitio.NewCountWriter(os.Stdout)
	if err := address.MarshalZserio(writer); err != nil {
		panic(fmt.Sprintf("error serializing address: %v", err))
	}
	if err := writer.Close(); err != nil {
		panic(fmt.Sprintf("error closing writer: %v", err))
	}
}
