package isPrime

// import (
// 	"fmt"
// 	"math"
// 	"sync"
// )

// func isPrime(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}
// 	sqrtNum := int(math.Sqrt(float64(n)))
// 	for i:=2; i<=sqrtNum; i++{
// 		if n%i==0{
// 			return false
// 		}
// 	}
// 	return true
// }

// func checkPrimes(start, end int, ch chan int, wg *sync.WaitGroup){
// 	defer wg.Done()
// 	for i:= start; i<=end;i++{
// 		if isPrime(i){
// 			ch <-i
// 		}
// 	}
// }

// func main(){
// 	const max = 100
// 	const workers = 5

// 	var wg sync.WaitGroup
// 	primes := make(chan int)

// 	chunk := max/ workers

// 	for i:= 0; i<workers;i++{
// 		start := i*chunk + 1
// 		end := (i+1) *chunk
// 		if i == workers -1{
// 			end =max
// 		}
// 		wg.Add(1)
// 		go checkPrimes(start, end, primes, &wg)
// 	}

// 	go func(){
// 		wg.Wait()
// 		close(primes)
// 	}()
// 	fmt.Println("Prime numbers:")
// 	for p := range primes {
// 		fmt.Print(p, " ")
// 	}
// 	fmt.Println("\nDone.")
// }

