package test03

import "testing"

type A struct {
	Name string
}

func Normal(v A) string {
	return v.Name
}

func NormalRef(v *A) string {
	return v.Name
}

func Benchmark_Normal(b *testing.B) {
	var a = A{"123"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Normal(a)
	}
}

func Benchmark_NormalRef(b *testing.B) {
	var a = &A{"123"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NormalRef(a)
	}
}

func Interface(v interface{}) string {
	return v.(A).Name
}

func Benchmark_Interface(b *testing.B) {
	var a = A{"123"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Interface(a)
	}
}

func InterfaceRef(v interface{}) string {
	return v.(*A).Name
}

func Benchmark_InterfaceRef(b *testing.B) {
	var a = &A{"123"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InterfaceRef(a)
	}
}
