package handlers

import (
	"fmt"

	"github.com/lilleyz7/nutritionCLI/data"
	nutritionservice "github.com/lilleyz7/nutritionCLI/nutritionService"
)

func AddHandler() error {
	facts, err := nutritionservice.GetData(FoodTerm)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(facts)

	err = data.AddEntry(facts)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully added %s to you nutrition counter", facts.Name)
	return nil
}
