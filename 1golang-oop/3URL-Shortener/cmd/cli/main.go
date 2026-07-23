package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"URL-shortener/internal/config"
	"URL-shortener/internal/repository"
	"URL-shortener/internal/service"
)

func main() {
	cfg := config.Load()
	repo := repository.NewFileURLRepository(cfg.StorePath)
	svc := service.NewURLService(repo)

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "shorten":
		fs := flag.NewFlagSet("shorten", flag.ExitOnError)
		fs.Parse(os.Args[2:])
		if fs.NArg() == 0 {
			fail("shorten requires a URL, e.g. `urlmgr shorten https://example.com`")
		}
		longURL := joinArgs(fs.Args())
		url, err := svc.Shorten(longURL)
		if err != nil {
			fail(err.Error())
		}
		fmt.Printf("shortened: %s -> %s/%s\n", longURL, cfg.BaseURL, url.Code)

	case "resolve":
		code := firstArg(os.Args[2:], "resolve")
		url, err := svc.Resolve(code)
		if err != nil {
			fail(err.Error())
		}
		fmt.Printf("%s -> %s\n", code, url.LongURL)

	case "list":
		urls, err := svc.ListAll()
		if err != nil {
			fail(err.Error())
		}
		if len(urls) == 0 {
			fmt.Println("no URLs yet")
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)
		fmt.Fprintln(w, "CODE\tLONG URL")
		for _, u := range urls {
			fmt.Fprintf(w, "%s\t%s\n", u.Code, u.LongURL)
		}
		w.Flush()

	case "delete":
		code := firstArg(os.Args[2:], "delete")
		if err := svc.Delete(code); err != nil {
			fail(err.Error())
		}
		fmt.Printf("deleted short URL %s\n", code)

	case "help", "-h", "--help":
		usage()

	default:
		fail(fmt.Sprintf("unknown command %q", os.Args[1]))
	}
}

func firstArg(args []string, cmd string) string {
	if len(args) == 0 {
		fail(fmt.Sprintf("%s requires a short code, e.g. `urlmgr %s abc123`", cmd, cmd))
	}
	return args[0]
}

func joinArgs(args []string) string {
	out := ""
	for i, a := range args {
		if i > 0 {
			out += " "
		}
		out += a
	}
	return out
}

func fail(msg string) {
	fmt.Fprintln(os.Stderr, "error:", msg)
	os.Exit(1)
}

func usage() {
	fmt.Print(`urlmgr - a tiny CLI URL shortener

Usage:
  urlmgr shorten <url>     shorten a URL
  urlmgr resolve <code>    resolve a short code to the original URL
  urlmgr list              list all shortened URLs
  urlmgr delete <code>     delete a short URL
  urlmgr help              show this help

URLs are stored in URL_STORE (default ~/.urlshortener/urls.json).
`)
}