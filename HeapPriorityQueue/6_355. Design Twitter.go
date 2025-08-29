package main

type Twitter struct {
	// userId:userId:bool
	relations map[int]map[int]bool
	// userId:postId
	feeds      map[int]*[]HeapData
	idIterator int
}

func ConstructorTwitter() Twitter {
	return Twitter{
		relations:  map[int]map[int]bool{},
		feeds:      map[int]*[]HeapData{},
		idIterator: 0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	tweetsHeap, exists := this.feeds[userId]
	if !exists {
		this.feeds[userId] = &[]HeapData{}
	}
	MinHeapInsert(this.feeds[userId], HeapData{key: float32(this.idIterator), value: tweetId})
	this.idIterator++
	if len(*this.feeds[userId]) > 10 {
		MinHeapPop(tweetsHeap)
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	tempHeap := []HeapData{}
	total := []int{userId}

	for folowingId := range this.relations[userId] {
		total = append(total, folowingId)
	}

	for _, folowingId := range total {
		currentFeedHeap, exists := this.feeds[folowingId]
		if exists {
			for i := 0; i < len(*currentFeedHeap); i++ {
				MinHeapInsert(&tempHeap, (*currentFeedHeap)[i])
				if len(tempHeap) > 10 {
					MinHeapPop(&tempHeap)
				}
			}
		}
	}

	MinHeapSort(&tempHeap)

	result := []int{}
	for _, num := range tempHeap {
		result = append(result, num.value.(int))
	}

	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	_, exists := this.relations[followerId]
	if !exists {
		this.relations[followerId] = map[int]bool{}
	}
	this.relations[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.relations[followerId], followeeId)
}

func runCommandsTwitter(commands []string, args [][]int) []any {
	var obj Twitter
	results := []any{}
	for i := range commands {
		args := args[i]
		switch commands[i] {
		case "Twitter":
			obj = ConstructorTwitter()
		case "postTweet":
			obj.PostTweet(args[0], args[1])
			results = append(results, nil)
		case "follow":
			obj.Follow(args[0], args[1])
			results = append(results, nil)
		case "getNewsFeed":
			results = append(results, obj.GetNewsFeed(args[0]))
		case "unfollow":
			obj.Unfollow(args[0], args[1])
			results = append(results, nil)
		}
	}
	return results
}

// func main() {
// 	fmt.Println(
// 		runCommandsTwitter(
// 			[]string{"Twitter", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "getNewsFeed", "follow", "getNewsFeed", "unfollow", "getNewsFeed"},
// 			[][]int{{}, {1, 5}, {2, 3}, {1, 101}, {2, 13}, {2, 10}, {1, 2}, {1, 94}, {2, 505}, {1, 333}, {2, 22}, {1, 11}, {1, 205}, {2, 203}, {1, 201}, {2, 213}, {1, 200}, {2, 202}, {1, 204}, {2, 208}, {2, 233}, {1, 222}, {2, 211}, {1}, {1, 2}, {1}, {1, 2}, {1}},
// 		),
// 	)
// 	// fmt.Println(
// 	// 	runCommands(
// 	// 		[]string{"Twitter", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "postTweet", "getNewsFeed", "follow", "getNewsFeed", "unfollow", "getNewsFeed"},
// 	// 		[][]int{{}, {1, 5}, {2, 3}, {1, 101}, {2, 13}, {2, 10}, {1, 2}, {1, 94}, {2, 505}, {1, 333}, {2, 22}, {1, 11}, {1, 205}, {2, 203}, {1, 201}, {2, 213}, {1, 200}, {2, 202}, {1, 204}, {2, 208}, {2, 233}, {1, 222}, {2, 211}, {1}, {1, 2}, {1}, {1, 2}, {1}},
// 	// 	),
// 	// )
// }
