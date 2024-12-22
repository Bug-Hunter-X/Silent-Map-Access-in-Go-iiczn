# Silent Map Access in Go

This example demonstrates a subtle issue in Go's map handling: accessing a non-existent key returns the zero value instead of causing a panic or error. This behavior can lead to unexpected results if not properly handled.

## The Bug

The `main` function attempts to access the value associated with the key "a" in the map `m`.  However, `m` is uninitialized, so there is no entry for this key.
Instead of panicking, Go returns 0 (the zero value for `int`), which might be mistaken for a valid value.

## The Solution

The solution involves checking if the key exists using `m["a"]`, which returns the value and the boolean of whether it's present or not. This provides a clean and efficient way to handle potential key absence.
