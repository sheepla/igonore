package main

import (
	"flag"
	"fmt"
	"os"
)

const ignoreName = ".gitignore"

func main() {
	flag.Parse()
	args := flag.Args()

	if inStr(args, "help") {
		showHelp()
		os.Exit(0)
	}

	if len(args) == 0 {
		if err := intaractive(); err != nil {
			fmt.Fscanln(os.Stderr, err)
		}
	} else {
		ignore, err := fetch(args)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err := save(ignoreName, ignore); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	fmt.Println("Done! 🎉")
}

func intaractive() error {
	var selectLangs []string

	for {
		err := Prompt("Do you want to add the specified language")
		if err != nil {
			break
		} else {
			fmt.Println("Open the viewfinder...")
		}

		lang, err := finder()
		if err != nil {
			return err
		}

		selectLangs = append(selectLangs, lang)
	}

	ignore, err := fetch(selectLangs)
	if err != nil {
		return err
	}

	if err := save(ignoreName, ignore); err != nil {
		return fmt.Errorf("failed to save .gitignore: %w", err)
	}

	return nil
}

func save(fileName string, ctn string) (err error) {
	_, err = os.Stat(fileName)
	if err == nil {
		err := Prompt("Overwrite file. Are you sure")
		if err != nil {
			return fmt.Errorf("an error occurred on prompt: %w", err)
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create the file(%s): %w: err", fileName, err)
	}

	defer func() {
		if e := file.Close(); e != nil {
			err = fmt.Errorf("failed to close the file(%s): %w", fileName, err)
		}
	}()

	_, err = file.Write([]byte(ctn))
	if err != nil {
		return fmt.Errorf("failed to write the content to file(%s): %w", fileName, err)
	}

	return nil
}

func showHelp() {
	const help = `
  igonore - .gitignore generator written in Go

  Usage: igonore
         Generate .gitignore interactively.

         igonore langage langage2...
         Generate a .gitignore file for the specified language.

         Ex.) igonore go node
  `
	fmt.Println(help)
}

func inStr(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}

	return false
}
