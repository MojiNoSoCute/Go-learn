package main

import (
	"fmt"
	"hello/cal"
)

//basic
//var (name) (type) = value .........   type = string , int, float, bool --- string must use ""
//(name) := value
//(name1), (name2) := value(n1),  value(n2)
//const (name) (type) = value    ...const = can't config

/*func main(){
	name := "Watcharakorn"
	age, High := 18, 177
	var W, H = 50, 177
	var pass bool = true
	const lastname string = "iamsamang"

	fmt.Println("Hello ",name)
	fmt.Println("age ", age)
	fmt.Println("pass =", pass)
	fmt.Println("lastname =", lastname)
	fmt.Println("hight = ", High)



}

// + - * /

func main(){
	n1, n2 := 10, 45
	var n3, n4 = 34, 23
	fmt.Println(n1*n2)
	fmt.Println(n3+n4)

}

// == != > < >= <=

func main(){
	n5, n6, n7,n8 := 35, 45, 56, 45
	fmt.Println(n5!=n6)
	fmt.Println(n7<n8)
	fmt.Println(n5==n6)
	fmt.Println(n7>n6)
	fmt.Println(n5<=n8)
    fmt.Println(n5>=n7)

}


//fmt.Scanf("(any%)", (name)) ... %s=string (Text), %d=int (integer), %f=float (floating point)

func main(){
	var name string
	var score int
	fmt.Print("enter name = ")
	fmt.Scanf("%s", &name)
	fmt.Print("enter ur score = ")
	fmt.Scanf("%d", &score)


	fmt.Println("hello ", name)
	fmt.Println("pass", score>=70)

}


//if_else -->  if___{   }, if___{  }else{  }

func main(){
  var score int
	fmt.Print("enter ur score = ")
	fmt.Scanf("%d", &score)
	fmt.Println("score =", score)

	if score >= 45 {
		fmt.Println("pass")
	}else{
		fmt.Println("not pass")
	}

	if score%2 ==0{
		fmt.Println("เลขคู่")
	}else{
		fmt.Println("เลขคี่")
	}


    var number int
	fmt.Print("enter number = ")
	fmt.Scanf("%d", &number)

	if number == 1 {
		fmt.Println("register")
	}else if number == 2 {
		fmt.Println("login")
	}else{
		fmt.Println("wrong")
	}



}


// switch_case ... look like if_else

func main(){
	var number int
	fmt.Print("enter number = ")
	fmt.Scanf("%d", &number)

	switch number{
	case 1:
		fmt.Println("register")
	case 2:
		fmt.Println("login")
	default:
		fmt.Println("wrong")
	}

}


//Array...เริ่มนับจำนวนที่เก็บตั่งแต่ง0...var (name) [จำนวนที่จะเก็บ] (type) =  [จำนวนที่จะเก็บ] (type){n1,n2,n3,...,nn}
//                               (name) :=  [จำนวนที่จะเก็บ] (type){n1,n2,...,nn}
func main(){
   Fname := [3]string{"Nai","New","Nite"}
   var number [2]int = [2]int{100, 200}
   numbers :=[...]int{123,2323,2323, 345345,345344,345345,45646,45646,64564} //no cap with[...]
   size := len(number)

   fmt.Println(Fname[0])
   fmt.Println(number[1])
   fmt.Println("size number ",size)
   fmt.Println(numbers)

}

//Slices look like Array ..... (name):=[]string {"som","chai"}
//index start at 0               can add value ..... (name)=append ((name),"gon")
func main(){
	numbers :=[]int{100,200}
	//numbers[1] = 300 //...change value
    numbers = append(numbers, 300) //add value
	numbers = append(numbers, 400)
	numbers = append(numbers, 500)
	fmt.Println(numbers)
	fmt.Println(numbers[1:])  //[1:]--> 1-last     [:1] = <1      [1:3]
	}



//Maps look like array..... var (name) map [key_type]value_type
//                           (name):=map [key_type]value_type

func main() {
	country:= map[string]string{"TH":"Thailand", "JP":"Japan"}
	country["EN"] = "English"

	value, checkkey := country["JP"]

	if checkkey{  //ถ้า check = true จะ print value
		fmt.Println(value)
	}else{
		fmt.Println("no value")
	}
}



//For_Loop ..... 	for count :=1;count<=10 ;count++ {______}           =1; <=1o = 1-10  , =10; >=1 =10-1
//                                      ++เพิ่มที่ละ1    --ลดทีละ1
// for ค่าเริ่มต้นของตัวแปร; เงื่อนไข; เปลี่ยนแปลงสค่าตัวแปร{คำสั่งเมื่อเมื่อเงื่อนไขเป็นจิง}
func main() {
	for count :=1; count<=3; count++ {
		fmt.Println(count)
	}
}



//break_continue -break = stop loop -continue = skip

func main() {
	for count :=1; count<=10; count++ {
		if count ==5 {
			continue
		}
		fmt.Println(count)
	}
}



//Range  ...better for
func main(){
	numbers := []int{10, 20, 30, 40, 50}
	for index := 0; index < len(numbers); index++{
        //fmt.Println(numbers[index])
		fmt.Println("index = ", index , "value = ", numbers[index])
	}

	for index, value := range numbers {
		fmt.Println("index = ", index , "value = ", value)
	}

	for _, value := range numbers {
		fmt.Println("value = ", value)
	}

    language := map[string]string{"TH":"Thailand", "JP":"Japan"}

	for key, value := range language {
		fmt.Println("key =", key , "value = ", value)
	}


}


//Function .... create func and use it in main

func main(){
	//showMessage("Nai")

	//total(50, 100)

	//delivery := delivery()
	//fmt.Println("delivery = ",delivery)

	mycart := totalcart(456,670)
	fmt.Println("totalcart =", mycart)
}
//แบบไม่รับและส่งข้อมูล
//func showMessage(name string) {
	//fmt.Println("Hello", name)
//}

//แบบรับข้อมูล
//func total(number1, number2 int){
     //fmt.Println("total =", number1+number2)
//}

//แบบส่งข้อมูล
//func delivery() int {
	//return 50
//}

//แบบรับและส่ง
func totalcart(number1, number2 int) int {
	total := number1 +number2
	return total

}



//return multiple value  ...... return multiple value in function

func main(){
	result, check := summation(5,2)
	fmt.Println("total =", result)
	fmt.Println(check)

}

func summation(number1, number2 int) (int, string) {
	total := number1 + number2
	status := ""
	if total%2 == 0 {
		status = "เลขคู่"
	}else{
		status = "เลขคี่"
	}
	return  total, status
}



//variadic function

func main(){
	result := summation(1,45,56,23,45)
	fmt.Println("total =", result)

}

func summation(number ... int) int {
	total := 0
	for _,value := range number {
		total += value                 //+=  ---> C=C+A
	}
	return  total
}



//Structure

type Product struct{
	name     string
	price    float32
	category string
	discount int
}

func main() {
	Product1 := Product{name:"ปากกา", price:50, category:"เครื่องเขียน", discount: 10}
	Product2 := Product{name:"Mouse", price:500, category:"IT", discount: 100}
	Product3 := Product{name:"Keyborad", price:550, category:"IT", discount: 120}
	Product1.price = 100
	fmt.Println(Product1, Product2, Product3)

}


//Package fix****************go mod init hello******************fix in terminal

func main() {
	fmt.Println(cal.Add(348,47))
	fmt.Println(cal.Add(5645,435))
}
*/