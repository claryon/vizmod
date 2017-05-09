package common

import "fmt"
import "io/ioutil"
import "os"
import "time"

func StateFlush(name string, contents string) {
	flush(name, contents)
	log(name, contents)
	log("full", name + ": "+ contents)
}

func flush(name string, contents string) {
        b := []byte(contents + "\n")
        name = "vizmod-" + name + ".flush"
        err := ioutil.WriteFile(name, b, 0644)
	check(err, name, contents)
}

func log(name string, contents string) {
        s := timestamp() + " " + contents + "\n"
        name = "vizmod-" + name + ".log"
	f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err, name, contents)
	defer f.Close()
	_, err = f.WriteString(s)
	check(err, name, contents)
}

func timestamp() string {
	return time.Now().Format("02 Jan 06 15:04:05")
}

func check(err error, file string, contents string) {
	if err != nil {
        	fmt.Printf("\nvizmod: demonstrator visualization failed to write to flush file %s: %s\n", file, contents)
                panic(err)
        }
}
