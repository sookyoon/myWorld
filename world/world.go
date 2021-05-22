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

//  func FromBishanToaPayoh() prints from pkg world - FromBishanToaPayoh
//  added by contributor
func FromBishanToaPayoh() {
	fmt.Println("from pkg world - BishanToaPayoh")
}

//  func FromBishanToaPayoh1() prints from pkg world - FromBishanToaPayoh
//  cloned  by contributor in visual studio code
func FromBishanToaPayoh1() {
	fmt.Println("from pkg world - BishanToaPayoh - from vs code")
}
