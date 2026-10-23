package main

import "os"

func main() {
	args := os.Args

	if len(args) < 2 {
		os.Stdout.WriteString("File name missing\n")
		return
	}
	if len(args) > 2 {
		os.Stdout.WriteString("Too many arguments\n")
		return
	}

	file, err := os.Open(args[1])
	if err != nil {
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			os.Stdout.Write(buffer[:n])
		}
		if err != nil {
			break
		}
	}
}
