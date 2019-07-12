package main

import (
	"fmt"
	"strconv"
	"strings"
)

func buildRomanMap(conversionRoman []string) (map[string]string, error) {
	romanMap := make(map[string]string)
	for _, roman := range conversionRoman {
		temp := strings.Split(roman, " ")
		romanMap[temp[0]] = temp[2]
	}
	for _, tmp := range romanMap {
		found := false
		for romanChar := range RomanNumerals {
			if tmp == romanChar {
				found = true
			}
		}
		if !found {
			return romanMap, fmt.Errorf("%s is not a roman character", tmp)
		}
	}
	return romanMap, nil
}

func buildObjectMap(conversionIntergalactic []string, romanMap map[string]string) (map[string]int, error) {
	objectMap := make(map[string]int)
	for _, intergalactic := range conversionIntergalactic {
		romanValues := ""
		temp := strings.Split(intergalactic, " ")
		for i, tmp := range temp {
			if x, ok := romanMap[tmp]; ok {
				romanValues += x
			}
			if tmp == "is" {
				rom, err := RomanToInt(romanValues)
				if err != nil {
					return objectMap, err
				}
				if len(temp) == i {
					return objectMap, fmt.Errorf("invalid format")
				}
				val, err := strconv.Atoi(temp[i+1])
				if err != nil {
					return objectMap, err
				}
				objectMap[temp[i-1]] = val / rom
				continue
			}
		}
	}
	return objectMap, nil
}

func buildResult(questions []string, romanMap map[string]string, objectMap map[string]int) ([]string, error) {
	result := make([]string, 0)
	for _, question := range questions {
		var tempResult string
		foundRomanValues := false
		temp := strings.Split(question, " ")
		romanValues := ""
		for i, tmp := range temp {
			if x, ok := romanMap[tmp]; ok {
				tempResult += tmp + " "
				romanValues += x
				foundRomanValues = true
			}
			if x, ok := objectMap[tmp]; ok {
				rom, err := RomanToInt(romanValues)
				if err != nil {
					return result, err
				}
				tempResult += fmt.Sprintf("%s is %d Credits", tmp, x*rom)
				break
			}
			if i == len(temp)-1 {
				if foundRomanValues {
					rom, err := RomanToInt(romanValues)
					if err != nil {
						return result, err
					}
					tempResult += fmt.Sprintf("is %d", rom)
				} else {
					tempResult = "I have no idea what you are talking about"
				}
			}
		}
		result = append(result, tempResult)
	}
	return result, nil
}
