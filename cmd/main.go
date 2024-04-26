package main

import (
    "fmt";
    "strings";
    "strconv";
    "bufio";
    "os";
)

// Create main menu for user prompt
func mainmenu() {
    fmt.Println("1. Link To Aggregate") 
    fmt.Println("2. Raw Times")
    fmt.Println("\nEnter Option: ")
}

func main(){
    // print out the main menu here
    mainmenu()
    // Create a builder to compile the seconds
    reader := bufio.NewReader(os.Stdin)

    // scan in user input
    option, err := reader.ReadString('\n')

    // Handle error for user scan
    if err != nil {
        fmt.Println("No error")
    } else {
        fmt.Println("Error exists")
    }

    fmt.Println(option)

    readValues := ReadCSV()
    //WriteCSV()

    if readValues != "" {
        fmt.Println("Data exits")
    } else {
        fmt.Println("Data does not exist")
    }
    

    fmt.Println(readValues)
    thisis := strings.Split(readValues, ",")

    for _,currentTime := range thisis {
        // handles extra garbage character at end
        if len(currentTime) > 1 {
            // split the times by colon (:)
            timesplit := strings.Split(currentTime, ":")

            // convert the 3 inputs into hours, minutes, and seconds
            hours, _ := strconv.Atoi(timesplit[0])
            minutes, _ := strconv.Atoi(timesplit[1])
            seconds, _ := strconv.Atoi(timesplit[2])

            totalseconds := (hours*3600) + (minutes*60) + (seconds)

            fmt.Println(currentTime, ",", totalseconds)
        }
    }

    fmt.Println("end")
}
