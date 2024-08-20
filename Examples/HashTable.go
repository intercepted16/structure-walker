package Examples

import "C"
import (
	"fmt"
	"unsafe"
)

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lmurmurhash -Wl,-rpath,/usr/local/lib
#include "../HashTable/HashTable.c"
#include <stdlib.h>
#include <string.h>
*/
import "C"

func ExampleHashTable() {
	// Create a new hash table
	ht := C.createHashTable()
	if ht == nil {
		fmt.Println("Failed to create hash table")
		return
	}
	defer C.freeHashTable(ht) // Ensure the hash table is freed when done

	// Convert Go strings to C strings and append to hash table
	key1 := C.CString("key1")
	defer C.free(unsafe.Pointer(key1)) // Free the C string when done
	value1 := C.CString("value1")
	defer C.free(unsafe.Pointer(value1)) // Free the C string when done
	C.appendToHashTable(ht, key1, unsafe.Pointer(value1), C.size_t(len("value1")+1))

	key2 := C.CString("key2")
	defer C.free(unsafe.Pointer(key2)) // Free the C string when done
	value2 := C.CString("value2")
	defer C.free(unsafe.Pointer(value2)) // Free the C string when done
	C.appendToHashTable(ht, key2, unsafe.Pointer(value2), C.size_t(len("value2")+1))

	// Retrieve a value
	value := C.getFromHashTable(ht, key1)
	if value == nil {
		fmt.Println("Value not found for key1")
	} else {
		fmt.Println("Value for key1:", C.GoString((*C.char)(value)))
	}
}
