package demo

import (
	"fmt"
	"reflect"
)
func Main(){
fmt.Println("hello")
}
func Array(){
var a[5] int
a[2]=2
var name[5] string= [5]string{"Srikantha", "Shreesha", "Adarsh", "Milyn", "Leander"}
num :=[...]int{1,2:3}//Initializing an Array with ellipses...
fmt.Println(reflect.TypeOf(a),a)
fmt.Println(len(num),"Elements:",num)
fmt.Print(name[:],"\n",name[0:4],"\n",name[len(name)-2:len(name)-1],"\n")


}
func Student(){
	var action int
    fmt.Println("Enter 1 for Student and 2 for Professional")
    fmt.Scanln(&action)
    //  Use of Switch Case   
    switch action {
        case 1:
            fmt.Printf("I am a  Student")
        case 2:
            fmt.Printf("I am a  Professional")
		default:
			panic(fmt.Sprintf("I am a  %d",action))
    }	
}
func Hashmap(){
	var num =map[string]int{"Mahesh:":10,"Ganesh":20}
	fmt.Println(num)
}