package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessCommand(command string) ([]string, int, error) {
	command = strings.TrimRight(command, "\r\n")
	args := strings.Split(command, " ")
	length := len(args)

	switch {
	case length == 1:
		return args, -1, nil

	case length == 2:
		if args[1] == "done" || args[1] == "in-progress" || args[1] == "todo" {
			return args, -1, nil
		}

		if id, isDigit := strconv.ParseInt(args[1], 10, 32); isDigit == nil {
			return args[:1], int(id), nil
		}
		description, err := parseDescription(args[1:])

		if err != nil {
			return nil, -1, err
		}
		return append(args[:1], description), -1, nil

	case length >= 3:
		if id, isDigit := strconv.ParseInt(args[1], 10, 32); isDigit != nil {
			description, err := parseDescription(args[1:])
			if err != nil {
				return nil, -1, err
			}
			return append(args[:1], description), -1, nil
		} else {
			description, err := parseDescription(args[2:])
			if err != nil {
				return nil, -1, err
			}
			return append(args[:1], description), int(id), nil
		}

	default:
		return nil, -1, nil
	}
}

func parseDescription(description []string) (string, error) {
	indexOpen, indexClose := 0, 0
	counter := 0
	for i := range description {
		if strings.HasPrefix(description[i], "\"") {
			indexOpen = i
			break
		}
		counter++
	}

	if counter != indexOpen {
		return "", fmt.Errorf("first quotation mark in missing")
	}
	tmp := description[indexOpen:]

	counter = len(tmp) - 1

	for i := counter; i >= 0; i-- {
		if strings.HasSuffix(tmp[i], "\"") {
			indexClose = i
			break
		}
		counter--
	}

	if counter != indexClose {
		return "", fmt.Errorf("second quotation mark in missing")
	}
	res := strings.Join(tmp[:indexClose+1], " ")
	res = strings.Trim(res, "\"")
	return res, nil
}
