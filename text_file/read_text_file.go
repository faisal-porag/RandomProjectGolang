package text_file

import (
    "fmt"
    "os"
    "bufio"
)

func ReadTextFileExample() {

    f, err := os.Open("data.txt")

    if err != nil {
        fmt.Println(err)
     }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {

        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}
