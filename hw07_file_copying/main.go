package main

import (
	"flag"
	"log"
	"os"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	logger := log.New(os.Stderr, "", 0)

	if len(from) == 0 {
		logger.Println("Provide '-from=<...>' argument")
		return
	}
	if len(to) == 0 {
		logger.Println("Provide '-to=<...>' argument")
		return
	}

	err := Copy(from, to, offset, limit)
	if err != nil {
		logger.Println(err)
	}
}
