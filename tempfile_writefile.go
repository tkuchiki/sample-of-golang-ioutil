package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) {
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("###############")
	fmt.Println(string(b))
	fmt.Println("###############")
	fmt.Println()
}

func main() {
	fp, err := ioutil.TempFile("./dir", "__prefix__")

	defer os.Remove(fp.Name())
	defer fp.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%T \n", fp)
	fmt.Printf("%v \n", fp.Name())
	fmt.Println()

	b := []byte(`hoge
foobar`)

	fp2, err := ioutil.TempFile("", "")

	defer os.Remove(fp2.Name())
	defer fp2.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(fp2.Name())

	ioutil.WriteFile(fp2.Name(), b, 0644)

	ReadFile(fp2.Name())
}
