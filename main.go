package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
	_ "github.com/joho/godotenv/autoload"
)

const (
	Version = "1.0.0"
	base    = "https://api.github.com/repos/"
)

type Label struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func main() {
	var user string
	var repo string
	var token string
	var filePath string
	var labels []Label

	resource := "labels"
	app := cli.NewApp()
	app.Name = "Github Label Maker"
	app.Author = "Ernie Casilla"
	app.Version = Version
	app.Email = "ecasilla@icloud.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "A file path to your labels. i.e -f labels.json",
			EnvVar:      "LABELS_FILE",
			Destination: &filePath,
		},
		cli.StringFlag{
			Name:        "user, u",
			Usage:       "Your Github username -u octocat",
			EnvVar:      "GITHUB_USER",
			Destination: &user,
		},
		cli.StringFlag{
			Name:        "token, t",
			Usage:       "Your Github OAuth Token -t github_token",
			EnvVar:      "GITHUB_TOKEN",
			Destination: &token,
		},
		cli.StringFlag{
			Name:        "repo, r",
			Usage:       "The GITHUB repo you want to add labels to -r octocat_repo",
			EnvVar:      "GITHUB_REPO",
			Destination: &repo,
		},
	}

	app.Run(os.Args)

	url := base + user + "/" + repo + "/" + resource + "?access_token=" + token

	if filePath == "" {
		return
	}

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Printf("error %v", err)
		panic(err)
	}

	if err := json.Unmarshal(file, &labels); err != nil {
		fmt.Printf("error %v", err)
		panic(err)
	}

	if err := createLabels(url, labels); err != nil {
		fmt.Printf("error %v", err)
		panic(err)
	}
}
func createLabels(url string, l []Label) error {
	for _, val := range l {
		fmt.Println("URL:>>>", url)

		b, err := json.Marshal(val)
		if err != nil {
			return err
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
		if err != nil {
			fmt.Printf("error %v", err)
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("error %v", err)
			return err
		}

		defer resp.Body.Close()

		fmt.Println()
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("Response Body:", string(body))
		fmt.Println()
	}
	return nil
}
