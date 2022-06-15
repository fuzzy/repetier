package main

import "fmt"

var (
	RED = 1
	GREEN = 2
	YELLOW = 3
	BOLD = 1
	END = 0
)

func red(s string) string {
	return fmt.Sprintf("\033[%d;3%dm%s\033[%dm", BOLD, RED, s, END)
}

func green(s string) string {
	return fmt.Sprintf("\033[%d;3%dm%s\033[%dm", BOLD, GREEN, s, END)
}

func yellow(s string) string {
	return fmt.Sprintf("\033[%d;3%dm%s\033[%dm", BOLD, YELLOW, s, END)
}
