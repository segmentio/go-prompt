package prompt

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
)

// String prompt.
func String(prompt string, args ...interface{}) string {
	var s string
	fmt.Printf(prompt+": ", args...)
	fmt.Scanln(&s)
	return s
}

// String prompt (required).
func StringRequired(prompt string, args ...interface{}) (s string) {
	for strings.Trim(s, " ") == "" {
		s = String(prompt, args...)
	}
	return s
}

// String default (required).
func StringDefault(prompt string, dvalue string) (s string) {
	fmt.Println()
	fmt.Printf("%s (default value: %s)", prompt, dvalue)
	fmt.Scanln(&s)
	if strings.Trim(s, " ") == "" {
		return dvalue
	}
	return s
}

// Stringln reads multi spaced sentence until newline
func Stringln(scanner *bufio.Scanner, prompt string) string {
	fmt.Printf(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}

// Integer prompt (required).
func IntegerRequired(prompt string, args ...interface{}) int {
	for {
		n := Integer(prompt, args...)
		if n != 0 {
			return n
		}
	}
}

// Integer prompt
func Integer(prompt string, args ...interface{}) int {
	s := String(prompt, args...)
	n, _ := strconv.Atoi(s)
	return n
}

// Confirm continues prompting until the input is boolean-ish.
func Confirm(prompt string, args ...interface{}) bool {
	for {
		switch String(prompt, args...) {
		case "Yes", "yes", "y", "Y":
			return true
		case "No", "no", "n", "N":
			return false
		}
	}
}

// Choose prompts for a single selection from `list`, returning in the index.
func Choose(prompt string, list []string) int {
	fmt.Println()
	for i, val := range list {
		fmt.Printf("  %d) %s\n", i+1, val)
	}

	fmt.Println()
	i := -1

	for {
		s := String(prompt)

		// index
		n, err := strconv.Atoi(s)
		if err == nil {
			if n > 0 && n <= len(list) {
				i = n - 1
				break
			} else {
				continue
			}
		}

		// value
		i = indexOf(s, list)
		if i != -1 {
			break
		}
	}

	return i
}

// Choose prompts for a single selection from `list`, returning in the index.
func ChooseInterface(prompt string, list []interface{}) int {
	fmt.Println()
	for i, val := range list {
		fmt.Printf("  %d) %+v\n", i+1, val)
	}

	fmt.Println()
	i := -1

	for {
		s := String(prompt)

		// index
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		if n > 0 && n <= len(list) {
			i = n - 1
			break
		} else {
			continue
		}

	}

	return i
}

// Password prompt.
func Password(prompt string, args ...interface{}) string {
	fmt.Printf(prompt+": ", args...)
	password, _ := gopass.GetPasswd()
	s := string(password[0:])
	return s
}

// Password prompt with mask.
func PasswordMasked(prompt string, args ...interface{}) string {
	fmt.Printf(prompt+": ", args...)
	password, _ := gopass.GetPasswdMasked()
	s := string(password[0:])
	return s
}

// index of `s` in `list`.
func indexOf(s string, list []string) int {
	for i, val := range list {
		if val == s {
			return i
		}
	}
	return -1
}
