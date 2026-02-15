package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	var w = tabwriter.NewWriter(
		os.Stdout, // output
		0,         // min width
		0,         // tab width
		1,         // padding
		'.',       // padding character
		tabwriter.AlignRight|tabwriter.Debug,
	)

	// Indent is misaligned when using multi-bytes character.
	//
	// ...a|...b|c
	// ..aa|..bb|cc
	// .aaa|.bbb|
	// ...あ|...い|う

	_, err := fmt.Fprintln(w, "a\tb\tc")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(w, "aa\tbb\tcc")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(w, "aaa\tbbb\t")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintln(w, "あ\tい\tう")
	if err != nil {
		log.Fatal(err)
	}

	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
