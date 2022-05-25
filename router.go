package main

import (
	"net/http"
	"twiter/tweetSystem/Service"
	"twiter/tweetSystem/Service/Pull_Plus_Push"
	"twiter/tweetSystem/Service/pull"
	"twiter/tweetSystem/Service/split_and_conquer"
)


func init(){



}

func main(){

	//今天把所有pull和push的接口写完
	//ShowNewsFeed
	http.HandleFunc("/show_news_feed", Service.ShowNewsFeed)
	//ShowTimeLine

	http.HandleFunc("/show_time_line", Service.ShowTimeLine)

	//PostTweetPull

	http.HandleFunc("/post_tweet_push", Service.PostTweetPush)

	//PostTweetFanOutTest

	http.HandleFunc("/post_tweetFanOutTest", Service.PostTweetFanOutTest)

	http.HandleFunc("/PostTweetFanOutSplit", split_and_conquer.PostTweetFanOutSplit)


	//todo push的方法的改进
	//todo Pull_Plus_Push.ShowTimeLinesPullPlusPush
	http.HandleFunc("/Show_TimeLines_Pull_Plus_Push", Pull_Plus_Push.ShowTimeLinesPullPlusPush)
	//



	//SearchTweetIdListForCommonPersons

	//_______________ -_下面是pull的方式_______________________________________________
	http.HandleFunc("/PostTweetPull", pull.PostTweetPull)
	//ShowNewsFeedPull

	http.HandleFunc("/ShowNewsFeedPull", pull.ShowNewsFeedPull)
	//show myself的这个  你自己发了啥 你自己心里就没点数吗？

	//todo  逻辑差不多  但是就是不用去找friends了  也不用去迭代friendsList    直接userId=userIdInput
	http.HandleFunc("/ShowMyselfContentPull", pull.ShowMyselfContentPull)
	//PostTweetFanOutSplit
	//MysqlInit()

	//PostTweetAddCelebrity



	//----------------    ---------------



	http.HandleFunc("/post_tweet_add_celebrity", Service.PostTweetAddCelebrity)




	err := http.ListenAndServe(":9091", nil) // 设置监听的端口
	if err != nil {
		//    logger.Fatal("ListenAndServe: ", err)
	}



}
