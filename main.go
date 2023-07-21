package main

import "os"

type Inputs struct {
	tag string
	ref string
	sha string
}

func main() {
	inputs := getInputs()
}

func getInputs() Inputs {
	return Inputs{
		ref: os.Getenv("GITHUB_REF"),
		sha: os.Getenv("GITHUB_SHA"),
	}
}
