package prompt

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
)

// String prompt.
func String(msg string, args ...interface{}) string {
	var s string
	prompt(msg, args...)
	fmt.Scanln(&s)
	return s
}

// String prompt (required).
func StringRequired(msg string, args ...interface{}) (s string) {
	for strings.Trim(s, " ") == "" {
		s = String(msg, args...)
	}
	return s
}

// String default (required).
func StringDefault(msg string, dvalue string) (s string) {
	fmt.Printf("%s (default value: %s)", msg, dvalue)
	fmt.Scanln(&s)
	if strings.Trim(s, " ") == "" {
		return dvalue
	}
	return s
}

// Stringln reads multi spaced sentence until newline
func Stringln(scanner *bufio.Scanner, msg string) string {
	prompt(msg)
	scanner.Scan()
	return scanner.Text()
}

// Integer prompt (required).
func IntegerRequired(msg string, args ...interface{}) int {
	for {
		n := Integer(msg, args...)
		if n != 0 {
			return n
		}
	}
}

// Integer prompt
func Integer(msg string, args ...interface{}) int {
	s := String(msg, args...)
	n, _ := strconv.Atoi(s)
	return n
}

// Confirm continues prompting until the input is boolean-ish.
func Confirm(msg string, args ...interface{}) bool {
	for {
		switch String(msg, args...) {
		case "Yes", "yes", "y", "Y":
			return true
		case "No", "no", "n", "N":
			return false
		}
	}
}

// Choose prompts for a single selection from `list`, returning in the index.
func Choose(msg string, list []string) int {
	fmt.Println()
	for i, val := range list {
		fmt.Printf("  %d) %s\n", i+1, val)
	}

	fmt.Println()
	i := -1

	for {
		s := String(msg)

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
func ChooseInterface(msg string, list []interface{}) int {
	fmt.Println()
	for i, val := range list {
		fmt.Printf("  %d) %+v\n", i+1, val)
	}

	fmt.Println()
	i := -1

	for {
		s := String(msg)

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
func Password(msg string, args ...interface{}) string {
	prompt(msg, args...)
	password, _ := gopass.GetPasswd()
	s := string(password[0:])
	return s
}

// Password prompt with mask.
func PasswordMasked(msg string, args ...interface{}) string {
	prompt(msg, args...)
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

func prompt(msg string, args ...interface{}) {
	fmt.Printf(msg+": ", args...)
}
