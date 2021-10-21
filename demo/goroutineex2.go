package demo

import(
	"fmt"
	"time"
)
func number(){
	for i:=0;i<5;i++{
		time.Sleep(1*time.Millisecond)
		fmt.Println(i)
	}
}
