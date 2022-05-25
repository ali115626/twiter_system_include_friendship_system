package Dao

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"time"
	cache "twiter/friendshipSystem/Dao/redis"
	//Dao "twiter/friendshipSystem/Dao/mysql"

	//"time"
)

func InsertFriend(userId string, friendId string)error{
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		return errors.New(fmt.Sprintf("open database error,err=", err))
	}
	createAtLocal := time.Now()
	fmt.Println("createAt=====",createAtLocal)
	//
	//shanghaiZone, _ := time.LoadLocation("Asia/Shanghai")
	//createAt := time.Now().String()
	//createAtLocal, _ := time.ParseInLocation("2006-01-02 15:04:05", createAt, shanghaiZone)
	//这里面要存两份
	//TODO 这里面要将其弄成事务吧
	//TODO  同时成功  同时失败  不然B的好友列表里面有A 但是A的好友列表中没有B
	//TODO insert的时候一定会调整那棵B+树 因为这个并不是按顺序的
	_, err = db.Exec("insert into FriendShip(fromUserId,toUserId,createAt) values(?,?,?)", userId, friendId, createAtLocal)
			if err != nil {
				fmt.Println("exec failed, err=", err)
				return errors.New(fmt.Sprintf("exec failed, err=", err))
			}
	_, err = db.Exec("insert into FriendShip(fromUserId,toUserId,createAt) values(?,?,?)",friendId,  userId, createAtLocal)
	if err != nil {
		fmt.Println("exec failed, err=", err)
		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}
	return nil
}
//	还是要像之前一样去insert进数据库  但是与此同时加一下redis   zset里面的操作

func InsertFriendNew(userId string, friendId string)error{
	//	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	//	if err != nil {
	//	fmt.Println("open database error,err=", err)
	//	return errors.New(fmt.Sprintf("open database error,err=", err))
	//}
	//	db,err:=MysqlInit()
		db:=MysqlConn

		fmt.Println("db=",db)

		createAtLocal := time.Now()
		fmt.Println("createAt=====",createAtLocal)
		//
		//shanghaiZone, _ := time.LoadLocation("Asia/Shanghai")
		//createAt := time.Now().String()
		//createAtLocal, _ := time.ParseInLocation("2006-01-02 15:04:05", createAt, shanghaiZone)
		//这里面要存两份
		//TODO 这里面要将其弄成事务吧
		//TODO  同时成功  同时失败  不然B的好友列表里面有A 但是A的好友列表中没有B
		//TODO insert的时候一定会调整那棵B+树 因为这个并不是按顺序的
		_, err := db.Exec("insert into FriendShip(fromUserId,toUserId,createAt) values(?,?,?)", userId, friendId, createAtLocal)
		if err != nil {
		fmt.Println("exec failed, err=", err)
		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}
		_, err = db.Exec("insert into FriendShip(fromUserId,toUserId,createAt) values(?,?,?)",friendId,  userId, createAtLocal)
		if err != nil {
		fmt.Println("exec failed, err=", err)
		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}

		//TODO 检测key-----------------------------------------------------------------------------

		//下面是redis的部分
		//把信息加到redis里面
		//这就相当于记录了每一个人的朋友圈
		conn, err := redis.Dial("tcp", "127.0.0.1:6379")
		if err != nil {
			fmt.Println("redis.Dial err=", err)
			//return
		}
		defer conn.Close()
		//followerId = userId
		boundLimit :=4
		//这个的话再反向操作一波  因为考虑到redis的事务
		FriendCricleKey := "FriendShip:FriendShipCricle:" + userId

		_, err = redis.Bool(conn.Do("sadd", FriendCricleKey,friendId))
		if err != nil {
			fmt.Println("sadd error,err=", err)
		}

		number1,err := redis.Int(conn.Do("scard", FriendCricleKey))
		if err!=nil{
			fmt.Println(err)
			return err
		}
		//加进去之后  看一下 数量吧
		userIdInt,err:=strconv.Atoi(userId)
		if err !=nil{
			return errors.New(fmt.Sprintf("convert into int error,err=",err))
		}
		if number1 > boundLimit{
			err:=ModifyIsCelebrityStatus(userIdInt)
			if err!=nil{
				//这里面要打印一下log
				fmt.Println("修改明星的状态失败！")
				return err
			}
			err=cache.AddcelebirtyZset(userIdInt)
			if err!=nil{
				fmt.Println("cache save celebirty error,err=",err)
				return err
			}

			//	将其加入明星的行列
			//sadd celebrity friendId
		//	todo 修改数据库了
		//	todo 或者加入zset
		//	todo 我这里面采用方案一
		//	todo 其实这个和下面的用一个函数就行呗
		//	ModifyIsCelebrityStatus
		//	Dao.
		//	update FriendShip set isToUserIdCelebrity="yes" where to UserId = userId
		}
		FriendCricleKey2 := "FriendShip:FriendShipCricle:" + friendId
		number2, err := redis.Int(conn.Do("sadd", FriendCricleKey2,userId))
		if err != nil {
			fmt.Println("sadd error,err=", err)
		}
		//这个要根据那个企业的数据  企业的经验来决定 因为我的机器可能比  企业的差一些    如果你嫌IO效率太低  你不行就能换成ssd硬盘
		//TODO if n > 1500{
		////	TODO添加那种粉丝的逻辑吧
		//
		//}
		friendIdInt,err:=strconv.Atoi(friendId)
		if err !=nil{
			return errors.New(fmt.Sprintf("convert into int error,err=",err))
		}
		if number2 > boundLimit{
			err:=ModifyIsCelebrityStatus(friendIdInt)
			if err!=nil{
				//这里面要打印一下log
				fmt.Println("修改明星的状态失败！")
				return err
			}
			err=cache.AddcelebirtyZset(friendIdInt)
			if err!=nil{
				fmt.Println("cache save celebirty error,err=",err)
				return err
			}
		//	将其加入明星的行列
		//sadd celebrity friendId
		}
		return nil
}
//update FriendShip set isToUserIdCelebrity="yes" where to UserId = friendId
//这个要根据那个企业的数据  企业的经验来决定 因为我的机器可能比  企业的差一些    如果你嫌IO效率太低  你不行就能换成ssd硬盘
//TODO if n > 1500{
////	TODO添加那种粉丝的逻辑吧
//
//}
//select * from FriendShip;
//	+--------------+------------+----------+---------------------+---------------------+
//	| FriendshipId | fromUserId | toUserId | createAt            | isToUserIdCelebrity
//convert into int

func ModifyIsCelebrityStatus(userId int) error {
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		return errors.New(fmt.Sprintf("open database error,err=", err))
	}
	res, err := db.Exec("update FriendShip set isToUserIdCelebrity=? where toUserId = ?","yes", userId)
	if err != nil {
		fmt.Println("exec failed, err=", err)
		return errors.New(fmt.Sprintf("exec failed, err=", err))
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ",err)
	}
	fmt.Println("update succ:",row)
	return nil
}


//update
//FriendShip
//set
//isToUserIdCelebrity = "yes"
//where
//to
//UserId = userId
////update FriendShip set isToUserIdCelebrity="yes" where to UserId = friendId
//res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 1)
//if err != nil {
//	fmt.Println("exec failed, ", err)
//	return
//}
//row, err := res.RowsAffected()
//if err != nil {
//	fmt.Println("rows failed, ",err)
//}
//fmt.Println("update succ:",row)