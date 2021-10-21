package demo

import(
	"fmt"
	"time"
)
func label(word string){
	for i:=0;i<5;i++{
		time.Sleep(10*time.Millisecond)
		fmt.Println(word)
	}
}

func Goroutine(){
	go label("hello")
	go label("hi")
 	label("bye")

}
