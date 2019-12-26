package main

import (
	"./lib"
	"flag"
	"fmt"
	"net"
	"os"
)

var VERSION = `0.0.0`
var BUILD = `dev`

const AUTHOR = `Pekka Järvinen`

func main() {
	flag.Usage = func() {
		fmt.Printf(`IP range to CIDR v%v build %v`+"\n", VERSION, BUILD)
		fmt.Printf(`(c) %v 2019-`+"\n", AUTHOR)
		fmt.Println(`Usage:`)
		fmt.Printf(`  %v <start IP> <end IP>`+"\n", os.Args[0])
		fmt.Println()
		fmt.Println(`Example:`)
		fmt.Printf(`  %v 192.168.0.0 192.168.0.255`+"\n", os.Args[0])
		fmt.Printf(`    Output: 192.168.0.0/24` + "\n")
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	start := net.ParseIP(flag.Arg(0))
	end := net.ParseIP(flag.Arg(1))

	ip, mask, err := lib.ToCidr(start, end)

	if err != nil {
		fmt.Printf(`error: %v`, err)
		os.Exit(1)
	}

	fmt.Printf(`%s/%d`, ip, mask)
}
