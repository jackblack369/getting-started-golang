package main

import "flag"

func helloFunc() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
