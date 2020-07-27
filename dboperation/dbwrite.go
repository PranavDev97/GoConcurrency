package dboperation;

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

func DBwriter(){
	db,err := sql.Open("mysql","root:password@/personal_project");
	if err!=nil {
		fmt.Println("db connection error:",err)
	}
	stmt,err := db.Prepare("INSERT INTO csvdata(data) values(?)");
	if err!=nil {
		fmt.Println("statement creation failed:",err)
	}
	res,err := stmt.Exec("hi");
	if err!=nil {
		fmt.Println("insertion failed:",err)
	}
	fmt.Println(res);
}

