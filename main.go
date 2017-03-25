package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	chars = map[rune]rune{
		';': '\u037E',
		' ': '\u2002',
	}
	charsReverse = map[rune]rune{
		'\u037E': ';',
		'\u2002': ' ',
	}
)

func main() {
	var app = kingpin.New("godamnit", "By default takes text from stdin, encodes it and outputs it back to stdout. Can instead take file paths with the args below.")
	var shouldDecode = app.Flag("decode", "Decode instead of encode").Short('d').Bool()
	var shouldTest = app.Flag("test", "Dumps comma seperated unicode code points to stdout").Short('t').Bool()
	var infilePath = app.Arg("input file", "Input file path").String()
	var outfilePath = app.Arg("output file", "Output file path").String()
	kingpin.MustParse(app.Parse(os.Args[1:]))

	var text string

	if len(*infilePath) > 1 {
		f, err := ioutil.ReadFile(*infilePath)
		if err != nil {
			log.Println(err)
		}
		text = string(f)
	} else {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Println(err)
		}
		text = string(bytes)
	}

	if *shouldTest {
		str := encodedecode(text, shouldDecode)
		for _, c := range str {
			fmt.Printf("%U,", c)
		}
		return
	}

	if len(*outfilePath) > 1 {
		err := ioutil.WriteFile(*outfilePath, []byte(encodedecode(text, shouldDecode)), os.FileMode(0644))
		if err != nil {
			log.Println(err)
		}
	} else {
		fmt.Print(encodedecode(text, shouldDecode))
	}
}

func encodedecode(in string, shouldDecode *bool) string {
	var new string
	if *shouldDecode {
		for _, v := range in {
			if val, ok := charsReverse[v]; ok {
				new += string(val)
			} else {
				new += string(v)
			}
		}
		return new
	}
	for _, v := range in {
		if val, ok := chars[v]; ok {
			new += string(val)
		} else {
			new += string(v)
		}
	}
	return new
}
