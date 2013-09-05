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
    FileHandle os.File
    Tasks []Tasks
    LastError error
    NextId int
}

func main() {
    var usage string = "<[-l, -a, -f> [id]\n"
    if len(os.Args) > 1 {
        var file string = os.Getenv("HOME") + "/.config/t"
        // check to make sure our config exists
        if _, err := os.Stat(file); os.IsNotExist(err) { 
            fmt.Print("(!) No config found. Creating one for you now... ")
            if fh, ok := createConfig(file); ok {
                Task{FileHandle: fh}
                fmt.Println("OK, done. You're welcome")
            } else {
                fmt.Fprintf(os.Stderr, "We had trouble creating config. Bailing\n")
                os.Exit(2)
            }
        }
        
        t := Task{File: file}
        t.init()

        if ! t.parseOptions() {
            fmt.Fprintf(os.Stderr, "Usage: %s %s\n", usage, os.Args[0])
            os.Exit(3)
        }
    } else {
        fmt.Fprintf(os.Stderr, "Usage: %s %s\n", usage, os.Args[0])
        os.Exit(1)
    }
}

// function to initialise anything we might need
// like to grab the next available ID from our file
// or initialise Tasks and add them into the Task struct
func (t *Task) init() {
    //scanner := bufio.NewScanner(t.File)
}

func (t *Task) parseOptions() bool {
    var err error
    // remove the file from the start of the array
    args := append(os.Args[:0], os.Args[0+1:]...)
    
    // get the command
    switch(args[0]) {
        case "list":
            fmt.Println("Listing tasks")
        case "add":
            if t.addTask(args); err == nil {
                fmt.Println("Added task")
                return true
            } else {
                return false
            }
    }

    return true
}

func (t *Task) addTask(args []string) bool {
    if len(args) != 2 {
        return false
    }

    
    return true
}

func createConfig(f string) (file, bool) {
    fo, err := os.Create(f)
    if err != nil { panic(err) }

    // close the file, or panic if we can't
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
   
    return *fo, true
}
