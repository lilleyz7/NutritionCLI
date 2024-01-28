package main

import (
	"github.com/lilleyz7/nutritionCLI/initializers"
	"github.com/lilleyz7/nutritionCLI/repl"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	repl.RunREPL()

}
