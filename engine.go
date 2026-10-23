package main

func engine() (froEngine, toEngine chan string) {
	froEngine = make(chan string)
	toEngine = make(chan string)
	go func() {
		for cmd := range toEng {
			switch cmd {
			case "stop":
			case "quit":

			}
		}
	}()
	return
}
