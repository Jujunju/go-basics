package main

var myArray [4]string
var mySlices []string = []string{"jujun", "ina", "yono"}

func main() {

	mySlices = append(mySlices, "hi")
	myArray[0] = "ibra"
	myArray[1] = "ibra"
	myArray[2] = "ibra"
	myArray[3] = "ibra"

}