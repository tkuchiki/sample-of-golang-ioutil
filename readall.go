package main

import (
        "bufio"
        "fmt"
        "io/ioutil"
        "os"
)

func main() {
        fp, err := os.Open(os.Args[1])
        if err != nil {
                panic(err)
        }
        defer fp.Close()

        r := bufio.NewReader(fp)

        fileByte, err := ioutil.ReadAll(r)

        fmt.Printf("Type: %T \n\n", fileByte)
        fmt.Printf("%v", string(fileByte))
}
