package quotes

import (
	"math/rand"
	"strings"
)

func Get() string {

	listOfQuotes := getListOfQuotes()

	quoteNum := rand.Intn(len(listOfQuotes))

	return listOfQuotes[quoteNum]

}

func FromAuthor(author string) string {

	listOfQuotes := getListOfQuotes()

	quotesFromAuthor := getQuotesFromAuthor(author, listOfQuotes)

	return FromList(quotesFromAuthor)

}

func FromList(listOfQuotes []string) string {

	quoteNum := rand.Intn(len(listOfQuotes))

	return listOfQuotes[quoteNum]
}

func getQuotesFromAuthor(author string, listOfQuotes []string) []string {

	response := make([]string, len(listOfQuotes)/3)

	for _, quote := range listOfQuotes {

		lowerCaseQuote := strings.ToLower(quote)
		lowerCaseAuthor := strings.ToLower(author)

		if strings.Contains(lowerCaseQuote, lowerCaseAuthor) {

			response = append(response, quote)
		}

	}

	return response
}
