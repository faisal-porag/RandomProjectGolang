package text_file

import (
    "fmt"
    "os"
    "bufio"
    "log"
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

func ReadExample2() {

    f, err := os.Open("data.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)

    for scanner.Scan() {

        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
