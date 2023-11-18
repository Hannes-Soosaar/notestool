package main

/* Version 0.0 of the "notes tool" created by Hannes Soosaar and Antonina Krjukova
comissioned by the 2023 Kood/Johvi sprints second raid*/

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	SHOW_NOTE        = 1
	ADD_NOTE         = 2
	DELETE_NOTE      = 3
	EXIT             = 4
	WRONG_INPUT      = "Error reading input:"
	ERR_OPENING_FILE = "Error opening file:"
	NOTE_NOT_SAVED   = "Error the note was not saved. Error:"
	NOTES_FILE_PATH  = "notes.json"
)

type Note struct {
	Index    int
	UserNote string
}

var noteIdIndex int = 1 // starts the count from 1

func WelcomeScreen() {
	fmt.Println("Welcome to the notes tool!")
}
func UsageScreen() {
	fmt.Println(" Usage: ./notesTool [tag]")
}
func DisplayDeleteMenu() {
	fmt.Println("Enter the number of note to remove or 0 to cancel:")
}
func DisplayWrongSellection() {
	fmt.Println("\nInvalid sellection! Please enter the number of the operation!")
}
func DisplayDeleteNoteMenu() {
	fmt.Println("\n Enter the number of note to remove or 0 to cancel:")
}
func DispalyNotes() {
	fmt.Println("\nNote:")
}
func DispalyAddNoteMenu() {
	fmt.Println("\nEnter the note text:")
}
func DisplayMainMenu() {
	fmt.Println("\nSelect operation:\n1. Show notes.\n2. Add a note.\n3. Delete a note.\n4. Exit.")
}
func ThrowError(errorText string, newError error) { // creatr error text to print
	fmt.Println(errorText, newError)
}
func MenuSelectInput() int {
	var selectedNumber = 0
	_, err := fmt.Scan(&selectedNumber) // _, palce holder could be used for numebr or items scanned, err stores the eror
	if err != nil || selectedNumber > 4 {
		DisplayWrongSellection()
		DisplayMainMenu()
		MenuSelectInput()  // if there is an error it will start the function over again
		selectedNumber = 0 // restes the selectednumber
	}
	return selectedNumber
}
func SellectOperation(userOperation int) {
	switch userOperation {
	case SHOW_NOTE:
		GetNotes()
	case ADD_NOTE:
		AddNewUserNoteToFile()
	case DELETE_NOTE:
		DeleteOperationSellection(GetNoteIndexDelete())
	case EXIT:
		return
	}
	userOperation = 0 // resets the userInput
}
func InitializeNoteIndex() {
	var notes []Note
	notesFromFile, err := ReadNotesFromFile() //T
	if err != nil {
		ThrowError("something is wrong with reading the file ", err)
	}
	notes = append(notes, notesFromFile...) // need to append the elements form the file to this temp variable
	noteIdIndex = len(notes) + 1            // starts count from one
}
func GetNotes() {
	var notes []Note
	DispalyNotes() // Prtinst the notes to console
	notesFromFile, err := ReadNotesFromFile()
	if err != nil {
		ThrowError("something is wrong with reading the file ", err)
	}
	notes = append(notes, notesFromFile...) // need to append the elements form the file to this temp variable
	for i := 0; i <= len(notes)-1; i++ {    // loops through the notes
		fmt.Printf("%03d %s \n", notes[i].Index, notes[i].UserNote) // Prints the notes to screen
	}
}
func ReadNotesFromFile() ([]Note, error) {
	namecol := os.Args[1]
	var notes []Note                                                       // creates a slice of notes
	file, err := os.OpenFile(namecol+".json", os.O_RDWR|os.O_CREATE, 0644) // save the file at the path to a custom type to be used later
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()                                              // says that we will close the file once we exit this function The opens lives only in this function as soon as it exits it dies
	decoder := json.NewDecoder(file)                                // decodes the  json
	if err := decoder.Decode(&notes); err != nil && err != io.EOF { // EOF is end of file
		ThrowError("error decoding", err)
	}
	return notes, nil
}
func GetUserNoteToAdd() Note {
	inputReader := bufio.NewScanner(os.Stdin)
	inputReader.Scan()
	line := inputReader.Text()
	return Note{noteIdIndex, line}
}
func AddNewUserNoteToFile() {
	DispalyAddNoteMenu()           // just prints stuff to the consol
	userNote := GetUserNoteToAdd() //read real text for real, man
	PostUserNoteToFile(userNote)
	noteIdIndex++
}
func PostUserNoteToFile(userNote Note) {
	namecol := os.Args[1]                                                  //
	file, err := os.OpenFile(namecol+".json", os.O_RDWR|os.O_CREATE, 0644) // save the file at the path to a custom type to be used later
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var notes []Note                                                       // creates a slice of notes
	decoder := json.NewDecoder(file)                                       // decodes the  json
	if err := decoder.Decode(&notes); err != nil && err.Error() != "EOF" { // decodes the file to the address of notes
		ThrowError("error decoding", err)
	}
	notes = append(notes, userNote)
	file.Seek(0, 0)  // finds the begining of the file
	file.Truncate(0) // empties the file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(notes); err != nil {
		ThrowError("Error encoding JSON:", err)
		return
	}
}

func DeleteOperationSellection(IndexToDelete int) {
	namecol := os.Args[1]
	if IndexToDelete == 0 {
		return
	}
	var notes []Note
	file, err := os.OpenFile(namecol+".json", os.O_RDWR|os.O_CREATE, 0644) // save the file at the path to a custom type to be used later
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&notes); err != nil && err.Error() != "EOF" { // decodes the file to the address of notes
		ThrowError("error decoding", err)
		if IndexToDelete > len(notes) || IndexToDelete < 1 {
			fmt.Printf("There's no note with the Index %d", IndexToDelete)
			return
		}
	}
	for i := 0; i < len(notes); i++ {
		if notes[i].Index == IndexToDelete {
			notes = append(notes[:i], notes[i+1:]...)
		}
	}
	file.Seek(0, 0)  // finds the begining of the file
	file.Truncate(0) // empties the file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(notes); err != nil {
		ThrowError("Error encoding JSON:", err)
		return
	}
}
func GetNoteIndexDelete() int {
	var selectedNumber = 0
	DisplayDeleteMenu()                 // enter func to display selection menu to delete
	_, err := fmt.Scan(&selectedNumber) // _, palce holder could be used for numebr or items scanned, err stores the eror
	if selectedNumber == 0 {
		return 0
	}
	if err != nil || selectedNumber < 0 {
		DisplayWrongSellection()
		GetNoteIndexDelete() // enter func to display selection menu to delete
		selectedNumber = 0   // restes the selectednumber
	}
	return selectedNumber
}
func main() {
	if len(os.Args) <= 1 || os.Args[1] == "help" {
		UsageScreen()
		return
	}
	var userOption int
	InitializeNoteIndex()    // sets the indexId to be one bigger than the last slice piece
	WelcomeScreen()          // method to display the welcome text
	for userOption != EXIT { // main loop
		userOption = 0    // resets the user option
		DisplayMainMenu() // a func that prints the menu text
		userOption = MenuSelectInput()
		SellectOperation(userOption) // sellects an operation based on the userOption sellected
	}
}
