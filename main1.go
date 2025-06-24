package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func ProcessLogs(inputFiles []string, outputFile string) error {
	errorChan := make(chan string)
	var wg sync.WaitGroup
	writerDone := make(chan struct{})
	go func() {
		outFile, err := os.Create(outputFile)
		if err != nil {
			log.Printf("Failed to create output file: %v", err)
			close(writerDone)
			return
		}
		defer outFile.Close()

		writer := bufio.NewWriter(outFile)
		for line := range errorChan {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				log.Printf("Failed to write to output: %v", err)
			}
		}
		writer.Flush()
		close(writerDone)
	}()
	for _, file := range inputFiles {
		wg.Add(1)
		go func(filename string) {
			defer wg.Done()

			f, err := os.Open(filename)
			if err != nil {
				log.Printf("Error opening file %s: %v", filename, err)
				return
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "ERROR") {
					errorChan <- line
				}
			}

			if err := scanner.Err(); err != nil {
				log.Printf("Error reading file %s: %v", filename, err)
			}
		}(file)
	}
	go func() {
		wg.Wait()
		close(errorChan)
	}()
	<-writerDone

	return nil
}
func createTestLogFiles() {
	files := map[string][]string{
		"server1.log": {
			"INFO: Server started successfully",
			"ERROR: Failed to connect to database",
			"INFO: Request processed",
			"ERROR: Timeout occurred while processing request",
		},
		"server2.log": {
			"INFO: Disk check",
			"ERROR: Disk space low",
			"INFO: Backup completed",
		},
		"server3.log": {
			"INFO: Init service",
			"ERROR: Unable to authenticate user",
			"INFO: Running background task",
		},
	}

	for name, lines := range files {
		f, err := os.Create(name)
		if err != nil {
			log.Fatalf("Failed to create %s: %v", name, err)
		}
		for _, line := range lines {
			fmt.Fprintln(f, line)
		}
		f.Close()
	}
}

func main() {
	createTestLogFiles()
	inputFiles := []string{"server1.log", "server2.log", "server3.log"}
	outputFile := "errors.log"
	err := ProcessLogs(inputFiles, outputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Log processing completed. See 'errors.log' for output.")
}
