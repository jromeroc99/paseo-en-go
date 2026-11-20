// Command paseo serves the course locally or builds it as a static site.
//
//	paseo serve [-addr :8080]
//	paseo build [-out dist] [-base /]
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"paseo-en-go/content"
	"paseo-en-go/internal/site"
	"paseo-en-go/web"
)

const usage = `usage:
  paseo serve [-addr :8080]           serve the course for local development
  paseo build [-out dist] [-base /]   write the course as a static site`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

	var err error
	switch os.Args[1] {
	case "serve":
		err = serve(os.Args[2:])
	case "build":
		err = build(os.Args[2:])
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func serve(args []string) error {
	flags := flag.NewFlagSet("serve", flag.ExitOnError)
	addr := flags.String("addr", ":8080", "address to listen on")
	flags.Parse(args)

	s, err := site.New(content.Course, content.FS, web.FS, "/")
	if err != nil {
		return err
	}
	log.Printf("serving on http://localhost%s", *addr)
	return http.ListenAndServe(*addr, s.Handler())
}

func build(args []string) error {
	flags := flag.NewFlagSet("build", flag.ExitOnError)
	out := flags.String("out", "dist", "output directory")
	base := flags.String("base", "/", "URL prefix the site is published under")
	flags.Parse(args)

	s, err := site.New(content.Course, content.FS, web.FS, *base)
	if err != nil {
		return err
	}
	if err := s.Build(*out); err != nil {
		return err
	}
	log.Printf("site written to %s", *out)
	return nil
}
