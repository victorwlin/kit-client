package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type friend struct {
	FriendName  string
	Group       string
	DesiredFreq int
	LastContact string
	NextContact string
}

type friends struct {
	Friends []friend
}

const baseURL string = "https://kit-api.herokuapp.com/api/v1/friends"

var (
	envErr = godotenv.Load(".env")
	key    = os.Getenv("KEY")
)

func getFriends() (friends friends) {
	response, err := http.Get(baseURL + "?key=" + key)

	if err == nil {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		json.Unmarshal(data, &friends)

		response.Body.Close()
	} else {
		fmt.Println(err)
	}

	return friends
}

func getFriend(friendName string) (friends friends) {
	response, err := http.Get(baseURL + "/" + friendName + "?key=" + key)

	if err == nil {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		json.Unmarshal(data, &friends)

		response.Body.Close()
	} else {
		fmt.Println(err)
	}

	return friends
}

func addFriendAPI(friend friend) (res int) {
	byteFriend, err := json.Marshal(friend)
	if err != nil {
		fmt.Println(err)
	}
	jsonFriend := bytes.NewBuffer(byteFriend)

	response, err := http.Post(baseURL+"/"+friend.FriendName+"?key="+key, "application/json", jsonFriend)

	if err == nil {
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		res = response.StatusCode
		fmt.Println(string(data))

		response.Body.Close()
	} else {
		fmt.Println(err)
	}

	return res
}

func editFriendAPI(friend friend) (res int) {
	byteFriend, err := json.Marshal(friend)
	if err != nil {
		fmt.Println(err)
	}
	jsonFriend := bytes.NewBuffer(byteFriend)
	fmt.Println("apiCalls sent: ", jsonFriend)
	request, err := http.NewRequest(http.MethodPut, baseURL+"/"+friend.FriendName+"?key="+key, jsonFriend)
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err == nil {

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		res = response.StatusCode
		fmt.Println(string(data))

		response.Body.Close()

	} else {
		fmt.Println(err)
	}

	return res
}

func deleteFriendAPI(friend friend) (res int) {

	request, err := http.NewRequest(http.MethodDelete, baseURL+"/"+friend.FriendName+"?key="+key, nil)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err == nil {

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		res = response.StatusCode
		fmt.Println(string(data))

		response.Body.Close()

	} else {
		fmt.Println(err)
	}

	return res
}
