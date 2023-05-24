package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"roongjinAssignment0/helper"
)

func main() {

	//------------ os.Args reads the arguments from the command line -----------

	if len(os.Args) == 2 {
		helpPtr := flag.Bool("help", false, "help")
		flag.Parse()

		if *helpPtr {
			fmt.Println("default command is http GET")
			fmt.Println("you can specify GET PUSH PUT DELETE for other purposes")
		} else {
			resp, e := http.Get("https://" + os.Args[1])
			if e != nil {
				log.Fatal(e)
			} else {
				fmt.Println(resp)
			}
		}
	} else {
		switch os.Args[1] {
		case "get":
			helper.GetRequest()
		case "post":
			helper.PostRequest()
		case "delete":
			helper.DeleteRequest()
		case "put":
			helper.PutRequest()
		default:
			fmt.Println("Command not found")
		}

	}
}
