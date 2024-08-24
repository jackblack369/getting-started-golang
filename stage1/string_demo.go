package main

import "strings"

func main() {
	url := "sqlite:///Users/dongwei/.curveadm/data/curveadm.db"
	dataSourceName := strings.TrimPrefix(url, "sqlite://")
	println("dataSourceName is ", dataSourceName)
}
