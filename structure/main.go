package main

import (
	"io"
	"os"
)

// func counter() func() int{
// 	count := 0 // variable captured by closure
// 	return func() int{
// 		count ++ // modifies captured closure
// 		return count
// 	}
// }

func main(){
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
	myString = "Please give me one argument!"
	} else {
	myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}