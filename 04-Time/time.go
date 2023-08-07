package main

import (
	"fmt"
	"time"
)

func main()  {
	present := time.Now()
	create := time.Date(2020, time.April,10,23,23,0,0,time.UTC)

	fmt.Println("Present Date Without Format :",present)
	fmt.Println("Present Date With Format :", present.Format("02-01-2006 15:04:05 Monday"))
	fmt.Println("Created Date :",create)
}