package models

type NutritionFacts struct {
	Name        string  `json:"name" `
	Sugar       float32 `json:"sugar_g" `
	Fiber       float32 `json:"fiber_g" `
	Sodium      float32 `json:"sodium_mg" `
	Potassium   float32 `json:"potassium_mg" `
	TotalFat    float32 `json:"total_fat_g" `
	TotalSatFat float32 `json:"fat_saturated_g" `
	Calories    float32 `json:"calories" `
	Cholestoral float32 `json:"cholestoral_mg" `
	Protein     float32 `json:"protein_g"`
	Carbs       float32 `json:"carbohydrates_total_g"`
}
