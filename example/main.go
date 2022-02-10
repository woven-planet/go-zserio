package main

import (
	"fmt"
	"os"

	"github.com/woven-planet/go-zserio/ztype"
	"myprojects.home/zserio-example/contacts"
)

func main() {
	writer := ztype.NewWriter(os.Stdout)
	defer writer.Close()

	address := contacts.Address{Street: "Mainstreet"}
	if err := address.MarshalZserio(writer); err != nil {
		panic(fmt.Sprintf("error serializing address: %v", err))
	}
}
