package demo

import(
	"fmt"

)
func Channel(){
message:=make(chan string)//Unbuffered channel
go func(){
	message<-"Srikantha"
}()
msg:=<-message
fmt.Println(msg)
}
