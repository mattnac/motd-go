package main

import (
	"bufio"
	"flag"
	"fmt"
	"motd/message"
	"os"
	"strings"
)

func main() {
	// Define flag variables
	var name string
	var greeting string
	var prompt bool
	var preview bool

	// Parse flags
	flag.StringVar(&name, "name", "", "name to use within the message")
	flag.StringVar(&greeting, "greeting", "", "Phrase to use in the message")
	flag.BoolVar(&prompt, "prompt", false, "Use prompt to input name and greeting")
	flag.BoolVar(&preview, "preview", false, "Use preview to output message without writing.")
	flag.Parse()

	// Show usage if flags are invalid
	if prompt == false && (name == "" || greeting == "") {
		flag.Usage()
		os.Exit(1)
	}

	// Optionally print flags if DEBUG env var is set.
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name:", name)
		fmt.Println("Greeting:", greeting)
		fmt.Println("Prompt:", prompt)
		fmt.Println("Preview:", preview)
		os.Exit(0)
	}

	// Conditionally read from stdin.
	if prompt {
		name, greeting = renderPrompt()
	}

	// Generate message
	message := message.Greeting(name, greeting)

	// Either preview the message or write to file
	if preview {
		fmt.Println(message)
	} else {
		// Write content
		f, err := os.OpenFile("/tmp/motd", os.O_WRONLY, 0644)

		if err != nil {
			fmt.Println("Unable to open file.")
			os.Exit(1)
		}

		defer f.Close()

		_, err = f.Write([]byte(message))

		if err != nil {
			fmt.Println("Unable to write to file")
			os.Exit(1)
		}
	}
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	return
}
