package nosql_logic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"twiter/UserSystem/utitl"
	//"github.com/gomodule/redigo/redis"

)

//TODO 在nosql的数据库中  这个user表怎么差

//TODO  相当SQL 型数据库   按照索引去存

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

	userId := requestMap["userId"][0]


	userName := requestMap["username"][0]

	passwd := requestMap["passwd"][0]


	email := requestMap["email"][0]


	phone := requestMap["phone"][0]
	//TODO 数据库加密都用什么算法呢



	//TODO 把redis数据库打开一下

	//TODO 这里面建立三张表格吧
	//commentKey := "Comment:" + paperId + ":" + commentId
	//data, err := redis.Bytes(conn.Do("GET", commentKey))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//}




	UserInfoKeyId := "UserInfo"+userId

	jmap1 := map[string]string{"userName":userName,"passwd":passwd,"email":email,"phone":phone}

	//var commentMap map[string]string
	jsonValue, err := json.Marshal(jmap1)
	_, err = utitl.Conn.Do("set", UserInfoKeyId, jsonValue)
	if err != nil {
		fmt.Println("set key error,err=", err)
	}


	//todo  这里面多存几份数据   相当于是多建了几份索引


	UserInfoKeyUserName := "UserInfo"+userName
	_, err = utitl.Conn.Do("set", UserInfoKeyUserName, userId)
	if err != nil {
		fmt.Println("set key error,err=", err)
	}



	UserInfoKeyEmail := "UserInfo"+email
	_, err = utitl.Conn.Do("set", UserInfoKeyUserName, UserInfoKeyEmail)
	if err != nil {
		fmt.Println("set key error,err=", err)
	}









	//_, err = utitl.Db.Exec("insert into User(userName, passwd, email,phone) values(?, ?, ?,?)", userName, passwd, email, phone)
	//if err != nil {
	//	fmt.Println("exec failed, ", err)
	//	return
	//}
	fmt.Fprintf(w, "welcome, register sucessfully!") // 这个写入到 w 的是输出到客户端的







	//在数据库中insert



	//注册结果返回给端



	//



}

