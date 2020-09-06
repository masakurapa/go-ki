package ki_test

import (
	"bytes"
	"testing"

	"github.com/masakurapa/gooki/internal/ki"
	"github.com/masakurapa/gooki/pkg/gooki"
)

func BenchmarkMake(b *testing.B) {
	opt := gooki.DefaultOption()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ki.Make("../../testdata", opt)
		if err != nil {
			b.Fatalf("Make() error = %v", err)
		}
	}
}

func BenchmarkKi_Write(b *testing.B) {
	got, err := ki.Make("../../testdata", gooki.DefaultOption())
	if err != nil {
		b.Fatalf("Make() error = %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		if err = got.Write(buf); err != nil {
			b.Fatalf("Write() error = %v", err)
		}
	}
}
