package main

import (
	
	"fmt"
	"time"
)

type item struct{
	name string
	price int
	count int
}

type buyer struct{
	point int
	wish map[string]int 
}

func newBuyer() buyer{
	
	d := buyer{}
	d.point = 1000000
	d.wish = map[string]int{}
	
	return d
}
 

type delivery struct{
	status string
	onedeli map[string]int //배송되는 물품과 수량
	
}

func newDelivery() delivery{
	d := delivery{}
	d.onedeli = map[string]int{}
	return d
}

func buying(items []item , buyer1 *buyer , choice int , numbuy *int , ch chan bool , temp map[string]int){
	//복구용 
	defer func(){
		r := recover()
		if r != nil{
			fmt.Println(r)
			main()
		}
	}()
	
	 
	
	if(*numbuy >=5){
		fmt.Println("배송 한도를 초과햇습니다.")
		return
	}
	
	
	inputcount := -1
	inputmenu := -1
	fmt.Printf("수량을 입력하세요 :")
	fmt.Scanln(&inputcount)
	
	//수량이 포인트내로 결제 되는지 확인 , 주문가능한 수량인지
	if inputcount <=0 || inputcount > items[choice].count || buyer1.point < inputcount*items[choice].price{
		
		panic("수량을 잘못 입력하여 주문이 불가능합니다.")
	} 
	fmt.Println("구매 : 1 , 장바구니 2:")
	fmt.Scanln(&inputmenu)
	
	//구매 로직
	if inputmenu == 1{
		//포인트 차감
		buyer1.point -= items[choice].price * inputcount
		//수량 차감
		items[choice].count -= inputcount
	
		temp[items[choice].name] = inputcount
		*numbuy++
		ch <- true
		
	}else if inputmenu == 2{
	//장바구니 로직
		
		checklist := false
		
		//장바구니에 아이템 존재하는지 확인
		for itms := range buyer1.wish {
			if itms == items[choice].name{
			checklist = true		
			}
		}
		
		if checklist == true{
			if buyer1.wish[ items[choice].name] + inputcount <items[choice].count{
			buyer1.wish[ items[choice].name] += inputcount
			}else{
				fmt.Println("주문가능 수량을 초과햇습니다.")	
				return 
			}
		}else {
			buyer1.wish[ items[choice].name] = inputcount
		}
		fmt.Println("상품이 장바구니에 추가되었습니다.")
		
	}
}
func remainitem(items []item){
	
 
	for i:=0 ; i < len(items); i++{
		fmt.Println(i+1, ".",items[i].name , " : ",items[i].count ,"개", items[i].price)
		fmt.Println()
	}

}

func remainpoint(point int){
	
	fmt.Println("잔여 포인트 ",  point, "입니다.")
	
}
func checkshipping(deli []delivery){
	
		fmt.Println("배송목록 입니다.")
	for i :=0 ; i < len(deli) ; i++{
		
		fmt.Println(i+1 , " : ",deli[i].status)
	}
	
	
}
func checkwish(wish map[string]int){
	
	index := 1
	if(len(wish) == 0){
		fmt.Println("장바구니가 비어있습니다.")
	}else{
		
	for itm , cnt := range wish{
		fmt.Println(index ,itm, cnt)
		index++
	}
	}
}

func deliveryStatus(deli []delivery , ch chan bool , index int , numby *int , temp *map[string]int){
	
	for{
		if <-ch {
			
			for name , cnt :=range *temp{
				deli[index].onedeli[name] = cnt //임시저장한 데이터 배송 상품 데이터 저장 
				
			}
			
			//주문접수
			deli[index].status = "start"
			time.Sleep(time.Second * 10)
			
			//배송중
			deli[index].status = "shipping"
			time.Sleep(time.Second * 30)
			
			//배송완료
			deli[index].status = "end"
			time.Sleep(time.Second * 10)
			
			deli[index].status = ""
		
			// 배송완료
			*numby--
			//배송목록 초기화 
			*temp = map[string]int{}
		}
		
		
		
	}
	
}

func main() {
	//최대 주문 슈량
	numbuy := 0

	// 배송 객체 생성
	deliverylist := make([]delivery , 5)
	for i:=0 ; i<5 ; i++{
		deliverylist[i] = newDelivery()
		
	}
	
	//배송위한 채널 생성
	deliverystart := make(chan  bool)
	temp := map[string]int{}
	
	//items := [5]item{} // 아이템 객체 5개 item 배열 생성
	items := make([]item ,5) // 아이템 객체 5개 item 배열 생성
	
	for i:=0 ; i < 5 ; i++{
		
		time.Sleep(time.Millisecond)
		
		go deliveryStatus(deliverylist ,deliverystart , i ,  &numbuy , &temp)
		
	}
	buyer1 := newBuyer() // 구매자 객체
	
	items[0] = item{"텀블러" , 10000 , 30}
	items[1] = item{"롱패팅" , 500000 , 20}
	items[2] = item{"백팩" , 400000 , 20}
	items[3] = item{"운동화" , 150000 , 50}
	items[4] = item{"뺴뺴로" , 1200 , 500}

	//프로그램 메뉴	
	for{
		
		fmt.Printf(" 1.구매 \n 2.잔여 수량 확인 \n 3.잔여 마일리지 \n 4.배송 상태 확인 \n 5.장바구니 확인 \n 6.프로그램 종료 \n 7.실행할 메뉴 입력 :")
		var input int 
		fmt.Scan(&input)
		
		switch input {
			
			case 1 :
					remainitem(items) //물품 출력 
					choice := -1
					fmt.Printf("구매할 아이템 선택 :  ")
					fmt.Scanln(&choice)
					choice += -1
					if choice >0 || choice <6{
						buying(items , &buyer1 , choice,&numbuy , deliverystart , temp)	
					}else{
					fmt.Println("잘못입력햇습니다. ")
					}
			
			case 2 : remainitem(items)
			case 3 : remainpoint(buyer1.point)
			case 4 : checkshipping(deliverylist)
			case 5 : checkwish(buyer1.wish)
			
			
			fmt.Printf("1. 장바구니 상품 주문 2. 장바구니 초기화 3.메뉴로 돌아가기  : ")
			bucketinput := 0
			fmt.Scanln(&bucketinput)
			if bucketinput == 1{
				//장바구니내 물품 구매
				
				//현재 포인트내 구매 가능한지
				if !requirePoint(items , buyer1){
					fmt.Println("포인트가 부족해 주문이 불가능합니다.")
				}
				//현재 수량이 주문가능한지 
				if !checkwishcount(items , buyer1){
					fmt.Println("수량이 부족해 주문이 불가능합니다..")
				}
				bucketbuy(items , &buyer1, &numbuy, deliverystart , temp)
				
			}else if bucketinput == 2{
				//장바구니 초기화
				
			}else{
				//메뉴로 돌아가기
				break
			}
			
			case 6 : return 
			default : continue
			
		}
		
		
		
	}
	}

func bucketbuy(itms []item , bys *buyer , numbuy *int, ch chan bool, temp map[string]int){

 
	if(*numbuy >=5 ){
		fmt.Println("배송 한도를 초과 햇습니다.")
		return
	}
	if len(bys.wish) == 0{
		return
	}
	
	for itm, cnt := range bys.wish{
		for i:=0 ; i < len(itms) ; i++{
			if itm == itms[i].name{
				itms[i].count -= cnt
				bys.point -= itms[i].price * cnt
				temp[itms[i].name] = cnt			
			}
		}
		
	}
	
	
	//초기화
	bys.wish = map[string]int{}
	*numbuy++
	ch <- true
	fmt.Println("주문이 접수 되었습니다.")
}

func requirePoint(itms []item , bys buyer) (bool){
	
	totalprice := 0
	for itm , cnt := range bys.wish{
		for i:=0; i<len(itms); i++{
			if(itms[i].name == itm){
				totalprice += itms[i].price* cnt
			}
		}
	}
	
	fmt.Println("필요 포인트 : " ,totalprice)
	fmt.Println("잔여 포인트 : ", bys.point)
	
	
	if (totalprice > bys.point){
		return false	
	}else{
		return true
	}
	
}

func checkwishcount(itms []item , bys buyer)bool{
	
	for itm , cnt := range bys.wish{
		for i:=0 ; i<len(itms); i++{
			if(itm == itms[i].name){
				if(cnt > itms[i].count){
				return false	
				}
				
			}
		}
	}
	return true
}
