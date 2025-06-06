// Package declaration
package packs

// Import required packages for formatting and unsafe pointer operations
import (
	"fmt"
	"unsafe"
)

// Main function - entry point of the program
func DataTypes() {
	// Declare and initialize an integer variable
	x := 42

	var y rune = 100 // Declare another integer variable

	// Create a pointer to the integer x
	ptr := &x

	// Convert the pointer to uintptr using unsafe.Pointer
	// This allows direct pointer arithmetic (though generally unsafe)
	up := uintptr(unsafe.Pointer(ptr))

	// Print both the pointer and its uintptr representation
	fmt.Printf("Pointer: %p, uintptr: %x , Value : %v\n", ptr, up, x)

	fmt.Printf("Rune value: %v, Character: %c, Addr: %p\n", y, y, &y)
}
