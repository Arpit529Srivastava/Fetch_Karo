package main

import (
    "fmt"
    "log"
    //"os"
    "io/ioutil"

    fb "github.com/huandu/facebook/v2"
)

func main() {

    // Fetch Facebook data
    resp, err := fb.Get("/122096451866545061", fb.Params{
        "fields":       "name",
        "access_token": "EAAH7DLup7DsBO3suruh8ziXuq9yhGk1wSjZCDcPILn7rqZA7UFpfZBYOLM3r1oCnz4e1YXWZBTatj2taz9svDgZBRJ3UU1n79wL4RBcnsZB7yVLrYEbbzagsczyU3ai2ySDFVxZAVzjVDZBz0YVrSXn6uCduZB5Fv5bRLO7nZCWO6xyfoSsZCQOW8SOtUXIFvLoQpDZBboMdHANuZApC2Gk5XlGDZCdY5UAYKGy8hd5SEZD",
    })
    if err != nil {
        log.Fatalf("Failed to fetch Facebook data: %v", err)
    }
    name, ok := resp["name"].(string)
    if !ok {
        log.Fatalf("Failed to parse name from response")
    }
    output := fmt.Sprintf("Here is my Facebook name: %s\n", name)
    err = ioutil.WriteFile("id.txt", []byte(output), 0644)
    if err != nil {
        log.Fatalf("Failed to write to file: %v", err)
    }

    fmt.Println("Facebook name has been written to id.txt")
}
