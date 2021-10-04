// You can edit this code!
// Click here and start typing.
package main

import "fmt"


func main() {

	str1 := "hyejinseo"
	str2 := "2"
	
	num1 := 3 //0000 0011
	fmt.Println("name is " , str1 + str2)
	fmt.Println(num1 , str1);
	
	num1 &= 2 ; // 0000 0010 AND
	fmt.Printf("num1 &= 2 %08b \n " , num1)
	num1 |= 5    // 0000 0101 OR
	fmt.Printf("num1 |= 5 %08b \n " , num1)
	num1 ^= 1   // 0000 0001 XOR
	fmt.Printf("num1 ^= 1 %08b \n " , num1)
	num1 &^= 2 //AND OR
	fmt.Printf("num1 &^= 2 %08b \n " , num1)
	num1 <<= 2 //shift left
	fmt.Printf("num1 <<= 2 %08b \n " , num1)
	num1 >>= 2 //shift left
	fmt.Printf("num1 >>= 2 %08b \n " , num1)
	
	
	ch := make(chan int);
	
	go func(){
		ch <- 10
		
	}()
	
	result := <- ch
	fmt.Println(result)
	
}
