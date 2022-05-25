package script

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func AddCommonCelebrityFriendsToSomeone() {

	//	TODO 加明星的好友

	for {
		//	这里面再去postman
		params := url.Values{}
		Url, err := url.Parse("http://baidu.com?fd=fdsf")
		if err != nil {
			panic(err.Error())
		}
		params.Set("a", "fdfds")
		params.Set("id", string("1"))
		//如果参数中有中文参数,这个方法会进行URLEncode
		Url.RawQuery = params.Encode()
		urlPath := Url.String()
		resp, err := http.Get(urlPath)
		defer resp.Body.Close()
		s, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(s))
	}

	//	TODO 加普通人好友

	for {
		//	这里面再去postman
	}

}

func makePersonToCelebrity() {

	//	就给这里面的明星加一些好友呗

}
