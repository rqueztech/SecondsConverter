package main

import (
    "fmt";      // for printing and formatting
    "strings";  // for string manipulations
    "strconv";  // typecasting and mainuplation of strings
    "bufio";    // for reading input from strings
    "os";       // for all io operations
)

type Builder struct {
    strings.Builder
}

func (builder *Builder) timesplitter(linkname string, timestosplit []string) [][]string { 
    var writearray [][]string = make([][]string, len(timestosplit))
    var currentvalue int = 0


    // iterate through every time that was put in the csv
    for index,currentTime := range timestosplit {
        var builder Builder

        // handles extra garbage character at end
        if len(currentTime) > 1 {
            // split the times by colon (:)
            timesplit := strings.Split(currentTime, ":")

            var timesplitlength int = len(timesplit)

            // convert the 3 inputs into hours, minutes, and seconds
            hours, _ := strconv.Atoi(timesplit[0])
            minutes, _ := strconv.Atoi(timesplit[1])
            seconds, _ := strconv.Atoi(timesplit[2])

            // calculate each component to extract the number of seconds
            var totalseconds int = (hours*3600) + (minutes*60) + (seconds)

            var totalsecondstring = strconv.Itoa(totalseconds)

            // write into the array position we created. we need the array to write to a new csv.
            builder.WriteString("<a href = https://www.")
            builder.WriteString(linkname)
            builder.WriteString("&t=")
            builder.WriteString(totalsecondstring)
            builder.WriteString("\">")

            if timesplitlength < 3 {
                writearray[index] = []string {builder.String(), "</a>"};
            } else {
                writearray[index] = []string {builder.String(), timesplit[3], "</a>"};
            }

            currentvalue++
        }
    }

    writeCSV(writearray)

    return writearray
}

func displayarray(writearray [][]string) {
     // stopping point: array prints right
    for _, col := range(writearray) {
        for _, row := range(col) {
            fmt.Println(row)
        }
    }   
}

func main(){
    var times_csv string = "resources/times.csv"
    var builder Builder

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
    var readValues string = ReadCSV(times_csv)

    // check if the data exists (may be redundant)
    if readValues != "" {
        fmt.Println("Data exits")
    } else {
        fmt.Println("Data does not exist")
    }

    // split the time values here
    var timestosplit []string = strings.Split(readValues, ",")

    // make an array at the length of timesplit.
    var writearray [][]string = builder.timesplitter(linkname, timestosplit)

    displayarray(writearray)
}
