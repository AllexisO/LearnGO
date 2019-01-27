// Need to import fmt package three times with different names
// Need to print a few messages using those imports
// Ex. hello -> hey -> hi
package main

import "fmt"
import h "fmt"
import he "fmt"

func main() {
	fmt.Println("Hello")
	h.Println("Hey")
	he.Println("Hi")
}
