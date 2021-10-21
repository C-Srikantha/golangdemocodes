package demo

import (
	"fmt"
	"reflect"
    "sort"
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
//map example
func Hashmap() {
	//var num = map[string]int{"Mahesh": 10, "Ganesh": 20}
    var countryname map[string]string 
    countryname=make(map[string]string)
    countryname["India"]="New Delhi"
    countryname["Japan"]="Tokyo"
    countryname["England"]="London"
    fmt.Println(len(countryname))
    keys := make([]string, 0, len(countryname))//slice
 
	for k := range countryname {
		keys = append(keys,k)//we append key to the slice

	}
    sort.Strings(keys)
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

    first := map[string]int{"a": 1, "b": 2, "c": 3,"d":5}
	second := map[string]int{"z": 1, "f": 5, "e": 3, "d": 6}
 
	for k, v := range second {
		first[k] = v//merge map var  second to first
    }
    fmt.Println(first)
}
