package main

import (
	"flag"
	"log"
	"os"

	"github.com/JokeTrue/otus-golang/hw09_generator_of_validators/go-validate/core"
)

var modelFilePath string

var errorLogger = log.New(os.Stderr, "", 0)

func init() {
	flag.StringVar(&modelFilePath, "path", "", "Path to file with model")
	flag.Parse()
}

func main() {
	stat, err := os.Stat(modelFilePath)
	if err != nil {
		errorLogger.Fatalln(err)
	}
	if !stat.Mode().IsRegular() {
		errorLogger.Fatalln("not a regular file")
	}

	pkgName, validatedStructs, err := core.Parse(modelFilePath)
	if err != nil {
		errorLogger.Fatalln(err)
	}

	content, err := core.Generate(pkgName, validatedStructs)
	if err != nil {
		errorLogger.Fatalln(err)
	}

	err = core.WriteTemplateCode(modelFilePath, content)
	if err != nil {
		errorLogger.Fatalln(err)
	}
}
