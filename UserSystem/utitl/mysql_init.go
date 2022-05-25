package utitl

import (
	"database/sql"
	//"database/sql"
	//"database/sql"
	//"database/sql"
	//ll "database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
)
var Db *sql.DB
var err error

func MysqlInit(){
	Db, err = sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")

	if err != nil{
		fmt.Println("mysql open error err=",err)
	}



}
