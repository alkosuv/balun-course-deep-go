package main

import "testing"

var numberToTest32 uint32 = 0x01020304
var numberToTest64 uint64 = 0x0102030405060708

func BenchmarkToLittleEndianGeneric64(b *testing.B) {
	for b.Loop() {
		ToLittleEndianGeneric(numberToTest64)
	}
}
func BenchmarkToLittleEndianReverseGeneric64(b *testing.B) {
	for b.Loop() {
		ToLittleEndianReverseGeneric(numberToTest64)
	}
}

func BenchmarkToLittleEndianLib64(b *testing.B) {
	for b.Loop() {
		ToLittleEndianLib64(numberToTest64)
	}
}

func BenchmarkToLittleEndian32(b *testing.B) {
	for b.Loop() {
		ToLittleEndian32(numberToTest32)
	}
}

func BenchmarkToLittleEndianReverse32(b *testing.B) {
	for b.Loop() {
		ToLittleEndianReverse32(numberToTest32)
	}
}

func BenchmarkToLittleEndianLib32(b *testing.B) {
	for b.Loop() {
		ToLittleEndianLib32(numberToTest32)
	}
}
