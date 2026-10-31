package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

	for !quit {
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
			handlePosition(cmd)
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

func handleNewGame() {
	tell("nyi")
}

func handleIsReady() {
	tell("readyok")
}

func handleRegister() {
	tell("nyi")
}

func handleSetOption(option []string) {
	tell("not implemeneted", strings.Join(option, " "))
}

func handlePosition(cmd string) {
	// position [fen <fenstring> | startpos ] moves <move1> .... <movei>
	// separate the string into the current position and given move
	cmd = trim(strings.TrimPrefix(cmd, "position"))
	parts := strings.Split(cmd, "moves")
	if len(cmd) == 0 || len(parts) > 2 {
		err := fmt.Errorf("%v wrong length=%v", parts, len(parts))
		tell("info string Error", fmt.Sprint(err))
		return
	}
	alt := strings.Split(parts[0], " ")
	alt[0] = trim(alt[0])
	tell("info string position", alt[0], "not implemented")

	if alt[0] == "startpos" { //initialize the start position
		alt[0] = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	} else if alt[0] == "fen" {
		alt[0] = " "
	} else {
		err := fmt.Errorf("%v must be %#v or %#v", alt[0], "fen", "startpos")
		tell("info string Error", err.Error())
		return
	}
	parts[0] = strings.Join(alt, " ")

	//wish i could see how this part worked....
	parts[0] = trim(parts[0])
	fmt.Printf("info string prase %#v\n", parts[0])
	parseFEN(parts[0])

	if len(parts) == 2 {
		parts[1] = low(trim(parts[1]))
		fmt.Printf("info string parse %#v\n", parts[1])
		parseMvs(parts[1])
	}
}

func parseFEN(fen string) {
	fenIx := 0
	for row := 7; row >= 0; row-- {
		for sq := row * 8; sq < row*8+8; {

			char := string(fen[fenIx])
			fenIx++
			if char == "/" {
				continue
			}

			if i, err := strconv.Atoi(char); err == nil {
				fmt.Println(i, "empty from sq", sq)
				sq += i
				continue
			}

			fmt.Println(char, "on sq", sq)
			sq++
		}
	}
}

func parseMvs(move string) {
	mvs := strings.Split(move, " ")
	for _, mv := range mvs {
		fmt.Println("make move", mv)
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

func handleDebug(words []string) {
	tell("nyi")
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
