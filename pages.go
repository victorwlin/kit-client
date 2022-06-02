package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	friends := getFriends()

	tpl.ExecuteTemplate(res, "index.gohtml", friends.Friends)
}

func showFriend(res http.ResponseWriter, req *http.Request) {
	friendName := req.FormValue("friend")
	friend := getFriend(friendName)

	tpl.ExecuteTemplate(res, "showFriend.gohtml", friend.Friends[0])
}

func addFriend(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		friendName := req.FormValue("friendname")
		group := req.FormValue("group")
		desiredFreq := req.FormValue("desiredfreq")
		lastContact := req.FormValue("lastcontact")

		// Friend name cannot contain spaces.
		if strings.ContainsAny(friendName, " ") {
			http.Error(res, "Friend name cannot contain spaces.", http.StatusForbidden)
		}

		// check if fields have been filled out
		if friendName == "" || group == "" || desiredFreq == "" || lastContact == "" {
			http.Error(res, "All fields must be filled out.", http.StatusUnauthorized)
			return

		} else {

			// conversions
			desiredFreqInt, err := strconv.Atoi(desiredFreq)
			if err != nil {
				fmt.Println(err)
			}

			lastContactDate, err := time.Parse("2006-01-02", lastContact)
			if err != nil {
				fmt.Println(err)
			}
			lastContactString := lastContactDate.Format("2006-01-02")

			friend := friend{
				FriendName:  friendName,
				Group:       group,
				DesiredFreq: desiredFreqInt,
				LastContact: lastContactString,
			}

			response := addFriendAPI(friend)

			if response == 201 {
				http.Redirect(res, req, "/", http.StatusSeeOther)
			} else if response == 409 {
				http.Error(res, "Friend already exists.", http.StatusUnauthorized)
				return
			} else {
				fmt.Fprintf(res, strconv.Itoa(response))
				return
			}

		}

	}

	tpl.ExecuteTemplate(res, "addfriend.gohtml", nil)
}

func editFriend(res http.ResponseWriter, req *http.Request) {

	friendName := req.FormValue("friend")

	// get current friend details
	currentFriend := getFriend(friendName).Friends[0]

	if req.Method == http.MethodPost {

		newGroup := req.FormValue("newGroup")
		newDesiredFreq := req.FormValue("newDesiredFreq")
		newLastContact := req.FormValue("newLastContact")

		friend := friend{
			FriendName:  currentFriend.FriendName,
			Group:       currentFriend.Group,
			DesiredFreq: currentFriend.DesiredFreq,
			LastContact: currentFriend.LastContact,
		}

		if newGroup != "" {
			friend.Group = newGroup
		}

		if newDesiredFreq != "" {
			tempDesiredFreq, err := strconv.Atoi(newDesiredFreq)
			if err != nil {
				fmt.Println(err)
			} else {
				friend.DesiredFreq = tempDesiredFreq
			}
		}

		if newLastContact != "" {
			friend.LastContact = newLastContact
		}
		fmt.Println("pages sent: ", friend)
		response := editFriendAPI(friend)

		if response == 201 {
			http.Redirect(res, req, "/friend/?friend="+friend.FriendName, http.StatusSeeOther)
		} else {
			fmt.Fprintf(res, strconv.Itoa(response))
		}

	}

	tpl.ExecuteTemplate(res, "editfriend.gohtml", currentFriend)
}

func deleteFriend(res http.ResponseWriter, req *http.Request) {
	friendName := req.FormValue("friend")
	friend := getFriend(friendName).Friends[0]

	deleteFriendAPI(friend)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
