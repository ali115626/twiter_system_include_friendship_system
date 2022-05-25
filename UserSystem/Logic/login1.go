package Logic

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
	"twiter/UserSystem/utitl"
)

//用邮箱登录

func Login1(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}

	requestMap := r.Form

	email := requestMap["email"][0]

	passwd := requestMap["passwd"][0]

	var retrivePassswd string
	var userId string

	///给数据库建立索引   给username email passwd建立索引  这块给email建立前缀索引吗？
	err = utitl.Db.QueryRow("SELECT userId,passwd FROM User WHERE email=?", email).Scan(&retrivePassswd, &userId)

	//rows, err :=utitl.Db.Query("select passwd from blog_info where user_id=?", userName
	//if err != nil {
	//	fmt.Printf("insert data error: %v\n", err)
	//	return
	//}

	if passwd == retrivePassswd {
		fmt.Fprintf(w, "login successfully!")

		//这里面加入set_cookie的逻辑
		//id := uuid.NewV4()
		sessionKey := uuid.New().String()
		//
		//fmt.Println(messageId)
		//session_key := id.String()

		cookie := http.Cookie{Name: "sessionStr", Value: sessionKey}
		http.SetCookie(w, &cookie)

		//
		//select column_name
		//	from information_schema.'key_column_usage'
		//	where table_name='表名称' and constraint_name='primary'
		//从这以后 几分钟过期
		//expire_at:=time.Now().Unix()+time.Second.String() .Seconds()*60*24
		expirationTime := time.Now().Add(time.Minute * 30).Unix()

		//expire_at := strconv.Itoa(expirationTime)
		expireAt := strconv.FormatInt(expirationTime, 10)

		//TODO 这样做  天生支持多台设备登录
		_, err = utitl.Db.Exec("insert into User(session_key, userId, expire_at) values(?, ?, ?)", sessionKey, userId, expireAt)
		if err != nil {
			fmt.Println("exec failed, ", err)
			return
		}

		fmt.Fprintf(w, "登陆成功!")

		//TODO 你每次登录成功后，向那个表格中写入  session_key userID create_at

		//insert到数据库中
		//生成uuid   给他们建一个表格   建一个数据库

	} else {
		fmt.Fprintf(w, "密码错误 ，请重新输入!")

	}

}
