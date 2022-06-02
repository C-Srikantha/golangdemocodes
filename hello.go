package main

import (
	"fmt"
	//"reflect"
	//"strconv"
	"errors"
	"golang/demo"
	"time"
)

func operation(a, b int) (int, int, int, int, error) {
	sum := a + b
	sub := a - b
	mul := a * b
	if b == 0 {
		err := errors.New("var b cant be zero")
		return sum, sub, mul, 0, err
	} else {
		div := a / b
		return sum, sub, mul, div, nil
	}
}

func main() {
	/*var(
	        a=0
	       b=0
	    //c=3.0
	       // strName = "100"
	    )
	 //fmt.Println("Hello, World!")
	 //fmt.Println(a,b)
	 add,sub,mul,div,err:=operation(a,b)//calling func operation
	      if(err!=nil){
	      fmt.Println(err)
	      }else{
	         fmt.Println(add,sub,mul,div)
	      }
	/*typecon,err:=strconv.Atoi(strName)//Atoi func returns two value 1.result of conversion 2.error

	fmt.Println(typecon,err,reflect.TypeOf(typecon))//returns type of variable,typeconversion string to int
	repeat(add,sub,mul,div)//example for if statements*/
	/*demo.Array()//example for array
	  demo.Student()//
	  workday()//example for switch
	  demo.Hashmap()
	  mt.Println("Area of rectangle r1 is",demo.Area(10.0,2.0))
	  demo.Example()*/
	//demo.Interface1()
	//demo.Library()
	//demo.Goroutine()
	demo.Channel2()
}
func repeat(add, sub, mul, div int) {
	if add < mul {

		fmt.Println("Mul value is greater than addition")
	} else if add == mul {
		fmt.Println("Sub value is greater than div")
	} else {
		fmt.Println(add, sub, mul, div)
	}

}

//switch statements
func workday() {
	time := time.Now()
	switch time.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 11:
		fmt.Println("Today is 11th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Today its 31 .Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

func goodmoring() {
	fmt.Println("good morning")
}
