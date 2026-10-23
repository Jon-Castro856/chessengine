package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var tell = mainTell //sets tell to MmainTell by default
var saveBm = ""

// helper functions for more quickly writing repetitive functions
var trim = strings.TrimSpace
var low = strings.ToLower

func uci(input chan string) {
	fmt.Println("sup")

	frmEngine, toEngine := engine()
	quit := false
	cmd := ""

	for quit == false {
		select {
		case cmd = <-fromGUI:
		case bm = <-frmEngine:
			handleBm(bm)
			continue
		}
		switch cmd {
		case "uci":
		case "stop":
			handleStop(toEngine)
		case "quit", "q":
			quit = true
			continue

		}
	}
}

func handleBm(bm string) {
	tell(bm)
}
func handleStop(toEng chan string) {
	toEng <- "stop"
}

// this function starts a go routine that reads for inputs
func input() chan string {
	line := make(chan string)
	go func() {
		var reader *bufio.Reader
		reader = bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString("\n")
			if err != io.EOF && len(text) > 0 {
				line <- text
			}
		}
	}()
	return line
}

// text to send back to the GUI
func mainTell(text ...string) {
	toGUI := ""
	for _, t := range text {
		toGUI += t
	}
	fmt.Println(toGUI)
}
