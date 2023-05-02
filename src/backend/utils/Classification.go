package utils

import (
	"backend/db"
	"backend/models"
	"regexp"
)

func CheckDate(input string) (bool, string) {

	pattern := `\d{2}/\d{2}/\d{4}`
	re := regexp.MustCompile(pattern)
	matches := re.FindString(input)

	if len(matches) == 0 {
		return false, ""
	}

	return true, matches
}

func CheckCalculate(input string) (bool, string) {
	pattern := `^[\d*/+\-^]*$`

	re := regexp.MustCompile(pattern)
	matches := re.MatchString(input)

	if matches {
		return matches, input
	}

	return false, ""
}

func CheckInsertQuesAns(input string) (isInsert, isQuesSame bool, matches1, matches2 string) {
	pattern := "Tambahkan pertanyan (.+) dengan jawaban (.+)"

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	if len(matches) == 0 {
		return false, false, "", ""
	}

	isInsert, matches1, matches2 = true, matches[1], matches[2]

	_db, err := db.GetDatabase()
	if err != nil {
		return
	}

	var question models.QuestAns

	if _ = _db.Where("question = ?", matches1).First(&question); len(question.Question) != 0 {
		isQuesSame = true
	}
	isQuesSame = false

	return
}

func CheckEraseQues(input string) (isAsk, isFound bool) {
	pattern := "Hapus pertanyan (.+)"

	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)

	if len(matches) != 0 {
		isAsk = true
	} else {
		isAsk = false
	}

	_db, err := db.GetDatabase()
	if err != nil {
		return
	}

	var question models.QuestAns

	if _ = _db.Where("question = ?", matches[1]).First(&question); len(question.Question) != 0 {
		_ = _db.Delete(question)
		isFound = true
	}

	isFound = false
	return
}