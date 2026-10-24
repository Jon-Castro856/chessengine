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

func uci(input chan string) { //this command accepts an input and executes functions based on the input received
	bInfinite := false
	frmEngine, toEngine := engine()
	quit := false
	cmd := "" //the command received by input
	bm := ""  //the current best move
	words := []string{}

	for quit == false {
		select {
		case cmd = <-input:
			tell("info string uci got ", cmd)
		case bm = <-frmEngine:
			handleBm(bm, bInfinite)
			continue
		}

		words = strings.Split(cmd, " ")
		words[0] = trim(low(words[0]))
		switch words[0] {
		case "uci":
			handleUCI()
		case "isready":
			handleIsReady()
		case "ucinewgame":
			handleNewGame()
		case "position":
			handlePosition(words)
		case "debug":
			handleDebug(words)
		case "register":
			handleRegister()
		case "go":
			handleGo(words)
		case "setoption":
			handleSetOption(words)
		case "stop":
			handleStop(toEngine, &bInfinite)
		case "quit", "q":
			quit = true
			continue

		}
	}

	tell("info string leaving uci")
}

func handleBm(bm string, bInfinite bool) {
	if bInfinite {
		saveBm = bm
	}
	tell(bm)
}
func handleStop(toEng chan string, bInfinite *bool) {
	if *bInfinite {
		if saveBm != "" {
			tell(saveBm)
			saveBm = ""
		}
		toEng <- "stop"
		*bInfinite = false
	}
}

func handleUCI() { //handle initial uci check from the engine
	tell("id name Astro")
	tell("id author JonCastro")
	tell("option name Threads type spin default 1 min 1 max 16") //# of threads to use when searching for a position
	tell("option name Hash type spin default 32 min 1 max 1024") //Hash table size (i dont know what that is)

	tell("uciok")
}
func handleIsReady() {
	tell("readyok")
}
func handleSetOption(option []string) {
	tell("not implemeneted", strings.Join(option, " "))
}

func handlePosition(words []string) {
	// position [fen <fenstring> | startpos ] moves <move1> .... <movei>
	if len(words) > 1 {
		words[1] = trim(low(words[1]))
		switch words[1] {
		case "startpos":
			tell("not yet inplemented")
		case "fen":
			tell("not yet implemented")
		default:
			tell("string pos ", words[1], "not implemented")
		}
	}
}

func handleGo(words []string) {
	if len(words) > 1 {
		words[1] = trim(low(words[1]))
		switch words[1] {
		case "searchmoves":
			tell("nyi")
		case "ponder":
			tell("nyi")
		case "wtime":
			tell("nyi")
		case "btime":
			tell("nyi")
		case "winc":
			tell("nyi")
		case "binc":
			tell("nyi")
		case "movestogo":
			tell("nyi")
		case "depth":
			tell("nyi")
		case "nodes":
			tell("nyi")
		case "movetimes":
			tell("nyi")
		case "mate": //mate <X> mate in X moves
			tell("nyi")
		case "infinite":
			tell("nyi")
		default:
			tell("nyi")
		}
	} else {
		tell("nyi")
	}
}

// this function starts a go routine that reads for inputs
func input() chan string {
	line := make(chan string)
	go func() {
		var reader *bufio.Reader
		reader = bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
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
