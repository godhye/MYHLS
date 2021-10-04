package main

import "fmt"
import "reflect"

func main() {
	  
	//map
	//key 값, 그 값에 해당하는 value값 저장하는 hashtable
	//var mapa map[string]string //nil map 생성

	var mapb = map[string]string{
		"apple":"red",
		"banana":"yellow",
	}
	
	fmt.Println(mapb)
	
	var mapc = make(map[string]string )
	
	mapc["hyejinse"] ="seo"
	mapc["hyejinseo"] ="se"
	
	fmt.Println(mapc)
	delete(mapc , "hyejinse")
	
	fmt.Println(mapc)
	
	val , exist := mapc["hyejinse"]
	fmt.Println(val , exist)
	
/*	
	- zero value는 명시적인 초기값을 할당하지 않고 변수를 만들었을 때 해당 변수가 갖게 되는 값이다.
	- nil은 포인터, 인터페이스, 맵, 슬라이스, 채널, 함수의 zero value이다. 	
	출처: https://2kindsofcs.tistory.com/3 [세상에는 두 종류의 cs가 있습니다.]

var a []int b := make([]int, 0) c := []int{} var d []int d = nil e := make([]int, 0, 0) f := make([]int, 0, 1) fmt.Println(a == nil, // true b == nil, // false c == nil, // false d == nil, // true e == nil, // false f == nil) // false

출처: https://2kindsofcs.tistory.com/3 [세상에는 두 종류의 cs가 있습니다.]

*/


	if(val == empty(val)){
		fmt.Println( reflect.TypeOf(val),"nil value")
	
		
	}
	val , exist = mapc["hyejinseo"]
	fmt.Println(val , exist)
	
	_, exist = mapc["hyejinseee"]
	fmt.Println(exist)
}
