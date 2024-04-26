package main

import (
    "os";
    "fmt";
)

func CheckForCSV() bool {
     // file variable -> recieves pointer of csv file if found
    // err -> kicks error if this is not found.
    file, err := os.Open("resources/times.csv")

    // If the error is not null, meaning an error was found... then print out custom error. er is verbose, so we replace with our own
    if(err != nil) {
        fmt.Println("File Found...")
        return true
    } else {
        fmt.Println("File Not Found...")
        return false
    }   
}

func ReadCSV() {
    CheckForCSV()
}

func WriteCSV() {
    CheckForCSV()
}
