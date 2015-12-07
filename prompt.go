package prompt

import "github.com/howeyc/gopass"
import "strings"
import "strconv"
import "fmt"
import "bufio"
import "os"

// String prompt.
func String(prompt string, args ...interface{}) string {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(prompt+": ", args...)
	s, _ = reader.ReadString('\n')
	return s
}

// String prompt (required).
func StringRequired(prompt string, args ...interface{}) string {
	s := String(prompt, args...)
	if strings.Trim(s, " ") == "" {
		return StringRequired(prompt)
	} else {
		return s
	}
}

// Confirm continues prompting until the input is boolean-ish.
func Confirm(prompt string, args ...interface{}) bool {
	s := String(prompt, args...)
	switch s {
	case "yes", "y", "Y":
		return true
	case "no", "n", "N":
		return false
	default:
		return Confirm(prompt, args...)
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

// Password prompt.
func Password(prompt string, args ...interface{}) string {
	fmt.Printf(prompt+": ", args...)
	s := string(gopass.GetPasswd()[0:])
	return s
}

// Password prompt with mask.
func PasswordMasked(prompt string, args ...interface{}) string {
	fmt.Printf(prompt+": ", args...)
	s := string(gopass.GetPasswdMasked()[0:])
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
