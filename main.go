package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Println("Parse Form Successully")
	ans := longestPalindrome(r.FormValue("YourString"))
	fmt.Fprintf(w, "name: %v\n", ans)
	fmt.Fprintf(w, "length: %v\n", len(ans))
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "hello!\n")
}

func longestPalindrome(s string) string {
	l := len(s)
	fmt.Println(l)
	p := make([][]bool, l)
	for i := 0; i < l; i++ {
		p[i] = make([]bool, l)
	}
	a := 0
	b := 0
	for j := 0; j < l; j++ {
		p[j][j] = true
		for i := 0; i < j; i++ {
			if j == i+1 && s[i] == s[j] {
				p[i][j] = true
				fmt.Println(p[i][j], i, j)
			} else if j > i+1 {
				p[i][j] = s[i] == s[j] && p[i+1][j-1]
				fmt.Println(p[i][j], i, j)
			}
		}
	}
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if p[i][j] && j-i > b-a {
				a = i
				b = j
				fmt.Println(a, b)
			}
		}
	}

	return s[a : b+1]
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/form", formhandler)
	fmt.Println("Starting Service at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
