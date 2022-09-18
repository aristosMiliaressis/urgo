package main

import (
	"flag"
	"fmt"
)

type strslice []string
type Value interface {
	String() string
	Set(string) error
}

func (s *strslice) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *strslice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

type Options struct {
	Request     RequestOptions
	Extraction  ExtractionOptions
	Output      OutputOptions
	Performance PerformanceOptions
}

type RequestOptions struct {
	Method  string
	Headers strslice
}

type ExtractionOptions struct {
	StatusCode   bool
	ContentType  bool
	ResponseTime bool
	Title        bool
	Headers      strslice
	Regex        strslice
	FaviconHash  bool
}

type OutputOptions struct {
	Json      bool
	PrintHelp bool
}

type PerformanceOptions struct {
	Threads int
}

func (opts *Options) Parse() {
	flag.BoolVar(&opts.Output.PrintHelp, "h", false, "prints this help page")
	flag.BoolVar(&opts.Output.PrintHelp, "help", false, "prints this help page")
	flag.BoolVar(&opts.Output.Json, "oJ", true, "outputs in JSONL(ine) format")
	flag.BoolVar(&opts.Output.Json, "json", true, "outputs in JSONL(ine) format")
	flag.BoolVar(&opts.Extraction.StatusCode, "sC", false, "extracts the response status code")
	flag.BoolVar(&opts.Extraction.StatusCode, "status-code", false, "extracts the response status code")
	flag.BoolVar(&opts.Extraction.ResponseTime, "rT", false, "extracts the response time")
	flag.BoolVar(&opts.Extraction.ResponseTime, "response-time", false, "extracts the response time")
	flag.BoolVar(&opts.Extraction.Title, "T", false, "extracts the page title")
	flag.BoolVar(&opts.Extraction.Title, "title", false, "extracts the page title")
	flag.BoolVar(&opts.Extraction.FaviconHash, "f", false, "extracts favicon hash")
	flag.BoolVar(&opts.Extraction.FaviconHash, "favicon", false, "extracts favicon hash")
	flag.Var(&opts.Extraction.Regex, "rE", "regex extract (can be used multiple times)")
	flag.Var(&opts.Extraction.Regex, "regex-extract", "regex extract (can be used multiple times)")
	flag.StringVar(&opts.Request.Method, "m", "", "specifies http method to use")
	flag.StringVar(&opts.Request.Method, "method", "", "specifies http method to use")
	flag.Var(&opts.Request.Headers, "H", "adds request header (can be used multiple times)")
	flag.Var(&opts.Request.Headers, "req-header", "adds request header (can be used multiple times)")
	flag.Var(&opts.Extraction.Headers, "rH", "extracts the specified response header (can be used multiple times)")
	flag.Var(&opts.Extraction.Headers, "resp-header", "extracts the specified response header (can be used multiple times)")
	flag.IntVar(&opts.Performance.Threads, "t", 20, "specifies number of threads to use [default: 20]")
	flag.IntVar(&opts.Performance.Threads, "threads", 20, "specifies number of threads to use [default: 20]")
	flag.Parse()
}

func print_help() {
	fmt.Println("USAGE: cat urls.txt | urgo [OPTIONS]")
	fmt.Println()
	fmt.Println("REQUEST OPTIONS:")
	fmt.Println("\t-m, --method")
	fmt.Println("\t\tspecifies http method to use")
	fmt.Println("\t-H, --req-header")
	fmt.Println("\t\tadds request header (can be used multiple times)")
	fmt.Println()
	fmt.Println("EXTRACTION OPTIONS:")
	fmt.Println("\t-sC, --status-code")
	fmt.Println("\t\ttextracts the response status code")
	fmt.Println("\t-rT, --response-time")
	fmt.Println("\t\textracts the response time")
	fmt.Println("\t-rH, --resp-header")
	fmt.Println("\t\textracts the specified response header (can be used multiple times)")
	fmt.Println("\t-T, --title")
	fmt.Println("\t\textracts the page title")
	// fmt.Println("\t-rE, --regex-extract")
	// fmt.Println("\t\tregex extract (can be used multiple times)")
	fmt.Println("\t-f, --favicon")
	fmt.Println("\t\textracts favicon hash")
	fmt.Println()
	fmt.Println("OUTPUT OPTIONS:")
	fmt.Println("\t-oJ, --json")
	fmt.Println("\t\toutputs in JSONL(ine) format")
	fmt.Println("\t-h, --help")
	fmt.Println("\t\tprints this help page")
	fmt.Println()
	fmt.Println("PERFORMANCE OPTIONS:")
	fmt.Println("\t-t, --threads")
	fmt.Println("\t\tspecifies number of threads to use [default: 20]")
}
