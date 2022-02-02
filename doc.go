// zserio is a serialization framework designed for environments where
// storage or bandwidth is the most critical constraint. The primary use
// case is the NDS.Live format for automobile map data.
//
// This package implements the de(serialization) of the core zserio data
// types, as well as a code generator to create Go code from zserio
// schema definitions.
//
// Installation
//
// You can use the normal go install method to install go-zserio:
//
//    go install github.com/woven-planet/go-zserio/zserio@latest
//
// Code generation
//
// The most common, and strongly preferred, method to use zserio is to write
// a zserio schema, and then generate code to define the data structures and
// serialization code.
//
// As an example let's look at a zserio schema for a contact list. In a minimal
// example we only define an Addreses structure to store addres information.
//
//     package contacts;
//
//     // Address, so I know where your house lives
//     struct Address
//     {
//         string street;
//         optional uint8 number;
//     };
//
// After saving this in schema/contacts.zs you can generate Go code for this
// schema with this command:
//
//   zserio generate --out contacts -r myprojects.home/zserio-example/addressbook ./schema
//
// This will create a number of code files in the "contacts" directory. The root
// package is set to "myprojects.home/zserio-example/addressbook", and must be
// defined to import statements can be generated.
//
// Can you now import the generated code, and serialize Address records using zserio:
//
//     package main
//
//     import (
//         "fmt"
//         "os"
//
//         "github.com/icza/bitio"
//         "myprojects.home/zserio-example/contacts"
//     )
//
//     func main() {
//         address := contacts.Address{Street: "Mainstreet"}
//         writer := bitio.NewCountWriter(os.Stdout)
//         if err := address.MarshalZserio(writer); err != nil {
//             panic(fmt.Sprintf("error serializing address: %v", err))
//         }
//         if err := writer.Close(); err != nil {
//             panic(fmt.Sprintf("error closing writer: %v", err))
//         }
//    }
//
// See also
//
// https://zserio.org
package main
