package demo

import(
	"fmt"
	"reflect"
)
type cat string
type dog string 

type animal interface{
	say() string
}

func (c cat) say() string{
	return "meow"
}
func (d dog) say() string{
	return "bow bow"
}

func Interface1(){
	var c1,c2 animal
	c1=new(cat)
	c2=new(dog)
	fmt.Println(reflect.TypeOf(c1),c1.say())
	fmt.Println(c2.say())
}