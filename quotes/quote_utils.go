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

	if len(quotesFromAuthor) < 1 {
		return "No quotes from this author in the database"
	}

	return getRandomQuoteFromList(quotesFromAuthor)
}

func getRandomQuoteFromList(listOfQuotes []string) string {
	if len(listOfQuotes) == 1 {
		return listOfQuotes[0]
	}

	quoteNum := rand.Intn(len(listOfQuotes))

	return listOfQuotes[quoteNum]
}

func getQuotesFromAuthor(author string, listOfQuotes []string) []string {

	response := []string{}

	for _, quote := range listOfQuotes {

		lowerCaseQuote := strings.ToLower(quote)
		lowerCaseAuthor := strings.ToLower(author)

		if strings.Contains(lowerCaseQuote, lowerCaseAuthor) {

			response = append(response, quote)
		}

	}

	return response
}
