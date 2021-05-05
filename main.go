package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	filename, length := InputCheck()
	if length == -1 {
		length = 5
	}
	if filename == "." {
		RenameAll(length)
		return
	}
	outname := RandName(length)
	ex := filepath.Ext(filename)
	fmt.Println(filename + " =>" + outname + ex)
	err := os.Rename(filename, outname+ex)
	CheckErr(err)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////

func RenameAll(length int) {
	files, err := ioutil.ReadDir(".")
	CheckErr(err)
	for _, f := range files {
		outname := RandName(length)
		ex := filepath.Ext(f.Name())
		fmt.Println(f.Name() + "=> " + outname + ex)
		err = os.Rename(f.Name(), outname+ex)
	}
}

func RandName(length int) string {
	rbyte := make([]byte, length)
	_, err := rand.Read(rbyte)
	CheckErr(err)
	return hex.EncodeToString(rbyte)
}

func InputCheck() (filename string, length int) {
	if len(os.Args) < 2 || os.Args[1] == "-h" {
		fmt.Println(banner)
		os.Exit(0)
	}
	filename = os.Args[1]
	if len(os.Args) > 2 && os.Args[2] == "-l" {
		num := os.Args[3]
		length, err := strconv.Atoi(num)
		CheckErr(err)
		length = length / 2
		return filename, length
	}
	return filename, -1
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
var banner string = `
Usage : mvrand [file name] [options]

Options :
-h help
-l length of the random file name (default: 10)
`
