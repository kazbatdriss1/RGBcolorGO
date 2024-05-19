package main

import (
    "fmt"
)

// Function to return the ANSI escape code for an RGB foreground color
func rgb(r, g, b int) string {
    return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

// Function to return the ANSI escape code for an RGB background color
func bgRgb(r, g, b int) string {
    return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

// Reset code to return to default colors
const Reset = "\033[0m"

func main() {
    fmt.Println(rgb(255, 0, 0) + "This is red text" + Reset)
    fmt.Println(rgb(0, 255, 0) + "This is green text" + Reset)
    fmt.Println(rgb(0, 0, 255) + "This is blue text" + Reset)
    fmt.Println(rgb(128, 0, 128) + "This is purple text" + Reset)
    fmt.Println(bgRgb(255, 255, 0) + "This text has a yellow background" + Reset)
    fmt.Printf("%sThis is text with RGB foreground and %sbackground%s\n", rgb(255, 165, 0), bgRgb(0, 255, 255), Reset)
}

