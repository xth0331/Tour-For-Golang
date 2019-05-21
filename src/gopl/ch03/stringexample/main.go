// "program" in Japanese katakana
package main

import "fmt"

func main() {
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string(r))
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x9E1F)) // "京"
}
