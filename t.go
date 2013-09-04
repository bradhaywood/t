package main

import (
    "fmt"
    "os"
)

type Tasks struct {
    Id int
    Desc string
    Started string 
}

type Task struct {
    File string
    Tasks []Tasks
    LastError error
}

func main() {
    var usage string = "<[-l, -a, -f> [id]\n"
    if len(os.Args) > 1 {
        var file string = os.Getenv("HOME") + "/.config/t"
        // check to make sure our config exists
        if _, err := os.Stat(file); os.IsNotExist(err) { 
            fmt.Print("(!) No config found. Creating one for you now... ")
            if createConfig(file) {
                fmt.Println("OK, done. You're welcome")
            } else {
                fmt.Fprintf(os.Stderr, "We had trouble creating config. Bailing\n")
                os.Exit(2)
            }
        }
        
        t := Task{File: file}
        if t.parseOptions() {
            fmt.Println("Parsed fine!")
        } else {
            fmt.Fprintf(os.Stderr, "Usage: %s %s\n", usage, os.Args[0])
            os.Exit(3)
        }
    } else {
        fmt.Fprintf(os.Stderr, "Usage: %s %s\n", usage, os.Args[0])
        os.Exit(1)
    }
}

func (t *Task) parseOptions() bool {
    var flag string
    var i int = 0
    // remove the file form the start of the array
    args := append(os.Args[:0], os.Args[0+1:]...)
    for {
        i += 1
        // bail if we exceed the index max
        if i > len(args) { break }
        args = append(args[:0], args[i+0:]...)
        flag = args[0]
        switch(flag) {
            case "-l":
                fmt.Println("Listing tasks")

            case "-a":
                fmt.Println("Add a task")
        }
    }
    
    return true
}

func createConfig(f string) bool {
    fo, err := os.Create(f)
    if err != nil { panic(err) }

    // close the file, or panic if we can't
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    return true
}
