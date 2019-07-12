package main

import "log"

func main() {

	inputText := readInput()

	conversionRoman, conversionIntergalactic, questions := CategorizeInput(inputText)

	romanMap, err := buildRomanMap(conversionRoman)
	if err != nil {
		log.Fatal(err)
	}

	objectMap, err := buildObjectMap(conversionIntergalactic, romanMap)
	if err != nil {
		log.Fatal(err)
	}

	result, err := buildResult(questions, romanMap, objectMap)
	if err != nil {
		log.Fatal(err)
	}

	printResult(result)
}
