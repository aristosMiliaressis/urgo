package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	var opts Options

	opts.Parse()
	if opts.Output.PrintHelp {
		print_help()
		return
	}

	ext := &Extractor{
		ExtractionOptions: opts.Extraction,
		RequestOptions:    opts.Request,
	}

	var out Outputer
	out = GreppableOutputer{}
	if opts.Output.Json {
		out = JsonOutputer{}
	}

	s := bufio.NewScanner(os.Stdin)
	wg := sync.WaitGroup{}
	guard := make(chan struct{}, opts.Performance.Threads)
	for s.Scan() {
		guard <- struct{}{}
		wg.Add(1)
		go func(url string) {
			metadata, err := ext.Extract(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err.Error())
			}
			out.Output(metadata)
			wg.Done()
			<-guard
		}(s.Text())
	}
	wg.Wait()
}
