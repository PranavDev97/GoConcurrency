package dboperation;

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func DBwriter(datachan chan string){
	db,err := sql.Open("mysql","root:password@/personal_project");
	if err!=nil {
		fmt.Println("db connection error:",err)
	}
	stmt,err := db.Prepare("INSERT INTO csvdata(data) values(?)");
	if err!=nil {
		fmt.Println("statement creation failed:",err)
	}
	for {
		data:=<-datachan;
		if data=="endoffile" {
			break;
		}
		_,err := stmt.Exec(data);
		if err!=nil {
			fmt.Println("insertion failed:",err)
		}
	}
}

