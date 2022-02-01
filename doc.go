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
// serialization code. This can be done with the zserio command. To convert a
// zserio schema defined in a set of files in the "schema" directory you can
// use the following command:
//
//   zserio generate --out ./nds -r github.com/acme/myproject ./schema
//
// This will create a number of code files in the "nds" directory. The root
// package is set to "github.com/acme/myproject", and must be defined to import
// statements can be generated.
//
// See also
//
// https://zserio.org
package main
