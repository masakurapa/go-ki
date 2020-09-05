# gooki
list contents of directories in a tree-like format.

## Module

### Usage

```
import (
	"os"

	"github.com/masakurapa/gooki"
)

func main() {
	ki, _ := gooki.MakeByDefaultOption(".")
	ki.Write(os.Stdout)
}
```

## Command Line Tool

### Installing

```
$ go get -u github.com/masakurapa/gooki/cmd/gooki
```

### Usage

```
$ gooki -help

	  _____             _  ___
	 / ____|           | |/ (_)
	| |  __  ___   ___ | ' / _
	| | |_ |/ _ \ / _ \|  < | |
	| |__| | (_) | (_) | . \| |
	 \_____|\___/ \___/|_|\_\_|

Description:
	'gooki' is list contents of directories in a tree-like format.

Synopsis:
	gooki [options...] directory

Options:
  -a	Outputs all files. By default does not hidden files.
  -d	Outputs only directories.
  -f	Outputs full path prefix for each file.
  -help
    	Outputs a usage.
```

### Example
```
$ gooki
.
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── gooki
│       └── main.go
├── go.mod
├── gooki.go
└── internal
    ├── ki
    │   ├── interface.go
    │   ├── ki.go
    │   └── tree_writer.go
    └── opt
        └── option.go

5 directories, 10 files
```
