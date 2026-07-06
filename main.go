package main

import "fmt"

func main() {
	 scores := []int{50, 75, 66, 20, 32, 90}
	 
	 targetIdx := -1

	 for i, v := range scores {
		if v == 66 {
			targetIdx = i
			break
		}
	 }

	 if targetIdx != -1 {
	leftPart := append([]int{}, scores[:targetIdx+1]...)
	leftPart = append(leftPart, 88)

	scores = append(leftPart,scores[targetIdx+1:]... )
	 }

	 fmt.Println("Score setelah di sisipkan 88")
	 for _, score := range scores {
		fmt.Println(score)
	 }
}