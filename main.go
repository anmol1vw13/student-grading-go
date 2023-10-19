package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func convertScore(score string) int{
	val, err := strconv.Atoi(score)
	if err != nil {
		return 0
	}
	return val
}

func parseCSV(filePath string) []student {

	file, err := os.Open(filePath)
	check(err)
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	headerLine := fileScanner.Text()

	header := strings.Split(headerLine, ",")
	headerIndexMap := make(map[string]int)
	for index, val := range header {
		headerIndexMap[val] = index
	}

	students := make([]student,0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		lineSplit := strings.Split(line, ",")
		student := student {
			firstName: lineSplit[headerIndexMap["FirstName"]],
			lastName: lineSplit[headerIndexMap["LastName"]],
			university: lineSplit[headerIndexMap["University"]],
			test1Score: convertScore(lineSplit[headerIndexMap["Test1"]]),
			test2Score: convertScore(lineSplit[headerIndexMap["Test2"]]),
			test3Score: convertScore(lineSplit[headerIndexMap["Test3"]]),
			test4Score: convertScore(lineSplit[headerIndexMap["Test4"]]),
		}
		students = append(students, student)

	}
	return students
}

func grade(score float32) Grade {
	if score >= 70 {
		return A
	} else if score >= 50 && score < 70 {
		return B
	} else if score >=35 && score < 50 {
		return C
	} 
	return F
	
}


func (s student) calculateScoreAndGrade() (float32, Grade){
	score := float32(s.test1Score + s.test2Score + s.test3Score + s.test4Score)/4
	return score, grade(score)
}

func calculateGrade(students []student) []studentStat {
	stats := make([]studentStat, 0)
	for _, student := range students {
		score, grade := student.calculateScoreAndGrade()
		stats = append(stats, studentStat{student: student, finalScore: score, grade: grade})
	}
	return stats
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	highestScorer := gradedStudents[0]
	for _, stat := range gradedStudents[1:] {
		if stat.finalScore > highestScorer.finalScore {
			highestScorer = stat
		}
	}
	return highestScorer
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	
	statMap := make(map[string]studentStat)
	for _, stat:= range gs {
		highestScorer, ok := statMap[stat.student.university]
		if !ok {
			statMap[stat.student.university] = stat
		} else if stat.finalScore > highestScorer.finalScore {
			statMap[stat.student.university] = stat
		}
	}
	return statMap
}
