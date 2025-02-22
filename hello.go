package main

import(
	"fmt"
	"errors"
)

func divAndRemainder(numerator, denominator int)(quotient int, remainder int, err error){
	if denominator == 0{
		err = errors.New("Can't divide by 0")
		return
	}
	quotient = numerator / denominator
	remainder= numerator % denominator
	return
}

func main(){
	var numerator, denominator int

	fmt.Print("Enter Numerator:")
	fmt.Scan(&numerator)

	fmt.Print("Enter Denominator:")
	fmt.Scan(&denominator)

	quotient, remainder, err := divAndRemainder(numerator, denominator)
	if err != nil{
		fmt.Println("Error", err)
	}else{
		fmt.Printf("Quotient: %d, Remainder: %d\n", quotient, remainder)
	}


}