package main

import (
  "os"
  "fmt"
  "log"
  "strconv"
  "crypto/rand"
  "encoding/hex"
  "path/filepath"
)

func main() {
  filename, length := InputCheck()
  if length == -1 {
    length = 5
  }
  rbyte := make([]byte, length)
  _, err := rand.Read(rbyte)
  CheckErr(err)
  outname := hex.EncodeToString(rbyte)
  err = os.Rename(filename, outname + filepath.Ext(filename))
  CheckErr(err)
}

var banner string = `
Usage : mvrand [file name] [options]

Options :
-h help
-l length of the random file name (default: 10)
`


func InputCheck() (filename string, length int){
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
