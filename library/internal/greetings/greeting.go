package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string, LibraryName string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome to %v system", name, LibraryName)
    return message
}

