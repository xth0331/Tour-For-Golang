// basename removes direcroty components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b/c.go => b.c
package main

import "fmt"

func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything vefore last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	fmt.Println(basename("a.txt"))
}
