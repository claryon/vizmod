package common

import "io/ioutil"
import "fmt"

func ScreenWrite(message string) {
        fmt.Printf("vizmod: %s\n", message)
}

func StateFlush(name string, contents string) {
        b := []byte(contents + "\n")
        name = "vizmod-" + name
        err := ioutil.WriteFile(name, b, 0644)
        if err != nil {
                fmt.Printf("\nivizmod: demonstrator visualization failed to write to flush file %s: %s\n", name, contents)
                panic(err)
        }
}
