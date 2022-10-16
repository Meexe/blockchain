package tools

import "unsafe"

func Merge[T any](cs ...chan T) chan T {
	var res = make(chan T)

	for _, ch := range cs {
		go func(from, to chan T) {
			for v := range from {
				to <- v
			}
		}(ch, res)
	}
	return res
}

func Append[T any](slices ...[]T) []T {
	var res []T
	for _, slice := range slices {
		res = append(res, slice...)
	}
	return res
}

// IntToByteArray переводит int64 в массив байт
func IntToByteArray(num int64) []byte {
	size := int(unsafe.Sizeof(num))
	arr := make([]byte, size)
	for i := 0; i < size; i++ {
		byt := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&num)) + uintptr(i)))
		arr[i] = byt
	}
	return arr
}

// ByteArrayToInt переводит массив байт в int64
func ByteArrayToInt(arr []byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
