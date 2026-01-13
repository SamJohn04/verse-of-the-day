package internal

import "fmt"

func Bold(s string) string {
	return fmt.Sprintf("\033[1m%v\033[0m", s)
}
