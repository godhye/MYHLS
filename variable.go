// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//함수 밖에서는 var 반드시 붙이기 , 타입 생략 가능
var strglobalname = "standard";


func main() {
	
	const (
	//, 사용 안됨 , 값 명시하지 않으면 그전값 , iota 해당 변수의 idx 저장
	n1 = "hyejinseo";
	n2 
	n3 = 3;
	n4
	n5 = iota
	n6
	
	)
	fmt.Println("Hello, goorm!")
	
	var num1  = 3;
	var num2  = 0;
	
	num3 := num1 + num2;
	strname := "hyejinseo"
	stra := "a";
	
	var d = true;
	
	var str1 , str2 string = "hyejinse" , "seo";
	
	//const 타입 생략 가능
	const nconst  = 3;
	
	const strconstname string = "godhye";
	
	fmt.Println(num3);
	fmt.Println(strname);
	fmt.Println(stra);
	fmt.Println(d);
	fmt.Println(str1 , str2);

	fmt.Println(nconst , strconstname);
	fmt.Println(n1 ,n2 , n3 , n4 , n5 , n6);
	
}
