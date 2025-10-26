package main

import (
	"math/bits"
	"unsafe"
)

func ToLittleEndianGeneric[T uint16 | uint32 | uint64](number T) T {
	if number == 0 {
		return number
	}

	size := int(unsafe.Sizeof(number))
	result := make([]byte, size)
	u := unsafe.Pointer(&number)
	for i := range size {
		result[size-1-i] = *(*byte)(unsafe.Add(u, uintptr(i)))
	}

	return *(*T)(unsafe.Pointer(&result[0]))
}

func ToLittleEndian32(number uint32) uint32 {
	size := int(unsafe.Sizeof(number))
	res := make([]byte, size)

	u := unsafe.Pointer(&number)
	for i := range size {
		res[size-1-i] = *(*byte)(unsafe.Add(u, uintptr(i)))
	}

	return *(*uint32)(unsafe.Pointer(&res[0]))
}

func ToLittleEndianReverseGeneric[T uint16 | uint32 | uint64](number T) T {
	if number == 0 {
		return number
	}

	size := int(unsafe.Sizeof(number))
	var result T = 0
	var mask T = 0xff

	for i := range size {
		shiftRight := T(i * 8)
		byteValue := (number >> shiftRight) & mask

		shiftLeft := T((size - 1 - i) * 8)
		result |= byteValue << shiftLeft
	}

	return result
}

func ToLittleEndianReverse32(number uint32) uint32 {
	const size = 4
	var result uint32 = 0x00000000
	var mask uint32 = 0xff

	for i := range size {
		shiftRight := uint32(i * 8)
		byteValue := (number >> shiftRight) & mask

		shiftLeft := uint32((size - 1 - i) * 8)
		result |= byteValue << shiftLeft
	}

	return result
}

func ToLittleEndianLib32(number uint32) uint32 {
	return bits.ReverseBytes32(number)
}

func ToLittleEndianLib64(number uint64) uint64 {
	return bits.ReverseBytes64(number)
}
