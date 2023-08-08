package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
)

var (
	gh_authtoken string
	gh_user      string
	gh_repos     string
	gh_relpath   string
	gh_ref       string
)

// gittar --gh.token="token" --gh.user="user" --gh.repo="repos" --gh.path="path/to/file.type" --gh.ref="branch/tag/sha"

func main() {

	flag.StringVar(&gh_authtoken, "gh.token", "", "The Token to authorize against the github api.")
	flag.StringVar(&gh_user, "gh.user", "", "The User to authorize against the github api.")
	flag.StringVar(&gh_repos, "gh.repo", "", "The GitHub Repository.")
	flag.StringVar(&gh_relpath, "gh.path", "", "Root-relative path to the file.")
	flag.StringVar(&gh_ref, "gh.ref", "", "Branch/Tag/SHA to the file.")

	flag.Parse()

	log.Println("starting file upload")

	var client = github.NewClient(nil)

	fileContent, _, resp, err := client.Repositories.GetContents(context.TODO(), gh_user, gh_repos, gh_relpath, &github.RepositoryContentGetOptions{Ref: gh_ref})
	if err != nil {
		log.Printf("error accessing file content: %#v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Response from GitHub: %d - %s | %s\n", resp.StatusCode, resp.Status, resp.String())

	content, err := fileContent.GetContent()
	fmt.Printf("Content: %s\n", content)
	fmt.Printf("Error: %s\n", err)

	os.Exit(0)
}
