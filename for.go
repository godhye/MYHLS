package main

import "fmt"

func main() {
	var i , j ,k int;
	
	//i,j는 두 개의 반복문에 쓰일 변수
	
	fmt.Scanln(&k);
	
	for i=1;i<=k;i++ {
		
		
		for j=1;j<=i;j++ {
		
			if (j == i){
			fmt.Printf("* ")
			continue
		}
			fmt.Printf("o ")
		}
		fmt.Printf("\n")
	}
	
	
}
