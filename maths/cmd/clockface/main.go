package main

import (
	"os"
	"time"

	clockface "github.com/stevenstank/learning-go/maths"
)

func main() {
	clockface.SVGWriter(os.Stdout, time.Now())
}
