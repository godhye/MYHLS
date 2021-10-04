package main

import "fmt"

func main() {
	
	var num1 , num2 ,num3 int;
	
	fmt.Scanln(&num1 , &num2 , &num3);
	
	result := num1 * num2 +num3;
	
	fmt.Printf("%d x %d + %d = %d" , num1 , num2 , num3, result)
}
