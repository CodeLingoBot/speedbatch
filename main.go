package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/juju/errors"
)

func main() {
	lines, err := getLines(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	res, err := launchTabs(lines)
	if err != nil {
		log.Fatal(err)
	}
	fileData := strings.Join(res, "\n")
	ioutil.WriteFile("PRs.txt", []byte(fileData), 0600)
}

func getLines(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Trace(err)
	}
	splitData := strings.Split(string(data), "\n")

	return splitData, nil
}

func launchTabs(urls []string) ([]string, error) {
	var prurls []string
	scanner := bufio.NewScanner(os.Stdin)
	for _, url := range urls {
		if url == "" {
			continue
		}
		cmd := exec.Command("chromium-browser", url)
		cmd.Run()
		scanner.Scan()
		prurls = append(prurls, scanner.Text())
	}
	var filtered []string
	for _, url := range prurls {
		match, err := regexp.MatchString("https://github.com.*", url)
		if err != nil {
			return nil, errors.Trace(err)
		}
		if match {
			filtered = append(filtered, url)
		}
	}
	return filtered, nil
}
