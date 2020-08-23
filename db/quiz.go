package db

import (
	"fmt"
)

// CreateQuiz create a new quiz with quiz structure.
func CreateQuiz(quiz Quiz) error {

	quizFound, _ := GetQuiz(quiz.Number)
	if quizFound != nil {
		tpl := "quiz number %d already exists"
		return fmt.Errorf(tpl, quiz.Number)
	}

	createSQL := `
	INSERT INTO quiz (number, description, score, option_a, option_b, option_c, option_d, answer)
	SELECT $1::INT, $2::VARCHAR, $3::INT, $4::VARCHAR, $5::VARCHAR, $6::VARCHAR, $7::VARCHAR, $8::VARCHAR
	`
	tx := database.MustBegin()
	tx.MustExec(
		createSQL,
		quiz.Number,
		quiz.Description,
		quiz.Score,
		quiz.OptionA,
		quiz.OptionB,
		quiz.OptionC,
		quiz.OptionD,
		quiz.Answer,
	)
	tx.Commit()
	return nil
}

// GetQuiz get a quiz with specified Number
func GetQuiz(quizNumber int) (*Quiz, error) {

	getSQL := "SELECT * FROM quiz WHERE number=$1"
	quiz := Quiz{}
	err := database.Get(&quiz, getSQL, quizNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "quiz number %d not found"
			return nil, fmt.Errorf(tpl, quizNumber)
		}
		return nil, err
	}
	return &quiz, nil
}

// ListQuizzes will list all current quizes.
func ListQuizzes() ([]Quiz, error) {

	listSQL := "SELECT * FROM quiz"
	var quizzes []Quiz
	err := database.Select(&quizzes, listSQL)
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

// QueryQuizzes query all quizes with the tag.
func QueryQuizzes(tagName string) ([]Quiz, error) {

	querySQL := `
	SELECT q.id, q.number, q.description, q.score, q.option_a, q.option_b, q.option_c, q.option_d, q.answer
	FROM   quiz q
	JOIN quiz_to_tag q_t ON q.id = q_t.quiz_id
	JOIN tag t ON t.id = q_t.tag_id
    WHERE t.name = $1
	`
	var quizzes []Quiz
	err := database.Select(&quizzes, querySQL, tagName)
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

// DeleteQuiz will delete quiz with specified ID.
func DeleteQuiz(quizNumber int) {

	deleteSQL := "DELETE FROM quiz WHERE number=$1"
	tx := database.MustBegin()
	tx.MustExec(deleteSQL, quizNumber)
	tx.Commit()
}

// CreateTag create a new quiz tag.
func CreateTag(tagName string) error {

	createSQL := `
	INSERT INTO tag (name)
	SELECT $1::VARCHAR
	WHERE NOT EXISTS (
		SELECT * FROM tag
		WHERE tag.name = $1
	);
	`
	tx := database.MustBegin()
	tx.MustExec(createSQL, tagName)
	tx.Commit()
	return nil
}

// GetTag get the tag with specified ID.
func GetTag(tagName string) (*Tag, error) {

	getSQL := "SELECT * FROM tag WHERE name=$1"
	tag := Tag{}
	err := database.Get(&tag, getSQL, tagName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "tag %s not found"
			return nil, fmt.Errorf(tpl, tagName)
		}
		return nil, err
	}
	return &tag, nil
}

// ListTags will list all current tags.
func ListTags() ([]Tag, error) {

	listSQL := "SELECT * FROM tag"
	var tags []Tag
	err := database.Select(&tags, listSQL)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// QueryTags query all tags of an quiz.
func QueryTags(quizNumber int) ([]Tag, error) {

	querySQL := `
	SELECT t.id, t.name
	FROM   tag t
	JOIN quiz_to_tag q_t ON t.id = q_t.tag_id
	JOIN quiz q ON q.id = q_t.quiz_id
    WHERE q.number = $1
	`
	var tags []Tag
	err := database.Select(&tags, querySQL, quizNumber)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// DeleteTag will delete Tag with specified name.
func DeleteTag(tagName string) {

	deleteSQL := "DELETE FROM tag WHERE name=$1"
	tx := database.MustBegin()
	tx.MustExec(deleteSQL, tagName)
	tx.Commit()
}

// RegisterTag will mark a quiz with an tag.
func RegisterTag(quizNumber int, tagName string) error {

	registerSQL := `
	INSERT INTO quiz_to_tag (quiz_id, tag_id)
	SELECT $1::INT, $2::INT
	WHERE NOT EXISTS(
		SELECT * FROM quiz_to_tag
		WHERE quiz_id = $1 AND tag_id = $2
	);
	`
	quizFound, err := GetQuiz(quizNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "quiz number %s not found"
			return fmt.Errorf(tpl, quizNumber)
		}
		return err
	}
	quizID := quizFound.ID
	tagFound, err := GetTag(tagName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "tag %s not found"
			return fmt.Errorf(tpl, tagName)
		}
		return err
	}
	tagID := tagFound.ID
	tx := database.MustBegin()
	tx.MustExec(registerSQL, quizID, tagID)
	tx.Commit()
	return nil
}

// DeregisterTag will deregister a tag from a quiz.
func DeregisterTag(quizNumber int, tagName string) error {

	deregisterSQL := `
	DELETE FROM quiz_to_tag
	WHERE quiz_id = $1 AND tag_id = $2;
	`
	quizFound, err := GetQuiz(quizNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "quiz number %s not found"
			return fmt.Errorf(tpl, quizNumber)
		}
		return err
	}
	quizID := quizFound.ID
	tagFound, err := GetTag(tagName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "tag %s not found"
			return fmt.Errorf(tpl, tagName)
		}
		return err
	}
	tagID := tagFound.ID
	tx := database.MustBegin()
	tx.MustExec(deregisterSQL, quizID, tagID)
	tx.Commit()
	return nil
}
