package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/custom-type/note"
)

/*
type str string

func (text *str) log() {
	fmt.Print(*text)
}

func main() {
	var name str = "pratik"
	name.log()
}

*/

func main() {
	title, content:= getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving failed")
		return
	}

	fmt.Println("Savded successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content :=getUserInput("Note content: ")
	
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	/*
	var value string
	fmt.Scanln(&value) // leaves a \n in the buffer next scan consumes it
	*/
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}