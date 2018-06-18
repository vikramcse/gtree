package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := tree(arg, "")
		if err != nil {
			log.Printf("tree %s: %v", arg, err)
		}
	}
}

func tree(root, indent string) error {
	fi, err := os.Stat(root)
	if err != nil {
		return fmt.Errorf("could not able to start %s: %v", root, err)
	}

	fmt.Println(fi.Name())

	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return fmt.Errorf("could not able to read the files in dir %s: %v", root, err)
	}

	var names []string
	for _, info := range fis {
		if info.Name()[0] != '.' {
			names = append(names, info.Name())
		}
	}

	for i, name := range names {
		add := "│ "
		if i == len(names)-1 {
			fmt.Printf(indent + "└──")
			add = "  "
		} else {
			fmt.Printf(indent + "├──")
		}

		if err := tree(filepath.Join(root, name), indent+add); err != nil {
			return err
		}
	}
	return nil
}
