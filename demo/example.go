package demo

import (
	"fmt"
	"reflect"
    //"sort"
)
//array example
func Array() {
	var a [5]int
	a[2] = 2
	var name [5]string = [5]string{"Srikantha", "Shreesha", "Adarsh", "Milyn", "Leander"}
	num := [...]int{1, 2: 3} //Initializing an Array with ellipses...
	fmt.Println(reflect.TypeOf(a), a)
	fmt.Println(len(num), "Elements:", num)
	fmt.Print(name[:], "\n", name[0:4], "\n", name[len(name)-2:len(name)-1], "\n")
    for _,i:=range name{
        fmt.Println(i)
    }

}
func Student() {
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
		panic(fmt.Sprintf("I am a  %d", action))
	}
}
//map example
func Hashmap() {
	//var num = map[string]int{"Mahesh": 10, "Ganesh": 20}
    var countryname map[string]string 
    countryname=make(map[string]string)
    countryname["India"]="New Delhi"
    countryname["Japan"]="Tokyo"
    countryname["England"]="London"
    keys := make([]string, 0, len(countryname))//slice
 
	for k := range countryname {
		keys = append(keys,k)//we append key to the slice

	}
   // sort.Strings(keys)
    for _,k:=range keys{
        fmt.Println(k)
    }
    defer fmt.Println(reflect.ValueOf(countryname).Kind())//returs type of variable countryname
    fmt.Println("Before Deleting")
    fmt.Println("Size of Countryname variable:",len(countryname))
	for name := range countryname {
		fmt.Println("Capital of",name,"is", countryname[name])
	}
    delete(countryname,"Japan")//deletes particular elements
    //countryname=make(map[string]string)//deleting all elements from map variable
    fmt.Println("After Deleting")
    fmt.Println("Size of Countryname variable:",len(countryname))
    for name := range countryname {
		fmt.Println("Capital of",name,"is", countryname[name])
	}
}
