package quotes

import "math/rand"

func Get() string {

	listOfQuotes := getListOfQuotes()

	quoteNum := rand.Intn(len(listOfQuotes))

	return listOfQuotes[quoteNum]

}

func getListOfQuotes() []string {

	return []string{
		"You could be good today but you choose tomorrow - Marcus Aureilios",
		"Live, love, laugh - elif",
	}

}
