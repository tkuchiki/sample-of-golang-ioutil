package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
)

func main() {
        link := "http://example.com/"

        client := &http.Client{}

        response, err := client.Get(link)
        if err != nil {
                fmt.Println(err)
        }
        defer response.Body.Close()

        content, err := ioutil.ReadAll(response.Body)

        if err != nil {
                fmt.Println(err)
        }

        fmt.Println(string(content))

}

