package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	f := flag.String("f", "", "Input File Name")
	flag.Parse()
	lines := readLines(*f)
	solve(lines)
}

func solve(lines []string) {
	input := lines[0]
	findNum(input, "00000")  // part 1
	findNum(input, "000000") // part 2
}

func findNum(input, prefix string) {
	for i := 0; ; i++ {
		hash := md5.New()
		hash.Write([]byte(input + strconv.Itoa(i)))
		hashStr := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(hashStr, prefix) {
			fmt.Println(hashStr, i)
			break
		}
	}
}
