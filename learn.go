package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
)

// command line arguments
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	fmt.Println(printSpaceSeparatedCommandLineArgument(os.Args))
	fmt.Println(printCommandLineArgumentFromAIndex(os.Args, 2))
	printCountAndTextFromFile()
	countFromStandardInput()
	lissajous(os.Stdout)
	pointerExample()
	fmt.Println(f())
	v := 2
	incr(&v)
	fmt.Println(incr(&v))
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func printSpaceSeparatedCommandLineArgument(args []string) string {
	var res string = strings.Join(os.Args, " ")
	return res
}

func printCommandLineArgumentFromAIndex(args []string, index int) string {
	var res, sep string
	for _, args := range os.Args[index:] {
		res += sep + args
		sep = " "
	}
	return res
}

/*
Dup1 prints the text of each line that appears more than
once in the standard input, preceded by its count.
*/
func countFromStandardInput() {
	count := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++

		if input.Text() == "" {
			break
		}
	}
	for line, n := range count {
		fmt.Println(line, n)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/*
	prints the count and text of lines that appear more than once

in the input.  It reads from stdin or from a list of named files.
*/
func printCountAndTextFromFile() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for lines, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, lines)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {

}

// animated gif
var palate = []color.Color{color.Black, color.White}

const (
	blackIndex = 0
	whiteIndex = 1
)

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 1000
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palate)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// pointer
func pointerExample() {
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x)
}

func f() *int {
	a := 2
	return &a
}

// modifies the actual value
func incr(p *int) int {
	*p++
	return *p
}
