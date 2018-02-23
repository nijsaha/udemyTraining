package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	/*
		bs, _ := ioutil.ReadAll(res.Body)
		str := string(bs)
		fmt.Println(str)
		fmt.Println(len(str))
	*/

	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words[sc.Text()] = ""

	}
	if err := sc.Err(); err != nil {
		fmt.Println(err)
		//fmt.Fprintln(.Stderr, "reading input: ", err)
	}

	i := 0
	for k := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}
