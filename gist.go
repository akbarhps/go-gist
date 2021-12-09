package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type GistResponse struct {
	Id          string `json:"id"`
	IsPublic    bool   `json:"public"`
	Description string `json:"description"`
	APIUrl      string `json:"url"`
	HTMLUrl     string `json:"html_url"`
	ForksUrl    string `json:"forks_url"`
	CommitsUrl  string `json:"commits_url"`
	CommentsUrl string `json:"comments_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Message     string `json:"message"`
}

func (res *GistResponse) ToString() string {
	return fmt.Sprintf("Id: %s | Url: %s\n", res.Id, res.HTMLUrl)
}

type GistRequest struct {
	Description string                     `json:"description"`
	IsPublic    bool                       `json:"public"`
	Files       map[string]GistRequestFile `json:"files"`
}

func NewGistRequest() *GistRequest {
	return &GistRequest{
		Description: "",
		IsPublic:    false,
		Files:       make(map[string]GistRequestFile),
	}
}

func (req *GistRequest) AddContentFromDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		err = req.AddContentFromFile(dir, file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func (req *GistRequest) AddContentFromFile(execDir, fileName string) error {
	path := filepath.Join(execDir, fileName)
	escapedFileName := filepath.Base(path)

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	req.Files[escapedFileName] = GistRequestFile{Content: string(fileContent)}
	return nil
}

type GistRequestFile struct {
	Content string `json:"content"`
}

type Gist struct {
}

var gistUrl = "https://api.github.com/gists"

func (g *Gist) Create(req *GistRequest) (*GistResponse, error) {
	requestBuffer, err := InterfaceToBuffer(req)
	if err != nil {
		return nil, err
	}

	responseBytes, err := Fetch(http.MethodPost, gistUrl, requestBuffer)
	if err != nil {
		return nil, err
	}

	response := &GistResponse{}
	err = BytesToInterface(responseBytes, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (g *Gist) Delete(gistId string) error {
	deleteUrl := fmt.Sprintf("%s/%s", gistUrl, gistId)

	_, err := Fetch(http.MethodDelete, deleteUrl, nil)
	if err != nil {
		return err
	}

	return nil
}
