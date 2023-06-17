package main

import (
	"os"
	"go-exercises/lissajousFunc"
)

func main(){
	lissajousFunc.Lissajous(os.Stdout, lissajousFunc.LissajousOpts{-1, -1, -1, -1, -1})
}