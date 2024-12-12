package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	//args := os.Args

	//filePath := args[1]
	//file, err := os.Open(filePath)
	file, err := os.Open("./tests/step1/valid.json")
	if err != nil {
		fmt.Println("cannot find file")
		return
	}
	defer file.Close()

	queue := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] != '{' {
			fmt.Println("invalid file 1")
			return
		}

		shouldProcessValue := false
		for _, c := range line {
			if unicode.IsSpace(c) {
				continue
			}

			if shouldProcessValue {
				// TODO: handle types
				if queue[len(queue)-1] == ":" {
					queue = append(queue, "value")
				}
			}

			current := string(c)
			switch current {

			case "{":
				// push "{"
				queue = append(queue, current)

			case "}":
				if checkEmptyQueue(queue) || isValidClosing(&queue) == false {
					fmt.Println("invalid format")
					return
				}
			case "\"":
				if checkEmptyQueue(queue) {
					fmt.Println("invalid format")
					return
				}

				if queue[len(queue)-1] != "\"" {
					// push (")
					queue = append(queue, current)
				} else {
					// pop last element (")
					queue = queue[:len(queue)-1]
					queue = append(queue, "key")
				}
			case ":":
				if checkEmptyQueue(queue) {
					return
				}

				if queue[len(queue)-1] != "key" {
					fmt.Println("invalid format")
					return
				}

				shouldProcessValue = true
				queue = append(queue, ":")

			case ",":
				if checkEmptyQueue(queue) {
					return
				}

				if queue[len(queue)-1] != "value" {
					fmt.Println("invalid format")
					return
				}

				shouldProcessValue = false
			}
		}
	}

	if len(queue) > 0 {
		fmt.Println("invalid file 2")
		return
	}
	fmt.Println("VALID")
}

func isValidClosing(queue *[]string) bool {
	cnt := 0
	if len(*queue) > 0 && (*queue)[len(*queue)-1] == "value" {
		*queue = (*queue)[:len(*queue)-1]
		cnt++
	}
	if len(*queue) > 0 && (*queue)[len(*queue)-1] == ":" {
		*queue = (*queue)[:len(*queue)-1]
		cnt++
	}
	if len(*queue) > 0 && (*queue)[len(*queue)-1] == "key" {
		*queue = (*queue)[:len(*queue)-1]
		cnt++
	}
	if len(*queue) > 0 && (*queue)[len(*queue)-1] == "{" {
		*queue = (*queue)[:len(*queue)-1]
		cnt++
	}
	return cnt == 4
}

func checkEmptyQueue(queue []string) bool {
	return len(queue) == 0
}
