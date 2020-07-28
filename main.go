package main

import (
	"fmt"
	"time"
	"runtime"
	"concurrency/fileoperation"
	"concurrency/dboperation"
)

var exitstatus string;

func main()  {	
	datachan:=make(chan string);
	var noOfProcs int;
	fmt.Println("Enter the maximum number of logical processors required:")
	fmt.Scanln(&noOfProcs)
	runtime.GOMAXPROCS(noOfProcs);
	go dboperation.DBwriter(datachan);
	go fileoperation.FileRead(datachan,&exitstatus);
	for{
		if exitstatus=="true"{
			break;
		}
		time.Sleep(100*time.Millisecond);
	}
}