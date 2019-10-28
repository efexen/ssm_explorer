package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"
)

func main() {
	stdin, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if stdin.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("SSM Explorer")
		fmt.Println("Usage: aws ssm describe-parameters | ssm_explorer")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	params, err := ioutil.ReadAll(reader)

	if err != nil {
		fmt.Println("Error reading input")
		log.Fatal(err)
	}

	layout, err := template.ParseFiles("index.html")

	var out bytes.Buffer

	data := struct {
		Params string
	}{
		Params: string(params),
	}

	if err := layout.Execute(&out, data); err != nil {
		fmt.Println("Error rendering template")
		log.Fatal(err)
	}

	file, err := ioutil.TempFile("", "ssm_explorer*.html")

	if err != nil {
		fmt.Println("Error creating temporary HTML file")
		log.Fatal(err)
	}

	defer os.Remove(file.Name())

	_, err = out.WriteTo(file)
	if err != nil {
		fmt.Println("Error writing temp file")
		log.Fatal(err)
	}

	file.Sync()

	err = exec.Command("open", file.Name()).Start()

	if err != nil {
		fmt.Println("Error opening file in browser")
		log.Fatal(err)
	}

	fmt.Printf(file.Name())

	time.Sleep(1000 * time.Millisecond) // wait for browser open before we destroy temp file
}
