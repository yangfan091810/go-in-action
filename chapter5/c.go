package main

import (
	"fmt"
	"io"
	"bytes"
	"os"
)

func main(){
	var b bytes.Buffer

	b.Write([]byte("hello"))

	fmt.Fprintf(&b, " world \n")

	io.Copy(os.Stdout, &b)
}
