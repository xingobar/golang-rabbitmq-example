package main

import "fmt"

func main() {

}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Println("msg: ", err, msg)
		panic(msg)
	}
}

