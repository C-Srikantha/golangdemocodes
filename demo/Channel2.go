package demo

import "fmt"

func myfun(mychnl chan string) {
  
    for i:=1;i<=4;i++{
        mychnl <- "hello"
    }
    close(mychnl)
}
  
func Channel2() {
  

    c := make(chan string,5)
    go myfun(c)
  
    for {
        res, ok := <-c
        if ok == false {
            fmt.Println("Channel Close ", ok)
            break
        }
        fmt.Println("Channel Open ", res, ok)
		fmt.Println(len(c))
    }
	fmt.Println(len(c),cap(c))
}