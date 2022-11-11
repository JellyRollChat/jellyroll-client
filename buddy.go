package main

import "strings"

func getBuddyList() (Buddylist, []string, []string) {

	rawList := fileToArray(buddyListPath)

	var NamesList []string
	var ServerList []string

	GlobalBuddyList = Buddylist{}

	if len(GlobalBuddyList.Buddys) < 1 {
		writeFile(buddyListPath, "donuthandler::3ck0.com::b19d495c2dcf9f92b0b420ba4d2a975d44df9b70014930f3b883b8ddbd253cc4")
		// getBuddyList()
	}

	for _, s := range rawList {
		splitList := strings.Split(s, "::")
		NamesList = append(NamesList, splitList[0])
		ServerList = append(ServerList, splitList[1])
		GlobalBuddyList.Buddys = append(GlobalBuddyList.Buddys, Buddy{Username: splitList[0], Server: splitList[1], Pubkey: splitList[2]})
	}

	return GlobalBuddyList, NamesList, ServerList
}
