package gooki_test

import (
	"os"

	"github.com/masakurapa/gooki"
)

func ExampleMake() {
	ki, err := gooki.Make(".", gooki.DefaultOption())
	if err != nil {
		panic(err)
	}
	ki.Write(os.Stdout)
	// Output:
	// .
	// └── path
	//     └── to
	//         └── file.go
}

func ExampleMakeByDefaultOption() {
	ki, err := gooki.MakeByDefaultOption(".")
	if err != nil {
		panic(err)
	}
	ki.Write(os.Stdout)
	// Output:
	// .
	// └── path
	//     └── to
	//         └── file.go
}
