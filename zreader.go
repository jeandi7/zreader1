package main

import (
	"flag"
	"fmt"
	"os"
	"zreader1/zinterpreter"
)

func printHelp() {
	fmt.Println("2024 : See my blog https://jeandi7.github.io/jeandi7blog/")
	fmt.Println()
	fmt.Println("Usage: zanzireader [options]")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {

	// examples :
	// input := `definition monsujet { }`
	// input := `definition monsujet { relation marelation1: monsujet2 }`
	// input := `definition monsujet { relation marelation1: monsujet2 | monsujet3 }`
	// input := `definition monsujet { relation marelation1: monsujet2 | monsujet3  relation marelation2: monsujet2 }`
	// input := `definition monsujet { } definition monsujet2 { } definition maressource { relation marelation: monsujet | monsujet2   }`
	// input := `definition monsujet { } definition monsujet2 { } definition maressource { relation marelation: monsujet | monsujet2  relation mr2: monsujet | msj3  }`

	var input string = ""
	var schema string = ""
	var fschema string = ""
	var showHelp bool

	flag.StringVar(&schema, "schema", "", "Read schema")
	flag.StringVar(&fschema, "fschema", "", "Read schema file")
	flag.BoolVar(&showHelp, "help", false, "Show help message")
	flag.Parse()

	if (schema == "" && fschema == "") || (schema != "" && fschema != "") {
		fmt.Println("You must provide either -schema or -fschema, but not both.")
		printHelp()
		return
	}

	if schema != "" {
		input = schema
	}

	if fschema != "" {
		fileContent, err := os.ReadFile(fschema)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier : ", err)
		}
		input = string(fileContent)
	}

	if showHelp {
		printHelp()
		return
	}

	lexer := zinterpreter.NewLexer(input)
	lexer.NextToken()
	zschema, err := lexer.ReadZSchema()

	if err != nil {
		fmt.Println("syntax error:", err)
	} else {
		fmt.Println("parsed schema OK:", zschema)
	}
}
