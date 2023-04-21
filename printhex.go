package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	bs, _ := io.ReadAll(os.Stdin)
	fmt.Println(bs)
}
