package Logic

import (
	"fmt"
	"net/http"
	"twiter/UserSystem/utitl"
)

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	requestMap := r.Form
	//
	//if strings.Join(requestMap["username"], "") == "" {
	//	fmt.Fprintf(w, "用户名不能为空")
	//	return
	//}

	//userId := requestMap["userId"][0]


	userName := requestMap["username"][0]

	passwd := requestMap["passwd"][0]


	email := requestMap["email"][0]


	phone := requestMap["phone"][0]


	//打开mysql数据库
	//TODO 数据库加密都用什么算法呢
	//
	//passwdByte := []byte(strr)
	//has := md5.Sum(passwdByte)
	//fmt.Println(len(fmt.Sprintf("%x", has)))
	_, err = utitl.Db.Exec("insert into User(userName, passwd, email,phone) values(?, ?, ?,?)", userName, passwd, email, phone)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}
	fmt.Fprintf(w, "welcome, register sucessfully!") // 这个写入到 w 的是输出到客户端的







	//在数据库中insert



	//注册结果返回给端



	//



}






