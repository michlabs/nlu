package nlu

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type IntentUtterance struct {
	Intent string `json:"intent"`
	Utterance string `json:"utterance"`
}

func ReadIntentsFromFile(inputFP string) ([]IntentUtterance, error) {
	input, err := os.Open(inputFP)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var ius []IntentUtterance
	scanner := bufio.NewScanner(input)
	counter := 0
	for scanner.Scan() {
		counter += 1
		var iu IntentUtterance
		text := strings.Replace(scanner.Text(), `"`, ``, -1) // Remove double quotes if have
		tokens := strings.SplitN(text, ",", 2)
		if len(tokens) < 2 {
			log.Println("Skip line %d: it is not complete: %s", counter, scanner.Text())
			continue
		}
		iu.Intent, iu.Utterance = strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1])
		ius = append(ius, iu)
	}

	return ius, nil
}