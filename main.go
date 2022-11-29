package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func compress(bitmap string) string {
	if len(bitmap) != 25 {
		return "Invalid Input!"
	}
	var sb strings.Builder
	for i := 0; i < 25; i += 5 {
		ui, err := strconv.ParseUint(bitmap[i:i+5], 2, 64)
		if err != nil {
			return "error"
		}
		sb.WriteString(fmt.Sprintf("%02x", ui))
	}
	return strings.ToUpper(sb.String())
}

func decompress(hex string) string {
	if len(hex) != 10 {
		return "Invalid Input!"
	}
	var sb strings.Builder
	for i := 0; i < 10; i += 2 {
		ui, err := strconv.ParseUint(hex[i:i+2], 16, 64)
		if err != nil {
			return "error"
		}
		sb.WriteString(fmt.Sprintf("%05b", ui))
		sb.WriteString("\r\n")
	}
	return strings.ToUpper(sb.String())
}

func encode(w http.ResponseWriter, req *http.Request) {
	bitmap := strings.TrimPrefix(req.URL.Path, "/encode/")
	hex := compress(bitmap)
	_, err := fmt.Fprintf(w, hex)
	if err != nil {
		return
	}
}

func decode(w http.ResponseWriter, req *http.Request) {
	hex := strings.TrimPrefix(req.URL.Path, "/decode/")
	bitmap := decompress(hex)
	_, err := fmt.Fprintf(w, bitmap)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/encode/", encode)
	http.HandleFunc("/decode/", decode)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running ...")
	if err != nil {
		fmt.Println(err)
		return
	}
}
