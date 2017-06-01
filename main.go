package main

import (
	"fmt"
	"time"

	"github.com/elgs/gostrgen"
)

func main() {

	for {
		str, _ := getCharacters()
		fmt.Println(str)
	}
}

func getCharacters() (string, error) {
	charsToGenerate := 20
	charSet := gostrgen.Lower | gostrgen.Digit
	includes := "[]{}<>" // optionally include some additional letters
	excludes := "Ol"     //exclude big 'O' and small 'l' to avoid confusion with zero and one.

	str, err := gostrgen.RandGen(charsToGenerate, charSet, includes, excludes)
	if err != nil {
		fmt.Println(err)
		return "", err

	}
	time.Sleep(1 * time.Second)
	return str, nil

}
