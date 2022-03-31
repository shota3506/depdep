package main

import (
	"github.com/shota3506/depdep"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(depdep.Analyzer) }
