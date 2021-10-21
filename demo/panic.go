package demo

import(
	"fmt"
)
func handlepanic(){
	a:=recover()
	if a!=nil{
		fmt.Println("Recovered")
	}
}
func book(bookid *int,bookname *string){
	defer handlepanic()
	if bookid==nil {
		panic("Error:bookid cant be nil")

	}
	if bookname==nil{
		panic("Error:bookname cant be nil")
	}
	 fmt.Println("BookId:",*bookid,"BookName:",*bookname)
	fmt.Println("book function")
}
func Library(){
	id1:=1
	
	book(&id1,nil)
	fmt.Println("main defer")
}
