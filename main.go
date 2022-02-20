package main

import (
    "fmt"
    "math/rand"
    "time"
    "os"
    "bufio"
    "strconv"
)

func main() {
    // Nice random number setup
    rand.Seed(time.Now().UnixNano())
    var random = rand.Intn(10);

    for true {
        // Read a line from the console
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter something: ");line, _ := reader.ReadString('\n')

        // Check if that number is equal to a const
        value, _  := strconv.Atoi(line) // _ is error
        if random == value {
            fmt.Println("You got it mister");
            os.Exit(0)
        } else {
            fmt.Println("You did not get it mister");
        }
    }
}
