package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title: ")

	if err != nil {
		fmt.Println("Error: ", err)
		return "", "", err
	}
	content, err := getUserInput("Note body: ")
	return title, content, err

}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if value == "" {
		fmt.Println("You must enter a value")
		return "", errors.New("no value entered")
	}
	return value, nil

}
