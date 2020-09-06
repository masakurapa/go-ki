package ki_test

import (
	"bytes"
	"testing"

	"github.com/masakurapa/gooki/internal/ki"
	"github.com/masakurapa/gooki/pkg/gooki"
)

func TestKi_Write(t *testing.T) {
	path := "../../testdata"

	type args struct {
		option gooki.Option
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "デフォルトオプションの場合",
			args: args{
				option: gooki.DefaultOption(),
			},
			want: `../../testdata
├── example.txt
├── path1
│   ├── to.txt
│   ├── to1
│   │   ├── file1.go
│   │   └── file2.go
│   └── to2
│       └── sample.txt
├── path2
│   ├── example1.go
│   └── example1_test.go
└── symlink.txt -> example.txt

4 directories, 8 files
`,
		},
		{
			name: "AllFile = trueの場合",
			args: args{
				option: func() gooki.Option {
					o := gooki.DefaultOption()
					o.AllFile = true
					return o
				}(),
			},
			want: `../../testdata
├── .hidden
│   └── test.txt
├── .hidden1
├── example.txt
├── path1
│   ├── to.txt
│   ├── to1
│   │   ├── file1.go
│   │   └── file2.go
│   └── to2
│       ├── .hidden2
│       └── sample.txt
├── path2
│   ├── example1.go
│   └── example1_test.go
└── symlink.txt -> example.txt

5 directories, 11 files
`,
		},
		{
			name: "DirectoryOnly = trueの場合",
			args: args{
				option: func() gooki.Option {
					o := gooki.DefaultOption()
					o.DirectoryOnly = true
					return o
				}(),
			},
			want: `../../testdata
├── path1
│   ├── to1
│   └── to2
└── path2

4 directories
`,
		},
		{
			name: "ShowFullPath = trueの場合",
			args: args{
				option: func() gooki.Option {
					o := gooki.DefaultOption()
					o.ShowFullPath = true
					return o
				}(),
			},
			want: `../../testdata
├── ../../testdata/example.txt
├── ../../testdata/path1
│   ├── ../../testdata/path1/to.txt
│   ├── ../../testdata/path1/to1
│   │   ├── ../../testdata/path1/to1/file1.go
│   │   └── ../../testdata/path1/to1/file2.go
│   └── ../../testdata/path1/to2
│       └── ../../testdata/path1/to2/sample.txt
├── ../../testdata/path2
│   ├── ../../testdata/path2/example1.go
│   └── ../../testdata/path2/example1_test.go
└── ../../testdata/symlink.txt -> example.txt

4 directories, 8 files
`,
		},
		{
			name: "ModuleOption.IgnoreTest = trueの場合",
			args: args{
				option: func() gooki.Option {
					o := gooki.DefaultOption()
					o.ModuleOption.IgnoreTest = true
					return o
				}(),
			},
			want: `../../testdata
├── example.txt
├── path1
│   ├── to.txt
│   ├── to1
│   │   ├── file1.go
│   │   └── file2.go
│   └── to2
│       └── sample.txt
├── path2
│   └── example1.go
└── symlink.txt -> example.txt

4 directories, 7 files
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ki.Make(path, tt.args.option)
			if err != nil {
				t.Fatalf("Make() error = %v", err)
			}

			buf := new(bytes.Buffer)
			if err := got.Write(buf); err != nil {
				t.Fatalf("Write() error = %v", err)
			}
			if g := buf.String(); tt.want != g {
				t.Errorf("Write() output want = %v, got = %v", tt.want, g)
			}
		})
	}
}
