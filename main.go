package main

import (
	"log"
	"os"
)

var (
	stateDelete = "DELETE"
	stateCreate = "CREATE"
)

func main() {
	executeDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	gist := &Gist{}

	var currentState string
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}

		switch arg {
		case "-c":
			currentState = stateCreate
			continue
		case "-d":
			currentState = stateDelete
			continue
		}

		switch currentState {
		case stateCreate:
			request := NewGistRequest()
			err = request.AddContent(executeDir, arg)
			if err != nil {
				log.Fatal(err)
			}

			gistResponse, err := gist.Create(request)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(gistResponse.ToString())
		case stateDelete:
			err = gist.Delete(arg)
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("Gist %s deleted", arg)
		}
	}

	if currentState == "" {
		log.Println(`HELP:

to create gist, use:
	gist -c file1 file2 file3

to delete gist, user:
	gist -d gistId1 gistId2

to do both, use:
	gist -c file1 file2 -d gistId1 gistId2`)
	}
}
