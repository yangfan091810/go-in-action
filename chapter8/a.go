package main

import "log"

func init(){
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main(){
	defer log.Println("aaaaa")
	defer log.Println("bbbbb")
	log.Println("message")
	//log.Fatalln("fatal message")
	log.Panicln("panic message")
}
