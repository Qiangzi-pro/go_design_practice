package main

import "fmt"

type Def struct {
	s int
}

func main() {
	tmp := map[string]Def{}
	e := tmp["lunch"]

	fmt.Printf("%v, %T", e, e)
}
