package main

import (
	"fmt"
	"os"

	"github.com/woven-planet/go-zserio/ztype"
	"myprojects.home/zserio-example/contacts"
)

func main() {
	writer := ztype.NewWriter(os.Stdout)
	address := contacts.Address{Street: "Mainstreet"}
	if err := address.MarshalZserio(writer); err != nil {
		panic(fmt.Sprintf("error serializing address: %v", err))
	}
	if err := writer.Close(); err != nil {
		panic(fmt.Sprintf("error flushing buffer address: %v", err))
	}
}
