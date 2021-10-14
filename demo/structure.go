package demo
import "fmt"
type rectangle struct{
	length float64
	breath float64
	color string

}
type Semestermarks struct{
	sgpa float64
	cgpa float64
}
type Student struct{
	name string
	usn string
	marks []Semestermarks

}
func Area(a,b float64)(float64){
var r1 rectangle
r1.length=a
r1.breath=b
area:=r1.length*r1.breath
return area
}
func Example(){
	var e[] Student=[]Student{
		Student{
		name:"Srikantha",
		usn:"4cb12cs032",
		marks:[]Semestermarks{
			Semestermarks{
				sgpa:7.0,
				cgpa:7.0},
			Semestermarks{
				sgpa:3.0,cgpa:7.0},
			Semestermarks{
				sgpa:8.0,
				cgpa:6.0},
			},

		},Student{
			name:"Adarsh",
			usn:"4cb12cs031",
			marks:[]Semestermarks{
				Semestermarks{
					sgpa:4.0,
					cgpa:5.0},
				Semestermarks{
					sgpa:8.0,cgpa:9.0},
				Semestermarks{
					sgpa:9.0,
					cgpa:8.0},
				},
	
			},
		}
		for i:=range e{
		fmt.Println("Name:",e[i].name,"USN:",e[i].usn)
		for j:=range e[i].marks{
			fmt.Println("SEM:",j+1,"sgpa:",e[i].marks[j].sgpa,"& cgpa:",e[i].marks[j].cgpa)
		}
	}
	
}
