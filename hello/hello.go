// package hello is a simple package to test calling of package created under the folder of main()
package hello

import "fmt"

//  func PrintHello() prints Hello that it's from pkg hello
func PrintHello() {
	fmt.Println("Hello from pkg hello")
}

//  func Hi() prints Hi that it's from pkg hello
func Hi() {
	fmt.Println("Hi from pkg hello")
	fmt.Println("Hello")
}
