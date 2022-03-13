package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	fileOne := flag.String("file1", "c1.txt", "File containing first set of data")
	fileTwo := flag.String("file2", "c2.txt", "File containing second set of data")
	fileOutput := flag.String("output", "output.txt", "File to save the XOR")
	flag.Parse()

	fmt.Sprintf("Reading %s and %s to output the XOR of their content", fileOne, fileTwo)

	fileOneHandle, err := os.Open(*fileOne)
	check(err)
	defer fileOneHandle.Close()
	fileTwoHandle, err := os.Open(*fileTwo)
	check(err)
	defer fileTwoHandle.Close()
	fileOutputHandle, err := os.Create(*fileOutput)
	check(err)
	defer fileOutputHandle.Close()

	var intOne uint64
	var intTwo uint64
	var output uint64
	for {
		err = binary.Read(fileOneHandle, binary.LittleEndian, &intOne)
		// fmt.Println(reflect.TypeOf(err))
		if err == io.EOF {
			break
		} else {
			check(err)
		}
		err = binary.Read(fileTwoHandle, binary.LittleEndian, &intTwo)

		output = intOne ^ intTwo

		binary.Write(fileOutputHandle, binary.LittleEndian, &output)
	}

}
