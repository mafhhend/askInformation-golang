package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

//this isn't variable, it's a type like 'int'
type User struct {
	userName        string
	Age             int
	FavouriteNumber float64
	OwnsADog        bool
}

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.userName = readString("what is your name ?")
	user.Age = readInt("How Old Are You??")
	user.FavouriteNumber = readFloat("What is your favourite number?")
	user.OwnsADog = readBool("Do you own a dog (y/n)")
	// S : string
	// f : format
	fmt.Println(fmt.Sprintf("Your name is %s and you are %d years old. your Favourite number is %.0f. Owns a dog: %t", user.userName, user.Age, user.FavouriteNumber, user.OwnsADog))
}

func prompt() {
	fmt.Println("->")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput != "" {
			return userInput
		} else {
			fmt.Println("Please enter a value")
		}
	}

}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please Enter a whole number")
		} else {
			return num
		}
	}
}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println("Please Enter a number")
		} else {
			return num
		}
	}
}

func readBool(s string) bool {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {

		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			fmt.Println("There is a problem !!!")
		}
		// var yesNo bool
		var yesNo string = strings.ToLower(string(char))
		fmt.Println(yesNo)
		if yesNo != "y" && yesNo != "n" {
			fmt.Println("Enter Y or N !!!")
		} else {
			if yesNo == "y" {
				return true
			} else {
				return false
			}
		}
	}
}
