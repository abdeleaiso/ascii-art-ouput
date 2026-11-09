package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Read(bannername string) []string {
	fileData, err := os.ReadFile(bannername)
	if err != nil {
		fmt.Println("Error: ", err)
		return []string{}
	}
	// in windows they have \r\n intead of \n
	fileText := strings.ReplaceAll(string(fileData), "\r\n", "\n")

	// split string with \n into slice of strings
	lines := strings.Split(fileText, "\n")

	return lines
}

func validateArgs(arg string) error {
	if len(arg) < 10 {
		return errors.New("invalide format")
	}
	if len(arg) > 9 && arg[:9] != "--output=" {
		return errors.New("invalide format")
	}
	if strings.HasPrefix(arg, "--output=") {
		filename := strings.TrimPrefix(arg, "--output=")

		if filename == "" || !strings.HasSuffix(filename, ".txt") {
			return errors.New("")
		}
	}
	return nil
}

func Error() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}

func main() {
	if len(os.Args) == 1 {
		Error()
		return
	}
	output := flag.String("output", "", "output in different format")
	err := validateArgs(os.Args[1])
	if err != nil {
		Error()
		return
	}
	flag.Parse()
	fprotected := []string{"standard.txt", "shadow.txt", "thinkertoy.txt", "go.mod", "main.go", "RAEDME.md", "tranformToAscii"}
	for _, v := range fprotected {
		if *output == v {
			fmt.Println("Error: please select other file than the program files!!")
			return
		}
	}

	args := flag.Args()
	if len(args) == 0 {
		Error()
		return
	}
	input := args[0]
	if input == "" {
		return
	}

	if len(args) == 1 {
		if *output == "" {
			lines := Read("standard.txt")
			res := TransformToAscii(input, lines)
			fmt.Print(res)
		} else {
			lines := Read("standard.txt")
			content := []byte(TransformToAscii(input, lines))

			err := os.WriteFile(*output, content, 0644)
			if err != nil {
				fmt.Println("Error Writing in file:", err)
				return
			}
		}
	} else if len(args) == 2 {
		bannername := args[1] + ".txt"
		// read banner file he return []byte
		lines := Read(bannername)

		content := []byte(TransformToAscii(input, lines))

		err := os.WriteFile(*output, content, 0644)
		if err != nil {
			fmt.Println("Error Writing in file:", err)
			return
		}
	} else {
		Error()
		return
	}
}
