package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/cheggaaa/pb"
)

func readInput(fileName string, BlockSize int, stream *chan []byte, progressStream *chan int64, wg *sync.WaitGroup) {
	// Defer waitgroup go-routine done before returning
	defer wg.Done()

	// Open input file
	file, err := os.Open(fileName)
	if err != nil {
		panic(err.Error())
	}

	// Defer file close and panic if there is an error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	// Get file info
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err.Error())
	}

	// Get input file size
	fileSize := fileInfo.Size()

	// Push input file size as the first input to progress stream
	*progressStream <- fileSize

	reader := bufio.NewReader(file)

	// Initialize offset
	offset := int64(0)
	i64BlockSize := int64(BlockSize)
	// Loop until the entire file is read
	for {
		// Determine the size of the block to read
		readSize := BlockSize
		if fileSize-offset < i64BlockSize {
			readSize = int(fileSize - offset)
		}

		// Create a block with the determined size
		block := make([]byte, readSize)

		// Read from the buffered reader to the block
		// Panic if there are any errors
		bytesRead, err := reader.Read(block)
		if err != nil {
			panic(err.Error())
		}

		// PUSH block to buffered stream channel
		*stream <- block

		// Increment offset by bytesRead
		offset += int64(bytesRead)

		// Push offset to buffered progress channel
		*progressStream <- offset

		// Check if the entire file is read, exit if so
		if fileSize-offset == 0 {
			// Send nil to stream to signal end of input
			*stream <- nil
			break
		}
	}
}

func writeOutput(fileName string, stream *chan []byte, wg *sync.WaitGroup) {
	// Defer waitgroup go-routine done before returning
	defer wg.Done()

	// Create output file
	file, err := os.Create(fileName)
	if err != nil {
		panic(err.Error())
	}

	// Defer file close and panic if there is an error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err.Error())
		}
	}()

	// Create a buffered writer
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Initialize offset
	offset := int64(0)
	for {
		// Read block from stream
		block := <-*stream

		// A nil block signals end, break out of loop
		if block == nil {
			// I WANT TO BREAK FREE
			// FROM FOR LOOP
			break
		}

		// Write block to file at offset
		// Panic if there are any errors
		bytesWritten, err := writer.Write(block)
		if err != nil {
			panic(err.Error())
		}

		// Increment offset by bytesWritten
		offset += int64(bytesWritten)
	}
}

func progressBar(fileSize int64, progressStream *chan int64) {
	// Create new progressbar with count
	bar := pb.Start64(fileSize)
	// Set template to full so we get remaining time
	bar.SetTemplate(pb.Full)
	// Set bytes to true so we get nicely formatted output
	bar.Set(pb.Bytes, true)
	var offset int64
	for offset < fileSize {
		offset = <-*progressStream
		// Set bar progress to current offset from the reader
		bar.SetCurrent(offset)
	}
	bar.Finish()
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readFromFile(filePath string) []byte {
	// Check if file exists and if not, print
	if fileExists(filePath) {
		data, err := os.ReadFile(filePath)
		if err != nil {
			panic(err.Error())
		}
		return data
	}
	// If file doesn't exist, print so and exit
	fmt.Println("File:", filePath, "seems to be nonexistent")
	os.Exit(1)
	// Fun fact: if you remove the below return
	// Compile will fail - as functions which return must
	// return at the end, BUT... The above OS Exit makes it
	// such that it'll never execute XD
	return nil
}

func printHelp() {
	// Ah, yes; help.
	fmt.Printf("%s is a CLI program which implements the SeaTurtle Block Cipher (http://github.com/sid-sun/seaturtle) in CFB (cipher feedback) mode with 256-Bit key length, using SHA3-256.", os.Args[0])
	fmt.Printf("\nDeveloped by Sidharth Soni (Sid Sun) <sid@sidsun.com>")
	fmt.Printf("\nOpen-sourced under The Unlicense")
	fmt.Printf("\nSource Code: http://github.com/sid-sun/seat-256-cfb\n")
	fmt.Printf("\nUsage:\n")
	fmt.Printf("    To encrypt: %s (--encrypt / -e) <input file> <passphrase file> <output file (optional)>\n", os.Args[0])
	fmt.Printf("    To decrypt: %s (--decrypt / -d) <encrypted input> <passphrase file> <output file (optional)>\n", os.Args[0])
	fmt.Printf("    To get version number: %s (--version / -v)\n", os.Args[0])
	fmt.Printf("    To get help: %s (--help / -h)\n", os.Args[0])
}
