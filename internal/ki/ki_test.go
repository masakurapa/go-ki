package ki_test

import (
	"path/filepath"
	"testing"

	"github.com/masakurapa/gooki/internal/ki"
	"github.com/masakurapa/gooki/pkg/gooki"
)

func TestMake(t *testing.T) {
	path, abs := testPath(t)
	defaultExpected := []expectedEda{
		{
			name: "example.txt",
			ha: expectedHappa{
				absPath: abs + "/example.txt", relPath: "example.txt", dir: ".", name: "example.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
			},
			hasChild: false,
			eda:      []expectedEda{},
		},
		{
			name: "path1",
			ha: expectedHappa{
				absPath: abs + "/path1", relPath: "path1", dir: ".", name: "path1", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
			},
			hasChild: true,
			eda: []expectedEda{
				{
					name: "to.txt",
					ha: expectedHappa{
						absPath: abs + "/path1/to.txt", relPath: "path1/to.txt", dir: "path1", name: "to.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
				{
					name: "to1",
					ha: expectedHappa{
						absPath: abs + "/path1/to1", relPath: "path1/to1", dir: "path1", name: "to1", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: true,
					eda: []expectedEda{
						{
							name: "file1.go",
							ha: expectedHappa{
								absPath: abs + "/path1/to1/file1.go", relPath: "path1/to1/file1.go", dir: "path1/to1", name: "file1.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
						{
							name: "file2.go",
							ha: expectedHappa{
								absPath: abs + "/path1/to1/file2.go", relPath: "path1/to1/file2.go", dir: "path1/to1", name: "file2.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
					},
				},
				{
					name: "to2",
					ha: expectedHappa{
						absPath: abs + "/path1/to2", relPath: "path1/to2", dir: "path1", name: "to2", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: true,
					eda: []expectedEda{
						{
							name: "sample.txt",
							ha: expectedHappa{
								absPath: abs + "/path1/to2/sample.txt", relPath: "path1/to2/sample.txt", dir: "path1/to2", name: "sample.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
					},
				},
			},
		},
		{
			name: "path2",
			ha: expectedHappa{
				absPath: abs + "/path2", relPath: "path2", dir: ".", name: "path2", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
			},
			hasChild: true,
			eda: []expectedEda{
				{
					name: "example1.go",
					ha: expectedHappa{
						absPath: abs + "/path2/example1.go", relPath: "path2/example1.go", dir: "path2", name: "example1.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
				{
					name: "example1_test.go",
					ha: expectedHappa{
						absPath: abs + "/path2/example1_test.go", relPath: "path2/example1_test.go", dir: "path2", name: "example1_test.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
			},
		},
		{
			name: "symlink.txt",
			ha: expectedHappa{
				absPath: abs + "/symlink.txt", relPath: "symlink.txt", dir: ".", name: "symlink.txt", isDir: false, isHiddenFile: false, isSymlink: true, realName: "example.txt",
			},
			hasChild: false,
			eda:      []expectedEda{},
		},
	}

	type args struct {
		option gooki.Option
	}
	tests := []struct {
		name string
		args args
		want []expectedEda
	}{
		{
			name: "デフォルトオプションの場合",
			args: args{
				option: gooki.DefaultOption(),
			},
			want: defaultExpected,
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
			want: []expectedEda{
				{
					name: ".hidden",
					ha: expectedHappa{
						absPath: abs + "/.hidden", relPath: ".hidden", dir: ".", name: ".hidden", isDir: true, isHiddenFile: true, isSymlink: false, realName: "",
					},
					hasChild: true,
					eda: []expectedEda{
						{
							name: "test.txt",
							ha: expectedHappa{
								absPath: abs + "/.hidden/test.txt", relPath: ".hidden/test.txt", dir: ".hidden", name: "test.txt", isDir: false, isHiddenFile: true, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
					},
				},
				{
					name: ".hidden1",
					ha: expectedHappa{
						absPath: abs + "/.hidden1", relPath: ".hidden1", dir: ".", name: ".hidden1", isDir: false, isHiddenFile: true, isSymlink: false, realName: "",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
				{
					name: "example.txt",
					ha: expectedHappa{
						absPath: abs + "/example.txt", relPath: "example.txt", dir: ".", name: "example.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
				{
					name: "path1",
					ha: expectedHappa{
						absPath: abs + "/path1", relPath: "path1", dir: ".", name: "path1", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: true,
					eda: []expectedEda{
						{
							name: "to.txt",
							ha: expectedHappa{
								absPath: abs + "/path1/to.txt", relPath: "path1/to.txt", dir: "path1", name: "to.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
						{
							name: "to1",
							ha: expectedHappa{
								absPath: abs + "/path1/to1", relPath: "path1/to1", dir: "path1", name: "to1", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: true,
							eda: []expectedEda{
								{
									name: "file1.go",
									ha: expectedHappa{
										absPath: abs + "/path1/to1/file1.go", relPath: "path1/to1/file1.go", dir: "path1/to1", name: "file1.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
									},
									hasChild: false,
									eda:      []expectedEda{},
								},
								{
									name: "file2.go",
									ha: expectedHappa{
										absPath: abs + "/path1/to1/file2.go", relPath: "path1/to1/file2.go", dir: "path1/to1", name: "file2.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
									},
									hasChild: false,
									eda:      []expectedEda{},
								},
							},
						},
						{
							name: "to2",
							ha: expectedHappa{
								absPath: abs + "/path1/to2", relPath: "path1/to2", dir: "path1", name: "to2", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: true,
							eda: []expectedEda{
								{
									name: ".hidden2",
									ha: expectedHappa{
										absPath: abs + "/path1/to2/.hidden2", relPath: "path1/to2/.hidden2", dir: "path1/to2", name: ".hidden2", isDir: false, isHiddenFile: true, isSymlink: false, realName: "",
									},
									hasChild: false,
									eda:      []expectedEda{},
								},
								{
									name: "sample.txt",
									ha: expectedHappa{
										absPath: abs + "/path1/to2/sample.txt", relPath: "path1/to2/sample.txt", dir: "path1/to2", name: "sample.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
									},
									hasChild: false,
									eda:      []expectedEda{},
								},
							},
						},
					},
				},
				{
					name: "path2",
					ha: expectedHappa{
						absPath: abs + "/path2", relPath: "path2", dir: ".", name: "path2", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: true,
					eda: []expectedEda{
						{
							name: "example1.go",
							ha: expectedHappa{
								absPath: abs + "/path2/example1.go", relPath: "path2/example1.go", dir: "path2", name: "example1.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
						{
							name: "example1_test.go",
							ha: expectedHappa{
								absPath: abs + "/path2/example1_test.go", relPath: "path2/example1_test.go", dir: "path2", name: "example1_test.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
					},
				},
				{
					name: "symlink.txt",
					ha: expectedHappa{
						absPath: abs + "/symlink.txt", relPath: "symlink.txt", dir: ".", name: "symlink.txt", isDir: false, isHiddenFile: false, isSymlink: true, realName: "example.txt",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
			},
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
			want: []expectedEda{
				{
					name: "path1",
					ha: expectedHappa{
						absPath: abs + "/path1", relPath: "path1", dir: ".", name: "path1", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: true,
					eda: []expectedEda{
						{
							name: "to1",
							ha: expectedHappa{
								absPath: abs + "/path1/to1", relPath: "path1/to1", dir: "path1", name: "to1", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
						{
							name: "to2",
							ha: expectedHappa{
								absPath: abs + "/path1/to2", relPath: "path1/to2", dir: "path1", name: "to2", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
					},
				},
				{
					name: "path2",
					ha: expectedHappa{
						absPath: abs + "/path2", relPath: "path2", dir: ".", name: "path2", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
			},
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
			want: defaultExpected,
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
			want: []expectedEda{
				{
					name: "example.txt",
					ha: expectedHappa{
						absPath: abs + "/example.txt", relPath: "example.txt", dir: ".", name: "example.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
				{
					name: "path1",
					ha: expectedHappa{
						absPath: abs + "/path1", relPath: "path1", dir: ".", name: "path1", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: true,
					eda: []expectedEda{
						{
							name: "to.txt",
							ha: expectedHappa{
								absPath: abs + "/path1/to.txt", relPath: "path1/to.txt", dir: "path1", name: "to.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
						{
							name: "to1",
							ha: expectedHappa{
								absPath: abs + "/path1/to1", relPath: "path1/to1", dir: "path1", name: "to1", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: true,
							eda: []expectedEda{
								{
									name: "file1.go",
									ha: expectedHappa{
										absPath: abs + "/path1/to1/file1.go", relPath: "path1/to1/file1.go", dir: "path1/to1", name: "file1.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
									},
									hasChild: false,
									eda:      []expectedEda{},
								},
								{
									name: "file2.go",
									ha: expectedHappa{
										absPath: abs + "/path1/to1/file2.go", relPath: "path1/to1/file2.go", dir: "path1/to1", name: "file2.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
									},
									hasChild: false,
									eda:      []expectedEda{},
								},
							},
						},
						{
							name: "to2",
							ha: expectedHappa{
								absPath: abs + "/path1/to2", relPath: "path1/to2", dir: "path1", name: "to2", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: true,
							eda: []expectedEda{
								{
									name: "sample.txt",
									ha: expectedHappa{
										absPath: abs + "/path1/to2/sample.txt", relPath: "path1/to2/sample.txt", dir: "path1/to2", name: "sample.txt", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
									},
									hasChild: false,
									eda:      []expectedEda{},
								},
							},
						},
					},
				},
				{
					name: "path2",
					ha: expectedHappa{
						absPath: abs + "/path2", relPath: "path2", dir: ".", name: "path2", isDir: true, isHiddenFile: false, isSymlink: false, realName: "",
					},
					hasChild: true,
					eda: []expectedEda{
						{
							name: "example1.go",
							ha: expectedHappa{
								absPath: abs + "/path2/example1.go", relPath: "path2/example1.go", dir: "path2", name: "example1.go", isDir: false, isHiddenFile: false, isSymlink: false, realName: "",
							},
							hasChild: false,
							eda:      []expectedEda{},
						},
					},
				},
				{
					name: "symlink.txt",
					ha: expectedHappa{
						absPath: abs + "/symlink.txt", relPath: "symlink.txt", dir: ".", name: "symlink.txt", isDir: false, isHiddenFile: false, isSymlink: true, realName: "example.txt",
					},
					hasChild: false,
					eda:      []expectedEda{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ki.Make(path, tt.args.option)
			if err != nil {
				t.Fatalf("Make() error = %v", err)
			}
			assertEda(t, ".", tt.want, got.Eda())
		})
	}
}

type expectedEda struct {
	name     string
	ha       expectedHappa
	hasChild bool
	eda      []expectedEda
}
type expectedHappa struct {
	absPath      string
	relPath      string
	dir          string
	name         string
	isDir        bool
	isHiddenFile bool
	isSymlink    bool
	realName     string
}

func testPath(t *testing.T) (string, string) {
	path := "../../testdata"
	abs, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	return path, abs
}

func assertEda(t *testing.T, name string, expected []expectedEda, actual []gooki.Eda) {
	if len(expected) != len(actual) {
		t.Errorf("%q gooki.Eda length want = %d, got = %d", name, len(expected), len(actual))
		return
	}

	for i := 0; i < len(expected); i++ {
		ee := expected[i]
		ae := actual[i]

		if got := ae.HasChild(); ee.hasChild != got {
			t.Errorf("[%s]gooki.Eda.HasChild() want = %v, got = %v", ee.name, ee.hasChild, got)
		}

		eh := ee.ha
		ah := ae.Happa()

		if got := ah.AbsPath(); eh.absPath != got {
			t.Errorf("[%s]gooki.Happa.AbsPath() want = %q, got = %q", ee.name, eh.absPath, got)
		}
		if got := ah.RelPath(); eh.relPath != got {
			t.Errorf("[%s]gooki.Happa.RelPath() want = %q, got = %q", ee.name, eh.relPath, got)
		}
		if got := ah.Dir(); eh.dir != got {
			t.Errorf("[%s]gooki.Happa.Dir() want = %q, got = %q", ee.name, eh.dir, got)
		}
		if got := ah.Name(); eh.name != got {
			t.Errorf("[%s]gooki.Happa.Name() want = %q, got = %q", ee.name, eh.name, got)
		}
		if got := ah.IsDir(); eh.isDir != got {
			t.Errorf("[%s]gooki.Happa.IsDir() want = %v, got = %v", ee.name, eh.isDir, got)
		}
		if got := ah.IsHiddenFile(); eh.isHiddenFile != got {
			t.Errorf("[%s]gooki.Happa.IsHiddenFile() want = %v, got = %v", ee.name, eh.isHiddenFile, got)
		}
		if got := ah.IsSymlink(); eh.isSymlink != got {
			t.Errorf("[%s]gooki.Happa.IsSymlink() want = %v, got = %v", ee.name, eh.isSymlink, got)
		}
		if got := ah.RealName(); eh.realName != got {
			t.Errorf("[%s]gooki.Happa.RealName() want = %q, got = %q", ee.name, eh.realName, got)
		}

		assertEda(t, ee.name, ee.eda, ae.Child())
	}
}
