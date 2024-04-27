package main

import (
    "bytes"
    "os";
    "fmt";
    "encoding/csv";
)

func ReadCSV() (string) {
     // file variable -> recieves pointer of csv file if found
    // err -> kicks error if this is not found.
    file, err := os.Open("resources/times.csv")

    // If the error is not null, meaning an error was found... then print out custom error. er is verbose, so we replace with our own
    if(err != nil) {
        fmt.Println("File Not Found...")
    } else {
        fmt.Println("File Found...")
    }

    // Create buffer to aggregate all input
    var buffer bytes.Buffer

    // Create the reader to take input
    reader := csv.NewReader(file)

    // read all of the records  out
    records, err := reader.ReadAll()

    // do error check
    if err != nil {
        fmt.Println(err)
    } else {
        // process ther records
        for _, record := range records {
            // process inner fields
            for _, field := range record {
                buffer.WriteString(field)
                buffer.WriteString(",")
            }
        }
    }

    return buffer.String()
}
