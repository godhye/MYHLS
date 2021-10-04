package main

import "fmt"

type student struct {
	name string
	sex	string
	data map[string]int //과목명 - 점수
}

func newStucet() student{
	
	s := student{}
	s.name = ""
	s.sex = "F"
	s.data = map[string]int{}
	return s
}

func main() {
	var nstudent , ncourse int
	var strname ,strsex , strcourse string
	
	fmt.Scanln(&nstudent , &ncourse)
	//fmt.Println(nstudent , &ncourse)
	students := []student{}
	
	for i:=0 ;i <nstudent; i++ {	
		fmt.Scanln(&strname , &strsex)
		
		s := newStucet()
		s.name = strname
		s.sex = strsex
		
		for j:=0 ; j<ncourse; j++{
			fmt.Scanln(&strcourse ,&ncourse )
			s.data[strcourse] = ncourse
		}
		
		fmt.Println(s)
		students = append(students ,s )
	}
	
	for i:=0 ; i<nstudent ; i++ {
		fmt.Println("----------")
		//fmt.Println(students[i].name , students[i].sex)
		
		for index, val := range students[i].data{
			fmt.Println(index, val)
		}
		
	}
	fmt.Println("----------")
}
