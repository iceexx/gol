package main;

import (
	"fmt"
	"os"
	"strconv"
	"strings"
);

func main() {
	var s string;
	for i:=0; i<len(os.Args); i++ {
		fmt.Println(os.Args[i]);
		s += os.Args[i] + " ";
	}
	fmt.Println(s);

	for index, arg := range os.Args {
		fmt.Println(strconv.Itoa(index) + " " + arg)
	}

	for _, arg := range os.Args {
		fmt.Println(arg);
	}

	fmt.Println(strings.Join(os.Args[1:], " * "));
	fmt.Println(os.Args[1:]);

	s1 := "s1";
	var s2 string;
	var s3 = "s3";
	var s4 string = "s4";

	fmt.Println(s1);
	fmt.Println(s2);
	fmt.Println(s3);
	fmt.Println(s4);
}