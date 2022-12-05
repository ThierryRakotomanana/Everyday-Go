package main

import "fmt"
import "os"

func main(){
	name := os.Getenv("USER")
	fmt.Printf("Well done  %s for having your first Go\n", name)
}