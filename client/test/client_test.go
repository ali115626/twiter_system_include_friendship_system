package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestAdd(t *testing.T){


	params := url.Values{}

	Url, err := url.Parse("http://127.0.0.1:9091/Show_TimeLines_Pull_Plus_Push?userId=10")
	if err != nil {
		panic(err.Error())

	}
	//params.Set("a", "fdfds")
	//params.Set("id", string("1"))
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	//defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(s))

}

