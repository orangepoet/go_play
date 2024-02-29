package cmd

import (
	"flag"
	"fmt"
	"os"
)

func foo() {
	name := flag.String("name", "world", "the name to greet.")

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("hello, %s\n", *name)
	} else {
		fmt.Printf("arg0->%s\n", flag.Arg(0))
		if flag.Arg(0) == "list" {
			files, _ := os.Open(".")
			defer files.Close()

			fileInfo, _ := files.Readdir(-1)
			for _, file := range fileInfo {
				fmt.Println(file.Name())
			}
		} else {
			fmt.Printf("hello, %s\n", *name)
		}
	}
}
