package inputs

//Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
import (
	"go-advent-of-code/file"
)
func Get(path string) []string {

	input:= file.Readfile(path)

	return input
	
}