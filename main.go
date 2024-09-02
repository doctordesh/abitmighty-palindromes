package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/profile"
)

const LIMIT = 1_000_000_000

func main() {
	limit := LIMIT
	prof := false

	flag.IntVar(&limit, "limit", LIMIT, "The limit of the calculation")
	flag.BoolVar(&prof, "profile", false, "Run CPU profile")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Version number is needed")
		os.Exit(1)
	}

	var p interface{}
	if prof {
		p = profile.Start(profile.CPUProfile, profile.ProfilePath("."), profile.NoShutdownHook)
	}

	var sum, count int

	switch flag.Args()[0] {
	case "v1":
		sum, count = PalindromeSum_v1(limit)
	case "v2":
		sum, count = PalindromeSum_v2(limit)
	case "v3":
		sum, count = PalindromeSum_v3(limit)
	case "v4":
		sum, count = PalindromeSum_v4(limit)
	default:
		fmt.Printf("no option for '%s'\n", flag.Args()[0])
		os.Exit(2)
	}
	fmt.Printf("Count: %d or %s in base22\n", count, NumberToBase_v1(count, 22))
	fmt.Printf("Sum:   %d or %s in base22\n", sum, NumberToBase_v1(sum, 22))
	fmt.Printf("Limit: %d\n", limit)

	if prof {
		(p.(*profile.Profile)).Stop()
	}
}
