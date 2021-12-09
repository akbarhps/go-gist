package main

import (
	"github.com/atotto/clipboard"
	"log"
	"path/filepath"
)

var gist *Gist
var args *Args

func init() {
	args = parseArgs()
	gist = &Gist{}
}

func main() {
	for key, value := range args.StateData {
		gistRequest := NewGistRequest()

		switch key {
		case stateCreateFromDir:
			for _, dir := range value {
				err := gistRequest.AddContentFromDir(filepath.Join(args.ExecDir, dir)); if err != nil {
					log.Fatal(err)
				}
				createAndPrintGist(gistRequest)
			}

		case stateCreateFromFile:
			for _, file := range value {
				err := gistRequest.AddContentFromFile(args.ExecDir, file); if err != nil {
					log.Fatal(err)
				}
				createAndPrintGist(gistRequest)
			}

		case stateCreateFromFileMerged:
			for _, file := range value {
				err := gistRequest.AddContentFromFile(args.ExecDir, file); if err != nil {
					log.Fatal(err)
				}
			}
			createAndPrintGist(gistRequest)

		case stateCreateFromClipboard:
			text, err := clipboard.ReadAll()
			if err != nil {
				log.Fatal(err)
			}

			if text == "" {
				log.Fatal("Clipboard is empty")
			}

			gistRequest.Files[value[0]] = GistRequestFile{text}
			createAndPrintGist(gistRequest)

		case stateDelete:
			for _, id := range value {
				err := gist.Delete(id); if err != nil {
					log.Fatal(err)
				}

				log.Printf("Gist %s deleted", id)
			}

		}
	}
}

func createAndPrintGist(gistRequest *GistRequest) {
	gistResponse, err := gist.Create(gistRequest)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(gistResponse.ToString())

	if args.SaveToClipboard {
		err = clipboard.WriteAll(gistResponse.ToString())
		if err != nil {
			log.Fatal(err)
		}
	}
}
