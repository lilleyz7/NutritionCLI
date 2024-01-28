package data

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/lilleyz7/nutritionCLI/models"
)

func AddEntry(newFact models.NutritionFacts) error {
	data := []string{
		newFact.Name,
		fmt.Sprintf("%d", int(newFact.Sugar)),
		fmt.Sprintf("%d", int(newFact.Fiber)),
		fmt.Sprintf("%d", int(newFact.Sodium)),
		fmt.Sprintf("%d", int(newFact.Potassium)),
		fmt.Sprintf("%d", int(newFact.TotalFat)),
		fmt.Sprintf("%d", int(newFact.TotalSatFat)),
		fmt.Sprintf("%d", int(newFact.Calories)),
		fmt.Sprintf("%d", int(newFact.Cholestoral)),
		fmt.Sprintf("%d", int(newFact.Protein)),
		fmt.Sprintf("%d", int(newFact.Carbs)),
	}

	file, err := os.OpenFile("./data.csv", os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {

		err = CreateCSV()
		if err != nil {
			fmt.Println(err)
		}
		file, err = os.Open("./data.csv")
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println("How did we get here")

	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()
	err = writer.Write(data)
	if err != nil {
		println(err)
	}

	writer.Flush()
	return nil
}
