package main

import "fmt"

type Student struct {
  Name         string
  MathScore    int
  ScienceScore int
  EnglishScore int
}

func main() {
  students := []Student{
      {"jamie", 67, 88, 90},
      {"michael", 77, 77, 90},
      {"aziz", 100, 65, 88},
      {"jamal", 50, 80, 75},
      {"eric", 70, 80, 65},
      {"john", 80, 76, 68},
  }

  getInfo(students)
  // output:
  // best student in math: aziz with 100
  // best student in science: jamie with 88
  // best student in english: jamie with 90
  // best student in average: aziz with 84

}

func getInfo(students []Student) {
	bestMath, bestScience, bestEnglish, bestAverage := findBest(students, 0, students[0], students[0], students[0], students[0])
	fmt.Printf("best student in math: %s with %d\n", bestMath.Name, bestMath.MathScore)
	fmt.Printf("best student in science: %s with %d\n", bestScience.Name, bestScience.ScienceScore)
	fmt.Printf("best student in english: %s with %d\n", bestEnglish.Name, bestEnglish.EnglishScore)
	averageScore := (bestAverage.MathScore + bestAverage.ScienceScore + bestAverage.EnglishScore) / 3
	fmt.Printf("best student in average: %s with %d\n", bestAverage.Name, averageScore)
}

func findBest(students []Student, index int, bestMath, bestScience, bestEnglish, bestAverage Student) (Student, Student, Student, Student) {
	if index == len(students) {
		return bestMath, bestScience, bestEnglish, bestAverage
	}
	current := students[index]
	averageScore := (current.MathScore + current.ScienceScore + current.EnglishScore) / 3
	bestAverageScore := (bestAverage.MathScore + bestAverage.ScienceScore + bestAverage.EnglishScore) / 3

	if current.MathScore > bestMath.MathScore {
		bestMath = current
	}
	if current.ScienceScore > bestScience.ScienceScore {
		bestScience = current
	}
	if current.EnglishScore > bestEnglish.EnglishScore {
		bestEnglish = current
	}
	if averageScore > bestAverageScore {
		bestAverage = current
	}

	return findBest(students, index+1, bestMath, bestScience, bestEnglish, bestAverage)
}