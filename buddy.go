package main

import "strings"

func getBuddyList() (Buddylist, []string) {
	rawList := fileToArray(buddyListPath)

	var NamesList []string

	GlobalBuddyList = Buddylist{}
	for _, s := range rawList {
		splitList := strings.Split(s, "::")
		NamesList = append(NamesList, splitList[0])
		GlobalBuddyList.Buddys = append(GlobalBuddyList.Buddys, Buddy{Username: splitList[0], Pubkey: splitList[1]})
	}

	return GlobalBuddyList, NamesList
}
