package xgb

/*
help.go is meant to contain a rough hodge podge of functions that are mainly
used in the auto generated code. Indeed, several functions here are simple
wrappers so that the sub-packages don't need to be smart about which stdlib
packages to import.

Also, the 'Get..' and 'Put..' functions are used through the core xgb package
too. (xgbutil uses them too.)
*/

import (
	"encoding/binary"
	"fmt"
	"math/bits"
	"strings"
)

// StringsJoin is an alias to strings.Join. It allows us to avoid having to
// import 'strings' in each of the generated Go files.
func StringsJoin(ss []string, sep string) string {
	return strings.Join(ss, sep)
}

// Sprintf is so we don't need to import 'fmt' in the generated Go files.
func Sprintf(format string, v ...interface{}) string {
	return fmt.Sprintf(format, v...)
}

// Errorf is just a wrapper for fmt.Errorf. Exists for the same reason
// that 'stringsJoin' and 'sprintf' exists.
func Errorf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}

// Pad a length to align on 4 bytes.
func Pad(n int) int {
	return (n + 3) & ^3
}

// PopCount counts the number of bits set in a value list mask.
func PopCount(mask0 int) int {
	return bits.OnesCount(uint(mask0))
}

var byteOrder = binary.LittleEndian

// Put16 takes a 16 bit integer and copies it into a byte slice.
func Put16(buf []byte, v uint16) {
	byteOrder.PutUint16(buf, v)
}

// Put32 takes a 32 bit integer and copies it into a byte slice.
func Put32(buf []byte, v uint32) {
	byteOrder.PutUint32(buf, v)
}

// Put64 takes a 64 bit integer and copies it into a byte slice.
func Put64(buf []byte, v uint64) {
	byteOrder.PutUint64(buf, v)
}

// Get16 constructs a 16 bit integer from the beginning of a byte slice.
func Get16(buf []byte) uint16 {
	return byteOrder.Uint16(buf)
}

// Get32 constructs a 32 bit integer from the beginning of a byte slice.
func Get32(buf []byte) uint32 {
	return byteOrder.Uint32(buf)
}

// Get64 constructs a 64 bit integer from the beginning of a byte slice.
func Get64(buf []byte) uint64 {
	return byteOrder.Uint64(buf)
}
