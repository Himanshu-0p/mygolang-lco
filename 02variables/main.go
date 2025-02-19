package main



import (
	"fmt"

)

const token string = "qwertyjhgfd"

func main() {
	var username string = "himanshu"
	fmt.Println(username)
	fmt.Printf("Variables is of type:%T \n",username)

	var islogged bool = true
	fmt.Println(islogged)
	fmt.Printf("Variables is of type:%T \n",islogged)

	var smallVaL uint8 = 7
	fmt.Println(smallVaL)
	fmt.Printf("Variables is of type:%T \n",smallVaL)

	var smallFloat float64 = 7.123456789098765432
	fmt.Println(smallFloat)
	fmt.Printf("Variables is of type:%T \n",smallFloat)

	//default values and some aliases
	var anotherVariables int
	fmt.Println(anotherVariables)
	fmt.Printf("Variables is of type:%T \n",anotherVariables)

	//implicit type
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variables is of type:%T \n",website)

	//no var style
	numberOfuser := 10000
	fmt.Println(numberOfuser)

	fmt.Println(token)
	fmt.Printf("Variable of type:%T \n", token)
}
