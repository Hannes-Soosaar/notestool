# notesTool

## Description

**notesTool** is a command-line tool(written in golang) that allows users to manage short single-line notes. Using this application, users can create collections of notes, open and view them, add new notes, or remove existing notes. The tool takes exactly one argument, which is the name of the collection.

## Notes

**The Notes** file is created in JSON type. **JSON** is a text-based data format following JavaScript object syntax. JSON exists as a string — useful when you want to transmit data across a network. In this case JSON string can be stored in its own file, which is basically just a text file with an extension of ```.json```.

### **When starting the program in terminal, you can create a new notes collection by adding to the command collection its name,if the collection already exists it just opens it and user interacts with it**:

```
go run notestool.go newcollection 
```
Then the **programm starts** and the collection is created, then user can interact with it by choosing the menu options.

If no argument is provided, the number of arguments is not one, or if the argument is help, the tool will display a brief help message that explains how to use the application.

```
$ ./notestool
Usage: ./todotool [TAG]
```
## Usage

To run the program you will need to follow these steps:
- Run the program
- When it starts you can see it the Welcome message
- Menu list suggests the user to choose one of the following opeartions(1. Show notes; 2. Add a note; 3. Delete a note; 4. Exit)

```
Welcome to the notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
```

- User enters "1" and the tool shows the existing list of notes
- Right after it was printed, it shows the main menu again - to let user choose oher options or exit

```
Notes:
001 - note one
002 - note two

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.

```

- User enters "2" and the tool suggest them to type in some note text
    
```
Enter the note text:
note three //for instance this is the user input
```

- Right after the operation is done, user is back to the main menu
- User can now check if the note was written to the notes' list by entering "1" again and will see the new note added, nd after will be back to the main menu

```
Notes:
001 - note one
002 - note two
003 - note three
```

- By choosing "3", user can delete the note out of the list by it's number written to the command line, and if they change their mind type in "0" and het returned to the main menu

```
Enter the number of note to remove or 0 to cancel:
```
- After the operation is done user is returned to the main menu, they can check the note to be deleted by choosing "1", and the edited notes list will be showed on the screen
```
Notes:
001 - note one
002 - note two
```
- After user is done, and wants to exit the tool they need to type in "4" and the tool is closed

### Build

Built needed to run the tool as a programm/app

To build the tool:

```
go build -o ./notestool
```

## Why this tool is needed

**Notes** can refer to a feature that allows users to jot down and store information, such as this tool.

...And much more:
1. Organization: Notes serve as an organizational tool, helping to structure and categorize information
2. Critical Thinking: The act of taking notes prompts critical thinking. It requires individuals to analyze information, extract key concepts, and make decisions about what to include
3. Goal Setting: Notes can be used to set and track goals
4. Creativity and Idea Generation: Notes provide a platform for recording ideas and fostering creativity



### Main purpose of this tool
In case of this tool notes are mostly jot down to better understand tool's code logic, learn how to use new functions and packages in golang. 

## noteTool authors

# Group Leader: **Hannes Soosaar**

# Team member: Antonina Krjukova