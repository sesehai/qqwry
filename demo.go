package main

import (
	"fmt"
)

func main() {
	var ip = "112.0.91.210"
	var qqwryfile = "./qqwry.dat"
	fmt.Println(ip)
	file, _ := qqwry.Getqqdata(qqwryfile)
	country, _ := qqwry.Getlocation(file, ip)
	fmt.Println(country)
}
