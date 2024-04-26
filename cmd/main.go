package main

import (
    "fmt";
)

func main(){
    readValues, err := ReadCSV()
    //WriteCSV()

    if err != nil {
        fmt.Println("CSV not working...")
    } else {
        fmt.Println(readValues)
    }
}
