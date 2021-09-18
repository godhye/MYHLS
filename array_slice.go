package main

import "fmt"

func main() {
	
	//배열
	var array1 [5]int ; //길이가 5인 int형 배열 선언
	fmt.Println(array1)
	
	array1 = [5]int{1,1,1,1,1} //배열 초기화
	fmt.Println(array1 , array1[1])
	
	array2 := [4]int{1,2,3,4} 
	array2[3] = 10;
	
	fmt.Println(array2 , len(array2))
	
	
	var multiarray [2][3]int; //2차원 배열 세로 x 가로
	
	multiarray = [2][3]int{
		{1,2,3},
		{2,3,4},
	}
	
	fmt.Println(multiarray)
	
	multiarray[0][0] =33;
	fmt.Println(multiarray)
	
	//슬라이스
	
	var slice []int = []int{1,2,3,4}
	fmt.Println(slice)
	slice[0] =100
	fmt.Println(slice)
	
	l := slice[0:1] //슬라이스 복제 시작인덱스 ~ 복사원하는갯수
	fmt.Println(l)
	
	var nilslice []string =[]string{"hyejin" ,"seo"} //nil형의 슬라이스생성
	
	if nilslice == nil{
		fmt.Println("이 슬라이스는 비어있습니다.")
	} else{
		fmt.Println("용량이" , cap(nilslice), "길이가" , len(nilslice), "입니다.")
		
	}
	
	//make 함수를 통한 컬렉션 선언
	//make( 슬라이스타입 , 슬라이스길이 , 슬라이스용량(생략시 길이와 같음))
	
	s := make([]int , 0 , 3)
	
	for i:=1; i<=10;i++ {
		s = append(s , i)
		fmt.Println(len(s), cap(s))
		//cap 증가 개념 , cap x2 더 들어갈 요소 있으면 +2 반복
	}
	fmt.Println(s)
		
	//슬라이스 추가 , 병합 , 복사
	//슬라이스에 공간있으면 추가 ,용량 초과 경우 설정한 용량만큼 새로운 배열 생성하고 기존 값 새 배열에 복제 후 다시 슬라이스를 할당
	//sllice A... 슬라이스 요소 집합으로 치환 ex {2,3,4}
	
	sliceA := []int{1,2,3}
	sliceB := []int{3,3,3}
	
	sliceC := make([]int , len(sliceB), cap(sliceB)*2)
	
	//슬라이스 자르기 slice[시작인덱스 , 마지막인덱스+1]
	//copy랑은 다르게 주소를 복사해오는거라 원본 값 바뀌면 같이 변경됨 
	sliceD := sliceB[0 :3]
	fmt.Println("SLICE D ,",sliceD)
	
	fmt.Println(sliceA)
	sliceA = append(sliceA , sliceB...)
	fmt.Println(sliceA , cap(sliceA))
	
	fmt.Println(sliceB,len(sliceB) , cap(sliceB))
	
	//요소는 복사하지만 속성은 변경되지않는다?
	copy(sliceC , sliceB);
	fmt.Println(sliceC, len(sliceC) , cap(sliceC))
	
	sliceB[0] =1000;
	fmt.Println(sliceC, len(sliceC) , cap(sliceC))
	
	fmt.Println("SLICE D ,",sliceD)
		
}
