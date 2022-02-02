package main

import (
	"fmt"
	"os"

	"github.com/icza/bitio"
	"myprojects.home/zserio-example/contacts"
)

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
