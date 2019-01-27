package main

import "fmt"

func hello() {
	// only this file can access the imported fmt package when others also do so, the can also access their own 'fmt' "name"
	bye()
	fmt.Println("Alexandru")
}
