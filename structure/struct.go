package main

import "fmt"

func main1() {
	title1, author1, year1 := "The divine comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	type book struct{
		title string
		author string 
		year int
	}

	myBook := book{
		title:  "The divine comedy",
		author: "Dante Alighieri",
		year:   1320,
	}

	fmt.Println(myBook)
}
