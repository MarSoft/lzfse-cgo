// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Sun, 06 Oct 2019 15:16:18 EDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package lzfse

/*
#cgo LDFLAGS: -L/usr/local/Cellar/lzfse/1.0/lib -L/usr/local/lib -llzfse
#include "lzfse.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import "unsafe"

// EncodeScratchSize function as declared in go-lzfse/lzfse.h:56
func EncodeScratchSize() uint {
	__ret := C.lzfse_encode_scratch_size()
	__v := (uint)(__ret)
	return __v
}

// EncodeBuffer function as declared in go-lzfse/lzfse.h:87
func EncodeBuffer(dstBuffer []byte, dstSize uint, srcBuffer string, srcSize uint, scratchBuffer unsafe.Pointer) uint {
	cdstBuffer, _ := (*C.uint8_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&dstBuffer)).Data)), cgoAllocsUnknown
	cdstSize, _ := (C.size_t)(dstSize), cgoAllocsUnknown
	csrcBuffer, _ := unpackPUint8String(srcBuffer)
	csrcSize, _ := (C.size_t)(srcSize), cgoAllocsUnknown
	cscratchBuffer, _ := scratchBuffer, cgoAllocsUnknown
	__ret := C.lzfse_encode_buffer(cdstBuffer, cdstSize, csrcBuffer, csrcSize, cscratchBuffer)
	__v := (uint)(__ret)
	return __v
}

// DecodeScratchSize function as declared in go-lzfse/lzfse.h:94
func DecodeScratchSize() uint {
	__ret := C.lzfse_decode_scratch_size()
	__v := (uint)(__ret)
	return __v
}

// DecodeBuffer function as declared in go-lzfse/lzfse.h:126
func DecodeBuffer(srcBuffer []byte) []byte {
	csrcBuffer, _ := unpackPUint8String(string(srcBuffer))
	csrcSize, _ := (C.size_t)(len(srcBuffer)), cgoAllocsUnknown

	dstBuffer := make([]byte, 4*len(srcBuffer))
	cdstBuffer, _ := (*C.uint8_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&dstBuffer)).Data)), cgoAllocsUnknown
	cdstSize, _ := (C.size_t)(4*len(srcBuffer)), cgoAllocsUnknown

	scratch := make([]byte, DecodeScratchSize())
	cscratchBuffer, _ := unsafe.Pointer(&scratch[0]), cgoAllocsUnknown

	__ret := C.lzfse_decode_buffer(cdstBuffer, cdstSize, csrcBuffer, csrcSize, cscratchBuffer)
	__v := (uint)(__ret)
	return dstBuffer[:__v]
}
