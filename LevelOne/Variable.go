package levelOne

import "fmt"

func DefaultValue() {
	var number int
	var float float32
	var useName string
	var boolean bool
	fmt.Println(number, float, useName, boolean)
}

func Variable() {
	var name string = "john"
	var age int16 = 20
	var qualification string = "B.E computer science"
	var CGPA float32 = 8.22
	user := fmt.Sprintf("My name is %s.I am a %d-year-old boy with a %s and a CGPA of %.1f.", name, age, qualification, CGPA)
	fmt.Println(user)

}

// statically typed
func ShortHandProperty() {
	name := "alice"
	age := 20
	qualification := "B.E computer science"
	CGPA := 7.33
	user := fmt.Sprintf("My name is %s.I am a %d-year-old boy with a %s and a CGPA of %.1f.", name, age, qualification, CGPA)
	fmt.Println(user)
}
