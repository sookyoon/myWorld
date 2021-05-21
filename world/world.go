// package world is a simple package to test calling of package created under the folder of main()
package world

import "fmt"

//  func PrintWorld() prints from pkg world - the Earth
func PrintWorld() {
	fmt.Println("from pkg world - the Earth")
}

//  func FromUpperThomson() prints from pkg world - Upper Thomson
func FromUpperThomson() {
	fmt.Println("from pkg world - Upper Thomson")
}
