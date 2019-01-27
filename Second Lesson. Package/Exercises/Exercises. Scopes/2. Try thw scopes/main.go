// Need to create two files - main.go and printer.go
// In printer.go need to create a func named hello, inside the hello func, print your name.
// In main.go need to create the usual func main, call your func just by using its name 'hello' and inside the bye func, print 'bye bye';
// In printer.go need to call the bye func from inside the hello func.

package main

import "fmt"

func main() {
	// as you can see, I don't need to import a package and I can call 'hello func here.
	// this is because package-scoped names are shared in the same package
	hello()
	//but here, I can't access the fmt package without importing it.
	// this is because it's in the printer.go's file scope. It imports it.
	// this main func can also call bye func here.
	// bye()
}

// printer.go can call this function
// this because bye func is in the package-scope of the main package now.
// main func can also call this.
func bye() {
	fmt.Println("Bye Bye")
}
