// +build windows
// +build 386 amd64

package internal

/*
#include <stdint.h>

typedef struct uint64aligned { uint64_t value[1]; } uint64aligned;
typedef struct int64aligned { int64_t value[1]; } int64aligned;
*/
import "C"

func (u64 C.uint64aligned) Get() uint64 {
	return uint64(u64.value[0])
}

func (u64 *C.uint64aligned) Set(value uint64) {
	u64.value[0] = C.uint64_t(value)
}

func (i64 C.int64aligned) Get() int64 {
	return int64(i64.value[0])
}

func (i64 *C.int64aligned) Set(value int64) {
	i64.value[0] = C.int64_t(value)
}
