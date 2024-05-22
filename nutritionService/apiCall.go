package nutritionservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/lilleyz7/nutritionCLI/models"
)

func GetData(food string) (models.NutritionFacts, error) {
	baseUrl := "https://nutrition-by-api-ninjas.p.rapidapi.com/v1/nutrition?query=%s"
	finalUrl := baseUrl + food

	var res []models.NutritionFacts
	err := ConvertJSON(finalUrl, &res)
	if err != nil {
		return models.NutritionFacts{}, nil
	} else {
		return res[0], nil
	}
}

func ConvertJSON(url string, target interface{}) error {
	key := os.Getenv("API_KEY")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-RapidAPI-Key", key)
	req.Header.Add("X-RapidAPI-Host", "nutrition-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(target)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
	return nil
}
