package main

var historics []string

func save(operation string, result string) {
	historics = append(historics, operation+result+". ")
}

func getHistory() []string {
	return historics
}
