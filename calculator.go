package main

import (
	"fmt" //print
	"bufio" //input by terminal
	//"log" //treat and show error
	"os"
	"strconv" //convert string to another type
	//"strings"
)

func main() {
		n1 := askNumber()
		oper := askOperation()
		n2 := askNumber()


		var result string

		switch oper {
			case "+":
				result = addition(n1,n2)
			case "-":
				result = subtraction(n1,n2)
			case "x":
				result = multiplication(n1,n2)
			case "/":
				result = division(n1,n2)
			default:
				fmt.Println("Input a Valid Operator")
				panic("exit")
			return
		}
		fmt.Println("The operation calculated is:",n1,oper,n2," = ", result)

}

func addition(n1, n2 float64) string{
	sum := n1 + n2

	return strconv.FormatFloat(sum, 'f', 2, 64)
}
func subtraction(n1, n2 float64) string{
	minus := n1 - n2

	return strconv.FormatFloat(minus, 'f', 2, 64)
}
func multiplication(n1, n2 float64) string{
	times := n1 * n2

	return strconv.FormatFloat(times, 'f', 2, 64)
}
func division(n1, n2 float64) string{
	divide := n1 / n2

	return strconv.FormatFloat(divide, 'f', 2, 64)
}

func askOperation() (operation string){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please, enter the operation(+,-,x,/): ")
	scanner.Scan()
	//operation = strings.TrimRight(scanner.Text(), "\n")
	operation = scanner.Text()

	return 
}

func askNumber() (number float64){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please, enter the number: ")
	scanner.Scan()
	number,_ = strconv.ParseFloat(scanner.Text(), 64)

	return 
}