package main

import (
	"fmt"
	"time"
)

func main() {
	dn := time.Now()
	fmt.Println(dn)

	dme := dn.Format("2/1/2006")
	fmt.Println(dme)
	tn := dn.Format("15:04")
	hms := dn.Format("15:04-05")
	fmt.Println(tn)
	fmt.Println(hms)
}
