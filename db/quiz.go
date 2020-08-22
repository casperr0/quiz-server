package db

import "fmt"

// CreateQuiz create a new quiz with quiz structure.
func CreateQuiz(quiz Quiz) error {

	createSQL := `
	INSERT INTO quiz (description, score, option_a, option_b, option_c, option_d, answer)
	SELECT $1::VARCHAR, $2::INT, $3::VARCHAR, $4::VARCHAR, $5::VARCHAR, $6::VARCHAR, $7::VARCHAR
	`
	tx := database.MustBegin()
	tx.MustExec(
		createSQL,
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

// GetQuiz get a quiz with specified ID
func GetQuiz(quizID int) (*Quiz, error) {

	getSQL := "SELECT * FROM quiz WHERE id=$1"
	quiz := Quiz{}
	err := database.Get(&quiz, getSQL, quizID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "quiz with id %d not found"
			return nil, fmt.Errorf(tpl, quizID)
		}
		return nil, err
	}
	return &quiz, nil
}

// ListQuizes will list all current quizes.
func ListQuizes() ([]Quiz, error) {

	listSQL := "SELECT * FROM quiz"
	var quizes []Quiz
	err := database.Select(&quizes, listSQL)
	if err != nil {
		return nil, err
	}
	return quizes, nil
}

// QueryQuizes query all quizes with the tag.
func QueryQuizes(tagName string) ([]Quiz, error) {

	querySQL := `
	SELECT q.id, q.description, q.score, q.option_a, q.option_b, q.option_c, q.option_d, q.answer
	FROM   quiz q
	JOIN quiz_to_tag q_t ON q.id = q_t.quiz_id
	JOIN tag t ON t.id = q_t.tag_id
    WHERE t.name = $1
	`
	var quizes []Quiz
	err := database.Select(&quizes, querySQL, tagName)
	if err != nil {
		return nil, err
	}
	return quizes, nil
}

// DeleteQuiz will delete quiz with specified ID.
func DeleteQuiz(quizID int) {

	deleteSQL := "DELETE FROM quiz WHERE id=$1"
	tx := database.MustBegin()
	tx.MustExec(deleteSQL, quizID)
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
func QueryTags(quizID int) ([]Tag, error) {

	querySQL := `
	SELECT t.id, t.name
	FROM   tag t
	JOIN quiz_to_tag q_t ON t.id = q_t.tag_id
	JOIN quiz q ON q.id = q_t.quiz_id
    WHERE q.id = $1
	`
	var tags []Tag
	err := database.Select(&tags, querySQL, quizID)
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
func RegisterTag(quizID int, tagName string) error {

	registerSQL := `
	INSERT INTO quiz_to_tag (quiz_id, tag_id)
	SELECT $1::INT, $2::INT
	WHERE NOT EXISTS(
		SELECT * FROM quiz_to_tag
		WHERE quiz_id = $1 AND tag_id = $2
	);
	`
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
func DeregisterTag(quizID int, tagName string) error {

	deregisterSQL := `
	DELETE FROM quiz_to_tag
	WHERE quiz_id = $1 AND tag_id = $2;
	`
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
