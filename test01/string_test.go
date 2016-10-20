package test01

import (
	"fmt"
	"testing"
)

func BenchmarkSimpleStringAppend(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s := "abc"
			s = s + "a" + s + "b" + s + "c" + s
		}
	})
}

func BenchmarkSimpleStringFormat(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			s := "abc"
			s = fmt.Sprintf("%sa%sb%sc%s", s, s, s, s)
		}
	})
}
