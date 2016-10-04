package main 

import "fmt"

const(
	// create a huge number by shifting a 1 bit left 100places
	// In other words, the bynary number that is 1 followeed by 100 zeroes
	Big =  1 << 100
	// shift 	it right again 99 palces, so we end up with 1 << 1, or 2
	Small = Big >> 99
)

func neddInt(x int)int {
	return x * 10 + 1
}

func needFload( x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(neddInt(Small))
	fmt.Println(needFload(Big))
	fmt.Println(needFload(Small))
}