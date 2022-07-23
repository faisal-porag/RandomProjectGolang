package text_file

import (
    "fmt"
    "log"
    "os"
)

func CreateTextFileExample() {

    f, err := os.Create("data.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    val := "old falcon"
    data := []byte(val)

    _, err2 := f.Write(data)

    if err2 != nil {
        log.Fatal(err2)
    }

    val2 := " and red fox\n"
    data2 := []byte(val2)

    var idx int64 = int64(len(data))

    _, err3 := f.WriteAt(data2, idx)

    if err3 != nil {
        log.Fatal(err3)
    }

    fmt.Println("done")
}
