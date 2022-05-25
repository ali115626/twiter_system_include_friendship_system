package Dao

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

//这是普通朋友

//把这两个弄完  再去结合一下    测试 一下能不能把明星和普通人都给   拿到

func SelectCommonFriendsFromFriendTable(userId int) ([]string, error) {
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}
	rows, err := db.Query("select toUserId from FriendShip where fromUserId=? and isToUserIdCelebrity=?", userId,"no")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("select data error: %v\n", err))
	}
	//defer func() {
	//	rows.Close() // 会释放数据库连接
	//}()
	var toUserId int
	var friendsList []string
	// 循环读取数据
	for rows.Next() {
		err := rows.Scan(&toUserId)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil, err
		}

		toUserIdStr:=strconv.Itoa(toUserId)

		friendsList = append(friendsList, toUserIdStr)
	}
	rows.Close()

	CommonFriendsList :=friendsList
	return CommonFriendsList, nil
}


func SelectStarsFromFriendTable(userId int) ([]string, error) {
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}
	rows, err := db.Query("select toUserId from FriendShip where fromUserId=? and isToUserIdCelebrity=?", userId,"yes")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("select data error: %v\n", err))
	}
	//defer func() {
	//	rows.Close() // 会释放数据库连接
	//}()
	var toUserId int
	var friendsList []string
	// 循环读取数据
	for rows.Next() {
		err := rows.Scan(&toUserId)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil, err
		}
		toUserIdStr:=strconv.Itoa(toUserId)

		friendsList = append(friendsList, toUserIdStr)
	}
	rows.Close()

	CommonFriendsList :=friendsList
	return CommonFriendsList, nil
}



func SelectFriendsFromFriendTable(userId int) ([]int, error) {
	db, err := sql.Open("mysql", "root:123456@/twiter_scheme?charset=utf8&loc=Local")
	if err != nil {
		fmt.Println("open database error,err=", err)
		return nil, errors.New(fmt.Sprintf("open database error,err=", err))
	}
	rows, err := db.Query("select toUserId from FriendShip where fromUserId=?", userId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("select data error: %v\n", err))
	}
	defer func() {
		rows.Close() // 会释放数据库连接
	}()
	var toUserId int
	var friendsList []int
	// 循环读取数据
	for rows.Next() {
		err := rows.Scan(&toUserId)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return nil, err
		}
		friendsList = append(friendsList, toUserId)
	}
	return friendsList, nil
}
