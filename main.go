package main

import "log"

func main() {

	inputText := ReadInput()

	intergalacticToRomanInputArray, metalValueInputArray, questions := CategorizeInput(inputText)

	intergalacticToRomanMap, err := BuildIntergalacticTointergalacticToRomanMap(intergalacticToRomanInputArray)
	if err != nil {
		log.Fatal(err)
	}

	metalValueMap, err := BuildMetalValueMap(metalValueInputArray, intergalacticToRomanMap)
	if err != nil {
		log.Fatal(err)
	}

	result, err := BuildResult(questions, intergalacticToRomanMap, metalValueMap)
	if err != nil {
		log.Fatal(err)
	}

	PrintResult(result)
}
