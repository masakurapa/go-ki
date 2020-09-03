# gooki
list contents of directories in a tree-like format.

## Module

### Usage

```
import "github.com/masakurapa/gooki"

func main() {
	gooki.Make(".", gooki.DefaultOption())
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
│       ├── main.go
│       └── option.go
├── go.mod
├── internal
│   ├── dirinfo
│   │   └── dirinfo.go
│   ├── fileinfo
│   │   └── fileinfo.go
│   └── tree
│       ├── tree.go
│       └── write.go
└── pkg
    └── option
        └── option.go

8 directories, 11 files
```
