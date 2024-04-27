package main

import (
    "fmt";      // for printing and formatting
    "strings";  // for string manipulations
    "strconv";  // typecasting and mainuplation of strings
    "bufio";    // for reading input from strings
    "os";       // for all io operations
)

func main(){

    fmt.Print("Enter the link name: ")

    // Create a builder to compile the seconds
    reader := bufio.NewReader(os.Stdin)
    linkname, err := reader.ReadString('\n')

    // check reader for errors
    if err != nil {
        fmt.Println("Error reading")
    }

    // trim space to handle issue with newreader reading in the newline
    linkname = strings.TrimSpace(linkname)

    // read the csv
    readValues := ReadCSV()

    // check if the data exists (may be redundant)
    if readValues != "" {
        fmt.Println("Data exits")
    } else {
        fmt.Println("Data does not exist")
    }

    // split the time values here
    thisis := strings.Split(readValues, ",")

    // iterate through every time that was put in the csv
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

            fmt.Printf("<a href = \"%s&t=%ds\"></a>\n", linkname, totalseconds)
        }
    }

    fmt.Println("\nend")
}
