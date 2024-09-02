package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, err := getUserInput("hello dear sir , what is your name : ", reader)
	if err != nil {
		fmt.Printf("an error has occured ! %v", err)
	}
	b := createNewBill(name)
	fmt.Printf("created bill - %v\n", b.name)
	return b
}
func prompOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, err := getUserInput("Select an option (a - add a new Item, s - save the bill, t - add a tip )", reader)
	if err != nil {
		fmt.Printf("an error has occured ! %v", err)
	}
	switch opt {
	case "a":
		item, _ := getUserInput("Item name : ", reader)
		price, _ := getUserInput("Item price : ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			prompOptions(b)
		}
		b.addItems(item, p)
		fmt.Printf("item added \n")
		prompOptions(b)

	case "s":
		b.saveBill()
		fmt.Println("you saved the bill : ", b.formatBill())

	case "t":

		tip, _ := getUserInput("Item tip : ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be a number")
			prompOptions(b)
		}
		b.updateTip(t)
		fmt.Printf("tip added : %v \n", t)
		prompOptions(b)

	default:
		fmt.Println("man choose a valid option or we'll increase the bill by £5")
		prompOptions(b)
	}

}
func main() {
	mybill := createBill()
	prompOptions(mybill)

	// fmt.Printf("\n\n\n the bill is %v", mybill.formatBill())
}
