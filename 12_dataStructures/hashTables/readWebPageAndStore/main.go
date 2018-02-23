package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//using hash Tables

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	/*
		byteslice, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("%s", byteslice)
	*/
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	buckets := make([][]string, 12)
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	for scanner.Scan() {
		//n := hashbucket(scanner.Text()) //Way 1
		word := scanner.Text()
		n := hashbucket(word, 12) //Way 2

		buckets[n] = append(buckets[n], word)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	//fmt.Println(buckets[65:122])
	fmt.Println(buckets[6])
}

/* way 1 - uneven distribution
func hashbucket(word string) int {
	return int(word[0])
}
*/

func hashbucket(word string, buckets int) int {
	// letter := int(word[0])
	// bucket := letter % buckets
	// return bucket
	var sum int
	for _, v := range word {
		sum += int(v)

	}
	return sum % buckets
}
