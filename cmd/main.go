package main

import (
	"flag"
	"io/ioutil"

	"github.com/camopy/sitemap-generator/generator"
)

func main() {
	parallelPtr := flag.Int("parallel", 1, "number of parallel workers")
	outputFilePtr := flag.String("output-file", "", "output file path")
	maxDepthPtr := flag.Int("max-depth", 0, "max depth of url navigation recursion")

	flag.Parse()

	if *parallelPtr == -1 || *outputFilePtr == "" || *maxDepthPtr == -1 {
		flag.Usage()
		return
	}

	baseUrl := flag.Arg(0)

	g := generator.New(baseUrl, *maxDepthPtr, *parallelPtr)
	xml, err := g.Generate()
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(*outputFilePtr, xml, 0644)
}
