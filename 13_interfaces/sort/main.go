package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	s := []string{"zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.Sort(sort.StringSlice(s)) //method 1
	sort.StringSlice(s).Sort()     //method 2
	sort.Strings(s)                //method 3
	fmt.Println(s)

}
