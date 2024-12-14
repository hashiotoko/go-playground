package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // エラー詳細を割愛
	// log.SetFlags(log.Ldate + log.Llongfile) // エラー詳細表示内容をbit値の加算で設定できる

	names := []string{
		"Mike",
		"Taro",
		"Samantha",
	}

	messages, err := greetings.Hellos(names)

	if (err != nil) {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
