package Examples

/*
#include "../Stack/Stack.c"
*/
import "C"

// ExampleStack Demonstrates a stack
func ExampleStack() {
	// Create a stack
	stack := C.createStack()
	// Push some sample data
	C.push(stack, 1)
	C.push(stack, 2)
	C.push(stack, 3)
	// Pop the data
	C.pop(stack)
	println("Top of stack: ", stack.top.val)
}
