package stage1

import "flag"

func helloFunc() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}
