package main

import (
	"log"
	"strings"
)

func loadBuddylist() (Buddylist, []string, []string) {

	rawList := fileToArray(buddyListPath)

	var NamesList, ServerList []string

	for _, s := range rawList {
		if !strings.Contains(s, "::") {
			log.Println("Didn't contain :: ", s)
			break
		} else if !strings.Contains(s, "@") {
			log.Println("Didn't contain @ ", s)
			break
		}
		newBuddy := Buddy{}

		fullNameAndPubkey := strings.Split(s, "::")
		nameAndServer := strings.Split(fullNameAndPubkey[0], "@")

		newBuddy.Pubkey = fullNameAndPubkey[1]
		newBuddy.Server = nameAndServer[1]
		newBuddy.Username = nameAndServer[0]

		NamesList = append(NamesList, nameAndServer[0])

	}

	return GlobalBuddyList, NamesList, ServerList
}

func loadRequestlist() (Buddylist, []string, []string) {

	rawList := fileToArray(requestListPath)

	var NamesList, ServerList []string

	for _, s := range rawList {
		newBuddy := Buddy{}

		fullNameAndPubkey := strings.Split(s, "::")
		nameAndServer := strings.Split(fullNameAndPubkey[0], "@")

		newBuddy.Pubkey = fullNameAndPubkey[1]
		newBuddy.Server = nameAndServer[1]
		newBuddy.Username = nameAndServer[0]

		NamesList = append(NamesList, nameAndServer[0])

	}

	return GlobalRequestList, NamesList, ServerList
}

func addFriendFromEntry(FullEntry string) bool {
	if !strings.Contains(FullEntry, "::") {
		log.Println("Didn't contain :: ", FullEntry)
		return false
	} else if !strings.Contains(FullEntry, "@") {
		log.Println("Didn't contain @ ", FullEntry)
		return false
	} else if len(FullEntry) < 78 {
		log.Println("Suspiciously short", FullEntry)
		return false
	}

	newBuddy := Buddy{}

	fullNameAndPubkey := strings.Split(FullEntry, "::")
	nameAndServer := strings.Split(fullNameAndPubkey[0], "@")

	newBuddy.Pubkey = fullNameAndPubkey[1]
	newBuddy.Server = nameAndServer[1]
	newBuddy.Username = nameAndServer[0]

	GlobalBuddyList.Buddys = append(GlobalBuddyList.Buddys, newBuddy)
	writeFile(buddyListPath, FullEntry)

	return true
}

func addRequestFromEntry(FullEntry string) bool {

	if !strings.Contains(FullEntry, "@") {
		log.Println("Didn't contain @ ", FullEntry)
		return false
	} else if len(FullEntry) < 10 {
		log.Println("Suspiciously short", FullEntry)
		return false
	} else if FullEntry == "" {
		log.Println("Entry is blank", FullEntry)
		return false
	}

	newBuddy := Buddy{}

	nameAndServer := strings.Split(FullEntry, "@")

	newBuddy.Server = nameAndServer[1]
	newBuddy.Username = nameAndServer[0]

	GlobalRequestList.Buddys = append(GlobalRequestList.Buddys, newBuddy)
	writeFile(requestListPath, FullEntry)

	return true
}
