// package main

// import "fmt"

// func main(){
// 	var employees map[string]string
// 	fmt.Printf("%#v\n",employees)

// 	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

// 	people := map[string]float64{}

// 	people["John"]= 21.4
// 	people["Marry"] = 10

// 	fmt.Println(people)

// 	map1 := make(map[int]int)
// 	map1[4] = 8

// 	balances := map[string]float64{
// 		"USD" : 323.11,
// 		"EUR" : 432.11,
// 	}
// 	fmt.Println(balances)

// 	m := map[int]int{1:10, 2:20,3:30}
// 	_ = m

// 	balances["USD"] = 500.2
// 	balances["NPR"] = 120.1
// 	fmt.Println(balances)

// }

package main

import "fmt"

func main(){
	a := map[string]string{"A": "X"}
	b := map[string]string{"B":"Y"}

	s1 := fmt.Sprintf("%s\n",a)
	s2 := fmt.Sprintf("%s\n",b)

	fmt.Println(s1,s2)

	if s1 == s2{
		fmt.Println("Maps are equal")
	}else{
		fmt.Println("Maps are not equal")
	}
}