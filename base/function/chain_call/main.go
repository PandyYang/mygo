package main

import (
	"fmt"
	"mygo/base/function/chain_call/friend"
)

func main() {

	findFriends, err := friend.Find("附近的人",
		friend.WithSex(1),
		friend.WithHeight(165),
		friend.WithWeight(50),
		friend.WithAge(27),
		friend.WithHobby("reading"))

	if err != nil {
		fmt.Println(err)
	}

	println(findFriends)
}
