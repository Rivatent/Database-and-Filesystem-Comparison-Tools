package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	usageMessage  = "usage: ./compareFS [-old] oldSnapshotFilepath [-new] newSnapshotFilepath"
	filenameError = "filename is required"
)

func main() {
	oldSnapshotFilepath, newSnapshotFilepath, error := CheckInputCompareFS()
	if error != nil {
		log.Fatal(error)
	}

	CompareSnapshots(os.Stdout, oldSnapshotFilepath, newSnapshotFilepath)
}

func CheckInputCompareFS() (oldSnapshotFilepath, newSnapshotFilepath string, err error) {
	oldSnapshotFlag := flag.String("old", "", usageMessage)
	newSnapshotFlag := flag.String("new", "", usageMessage)
	flag.Parse()
	if flag.NArg() != 0 {
		err = errors.New(usageMessage)
		return
	}
	oldSnapshotFilepath = *oldSnapshotFlag
	newSnapshotFilepath = *newSnapshotFlag

	if oldSnapshotFilepath == "" || newSnapshotFilepath == "" {
		err = errors.New(filenameError)
		return
	}

	return
}

func CompareSnapshots(out io.Writer, oldSnapshotFilepath, newSnapshotFilepath string) {

	oldSnapshotFile, error := os.Open(oldSnapshotFilepath)
	if error != nil {
		log.Fatal(error)
	}
	defer oldSnapshotFile.Close()

	oldSnapshotLines := map[string]bool{}

	oldScanner := bufio.NewScanner(oldSnapshotFile)
	for oldScanner.Scan() {
		oldSnapshotLines[oldScanner.Text()] = true
	}

	if err := oldScanner.Err(); err != nil {
		log.Fatalf("Error reading old snapshot: %v", err)
	}

	newSnapshotFile, err := os.Open(newSnapshotFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer newSnapshotFile.Close()

	newScanner := bufio.NewScanner(newSnapshotFile)
	for newScanner.Scan() {
		if !oldSnapshotLines[newScanner.Text()] {
			fmt.Fprintf(out, "ADDED %s\n", newScanner.Text())
		} else {
			delete(oldSnapshotLines, newScanner.Text())
		}
	}
	if err := newScanner.Err(); err != nil {
		log.Fatalf("Error reading new snapshot: %v", err)
	}

	for k := range oldSnapshotLines {
		fmt.Fprintf(out, "REMOVED %s\n", k)
	}
}
