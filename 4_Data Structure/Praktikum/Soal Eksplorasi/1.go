package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	records := readCSV("survey.csv")
	if records == nil {
		return
	}

	averageMentorScore, averageLearningScore, highestCourse, highestAverage := calculate(records)

	fmt.Printf("Average Mentor Satisfaction Score: %.2f\n", averageMentorScore)
	fmt.Printf("Average Learning Satisfaction Score: %.2f\n", averageLearningScore)
	fmt.Printf("Course with Highest Average Mentor Satisfaction: %s (%.2f)\n", highestCourse, highestAverage)
}

func readCSV(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return nil
	}
	return records
}

func calculate(records [][]string) (float64, float64, string, float64) {
	totalMentorScore := 0.0
	totalLearningScore := 0.0
	courseSatisfaction := make(map[string][]float64)

	for i, record := range records {
		if i == 0 {
			continue
		}
		learningScore := parseToFloat(record[3])
		mentorScore := parseToFloat(record[4])

		totalMentorScore += mentorScore
		totalLearningScore += learningScore

		courseName := record[2]
		courseSatisfaction[courseName] = append(courseSatisfaction[courseName], mentorScore)
	}

	numRecords := float64(len(records) - 1)
	averageMentorScore := totalMentorScore / numRecords
	averageLearningScore := totalLearningScore / numRecords

	highestCourse := ""
	highestAverage := 0.0

	for course, scores := range courseSatisfaction {
		sum := 0.0
		for _, score := range scores {
			sum += score
		}
		average := sum / float64(len(scores))
		if average > highestAverage {
			highestAverage = average
			highestCourse = course
		}
	}

	return averageMentorScore, averageLearningScore, highestCourse, highestAverage
}

func parseToFloat(s string) float64 {
	var result float64
	for i := 0; i < len(s); i++ {
		result = result*10 + float64(s[i]-'0')
	}
	return result
}
