package gooki_test

import (
	"os"

	"github.com/masakurapa/gooki"
)

func ExampleMake() {
	ki, err := gooki.Make("./testdata", gooki.DefaultOption())
	if err != nil {
		panic(err)
	}
	ki.Write(os.Stdout)
	// Output:
	// ./testdata
	// ├── example.txt
	// ├── path1
	// │   ├── to.txt
	// │   ├── to1
	// │   │   ├── file1.go
	// │   │   └── file2.go
	// │   └── to2
	// │       └── sample.txt
	// ├── path2
	// │   ├── example1.go
	// │   └── example1_test.go
	// └── symlink.txt -> example.txt
	//
	// 4 directories, 8 files
}

func ExampleMakeByDefaultOption() {
	ki, err := gooki.MakeByDefaultOption("./testdata")
	if err != nil {
		panic(err)
	}
	ki.Write(os.Stdout)
	// Output:
	// ./testdata
	// ├── example.txt
	// ├── path1
	// │   ├── to.txt
	// │   ├── to1
	// │   │   ├── file1.go
	// │   │   └── file2.go
	// │   └── to2
	// │       └── sample.txt
	// ├── path2
	// │   ├── example1.go
	// │   └── example1_test.go
	// └── symlink.txt -> example.txt
	//
	// 4 directories, 8 files
}
