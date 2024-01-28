package handlers

import (
	"fmt"

	nutritionservice "github.com/lilleyz7/nutritionCLI/nutritionService"
)

func SearchHandler() error {
	facts, err := nutritionservice.GetData(FoodTerm)
	if err != nil {
		fmt.Println(err)
		return err
	}
	

	fmt.Printf("Name: %s \n", facts.Name)
	fmt.Printf("Sugar: %f \n", facts.Sugar)
	fmt.Printf("Fiber: %f \n", facts.Fiber)
	fmt.Printf("Sodium: %f \n", facts.Sodium)
	fmt.Printf("Potassium: %f \n", facts.Potassium)
	fmt.Printf("Total Fat: %f \n", facts.TotalFat)
	fmt.Printf("Total Saturated Fat: %f \n", facts.TotalSatFat)
	fmt.Printf("Calories: %f \n", facts.Calories)
	fmt.Printf("Cholestoral: %f \n", facts.Cholestoral)
	fmt.Printf("Protein: %f \n", facts.Protein)
	fmt.Printf("Carbs: %f \n", facts.Carbs)
	return nil
}
