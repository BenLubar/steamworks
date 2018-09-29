// +build linux darwin
// +build 386 amd64

package internal

/*
#include <stdint.h>

typedef struct uint64aligned { uint32_t value[2]; } uint64aligned;
typedef struct int64aligned { uint32_t value[2]; } int64aligned;
*/
import "C"

func (u64 C.uint64aligned) Get() uint64 {
	return uint64(u64.value[0]) | uint64(u64.value[1])<<32
}

func (u64 *C.uint64aligned) Set(value uint64) {
	u64.value[0] = C.uint32_t(value)
	u64.value[1] = C.uint32_t(value >> 32)
}

func (i64 C.int64aligned) Get() int64 {
	return int64(uint64(i64.value[0]) | uint64(i64.value[1])<<32)
}

func (i64 *C.int64aligned) Set(value int64) {
	i64.value[0] = C.uint32_t(value)
	i64.value[1] = C.uint32_t(uint64(value) >> 32)
}
