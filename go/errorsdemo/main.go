// Error handling: Go me exceptions nahi hote, error ek normal value hai.
package main

import (
	"errors"
	"fmt"
)

// sentinel error: isko errors.Is se compare karte hain
var ErrNotFound = errors.New("not found")

// custom error type: extra data rakhna ho to
type ValidationError struct{ Field string }

func (e *ValidationError) Error() string { return "invalid field: " + e.Field }

func find(id int) error {
	if id == 0 {
		return &ValidationError{Field: "id"}
	}
	if id > 10 {
		// %w se error wrap hota hai, upar wala errors.Is se original nikaal sakta hai
		return fmt.Errorf("find(%d): %w", id, ErrNotFound)
	}
	return nil
}

func main() {
	err := find(99)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("wrapped error mila:", err)
	}

	err = find(0)
	var ve *ValidationError
	if errors.As(err, &ve) { // As = type se match karke nikaalo
		fmt.Println("validation fail, field:", ve.Field)
	}

	// panic sirf bugs ke liye, normal errors ke liye nahi. recover se pakad sakte hain
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	var arr []int
	_ = arr[5] // index out of range panic
}
