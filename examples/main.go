package main

import "fmt"
import "github.com/seborama/govcr"

func runExample(name, cassetteName string, f func()) {
	fmt.Println("Running " + name + "...")
	fmt.Println("1st run =======================================================")
	govcr.DeleteCassette(cassetteName, "")
	f()
	fmt.Println("2nd run =======================================================")
	f()
	fmt.Println("Complete ======================================================\n\n")
}
func main() {
	runExample("Example1", example1CassetteName, Example1)
	runExample("Example2", example2CassetteName, Example2)
	runExample("Example3", example3CassetteName, Example3)
	runExample("Example4", example4CassetteName, Example4)
}
