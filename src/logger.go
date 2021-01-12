package main

import (
    "os"
    "fmt"
)

func empty() {
    fmt.Fprintf(os.Stdout, "\n")
}

func debug(format string, a ...interface{}) {
	if *quiet == false {
		if *verbose {
	    	fmt.Fprintf(os.Stdout, "\x1b[0mdebug:\x1b[0m "+format+"\n", a ...)
	    }
	}
}

func info(format string, a ...interface{}) {
	if *quiet == false {
	    fmt.Fprintf(os.Stdout, "\x1b[2mlog:\x1b[0m "+format+"\n", a ...)
	    //fmt.Fprintf(os.Stdout, format+"\n", a ...)
	}
}

func warn(format string, a ...interface{}) {
	if *quiet == false {
	    fmt.Fprintf(os.Stdout, "\x1b[1mwarning:\x1b[0m "+format+"\n", a ...)
	}
}

func error(format string, a ...interface{}) {
    fmt.Fprintf(os.Stdout, "\x1b[91merror:\x1b[0m "+format+"\n", a ...)
    os.Exit(1)
}
