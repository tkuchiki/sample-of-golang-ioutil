package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dir, err := ioutil.TempDir("", "")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%v \n", dir)

	os.RemoveAll(dir)

	dir, err = ioutil.TempDir("./dir", "__prefix__")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%v \n", dir)

	os.RemoveAll(dir)
}
