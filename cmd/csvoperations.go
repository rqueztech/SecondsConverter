package main

import (
    "os";
    "fmt";
    "encoding/csv";
)

/*
func WriteCSV() {
     // file variable -> recieves pointer of csv file if found
    // err -> kicks error if this is not found.
    file, err := os.Open("resources/times.csv")

    // If the error is not null, meaning an error was found... then print out custom error. er is verbose, so we replace with our own
    if(err != nil) {
        fmt.Println("File Found...")
    } else {
        fmt.Println("File Not Found...")
    }   
}
*/

func ReadCSV() {
     // file variable -> recieves pointer of csv file if found
    // err -> kicks error if this is not found.
    file, err := os.Open("resources/times.csv")

    // If the error is not null, meaning an error was found... then print out custom error. er is verbose, so we replace with our own
    if(err != nil) {
        fmt.Println("File Found...")
    } else {
        fmt.Println("File Not Found...")
    }

    reader := csv.NewReader(file)
 

    // process ther records
    for _, record := range records {
        // process inner fields
        for _, field := range record {
            fmt.Println(field)
        }
    }
}
