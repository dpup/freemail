// +build ignore

// This program generates free.go, disposable.go, and spammy.go
// It can be invoked by running go generate
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	writeFile("free", readFile("free.txt"))
	writeFile("disposable", readFile("disposable.txt"))
	writeFile("spammy", readFile("spammy.txt"))
}

func readFile(name string) []string {
	b, err := ioutil.ReadFile("./data/" + name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(b), "\n")
}

func writeFile(name string, domains []string) {
	fmt.Printf("Writing %d domains to %s.go\n", len(domains), name)
	lines := []string{
		"package freemail",
		"",
		"var " + name + " = map[string]bool{",
	}
	for _, domain := range domains {
		if domain != "" {
			lines = append(lines, "	\""+domain+"\": true,")
		}
	}
	lines = append(lines, "}")
	content := strings.Join(lines, "\n")
	ioutil.WriteFile(name+".go", []byte(content), 0644)
}
