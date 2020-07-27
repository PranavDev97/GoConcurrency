package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"io"
	"time"
	"runtime"
)

var datachan chan string;
var exitstatus string;

func DBwriter(ch chan string){
	fmt.Println("4");
	db,err := sql.Open("mysql","root:password@/personal_project");
	if err!=nil {
		fmt.Println("db connection error:",err)
	}
	stmt,err := db.Prepare("INSERT INTO csvdata(data) values(?)");
	if err!=nil {
		fmt.Println("statement creation failed:",err)
	}
	for {
		data:=<-ch;
		fmt.Println("7");
		if data=="endoffile" {
			fmt.Println("9");
			break;
		}
		res,err := stmt.Exec(data);
		if err!=nil {
			fmt.Println("insertion failed:",err)
		}
		fmt.Println(res);
	}
}

func FileRead(ch chan string){
	fmt.Println("5");
	csvFile,err := os.Open("data.csv");
	if err != nil {
		fmt.Println("File open error :",err)
	}
	fileData := csv.NewReader(csvFile);
	for{
		fmt.Println("6");
		data , e :=fileData.Read();
		if e == io.EOF{
			fmt.Println("8");
			break;
		}
		ch<-data[0];
		if e!=nil {
			fmt.Println("data read error :",e)
		}
		fmt.Println(data[0]);
	}
	ch<-"endoffile"
	exitstatus="true";
}

func main()  {
	runtime.GOMAXPROCS(1);
	fmt.Println("1");
	go DBwriter(datachan);
	fmt.Println("2");
	go FileRead(datachan);
	fmt.Println("3");
	for{
		if exitstatus=="true"{
			break;
		}
		time.Sleep(100*time.Millisecond);
	}
	
}