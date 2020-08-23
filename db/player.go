package db

import (
	"fmt"

	// posetgreSQL databse driver required by sqlx
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PlayerScore describe the player with its total score.
type PlayerScore struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Score int    `db:"score"`
}

// CreatePlayer will create a new player by name.
func CreatePlayer(playerName string) error {

	player, _ := GetPlayer(playerName)
	if player != nil {
		tpl := "player %s already existed"
		return fmt.Errorf(tpl, playerName)
	}

	createSQL := `
	INSERT INTO player (name)
	SELECT $1::VARCHAR
	WHERE NOT EXISTS (
		SELECT 1 FROM player
		WHERE player.name = $1
	);
	`
	tx := database.MustBegin()
	tx.MustExec(createSQL, playerName)
	tx.Commit()
	return nil
}

// GetPlayer will get the player with specified playerName.
func GetPlayer(playerName string) (*Player, error) {

	getSQL := "SELECT * FROM player WHERE name=$1"
	player := Player{}
	err := database.Get(&player, getSQL, playerName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "player %s not found"
			return nil, fmt.Errorf(tpl, playerName)
		}
		return nil, err
	}
	return &player, nil
}

// RegisterAnswer will records whether the player's answer is correct or not.
func RegisterAnswer(playerName string, quizNumber int, correct bool) error {

	registerSQL := `
	INSERT INTO player_to_quiz (player_id, quiz_id, correct)
	SELECT $1::INT, $2::INT, $3::BOOLEAN
	WHERE NOT EXISTS (
		SELECT 1 FROM player_to_quiz
		WHERE player_to_quiz.player_id = $1 AND player_to_quiz.quiz_id = $2
	);
	`
	playerFound, err := GetPlayer(playerName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "player %s not found"
			return fmt.Errorf(tpl, playerName)
		}
		return err
	}
	playerID := playerFound.ID
	quizFound, err := GetQuiz(quizNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "quiz number %d not found"
			return fmt.Errorf(tpl, quizNumber)
		}
		return err
	}
	quizID := quizFound.ID
	tx := database.MustBegin()
	tx.MustExec(registerSQL, playerID, quizID, correct)
	tx.Commit()
	return nil
}

// ListPlayersOrderByScore will list all current players order by score.
func ListPlayersOrderByScore() ([]PlayerScore, error) {

	listSQL := `
	SELECT p.id, p.name, SUM(q.score) AS score
	FROM player p
	JOIN player_to_quiz p_q ON p.id = p_q.player_id
	JOIN quiz q ON q.id = p_q.quiz_id
	WHERE p_q.correct = TRUE
	GROUP BY p.id
	ORDER BY SUM(q.score) DESC
	`
	var playerScores []PlayerScore
	err := database.Select(&playerScores, listSQL)
	if err != nil {
		return nil, err
	}
	return playerScores, nil
}

// DeletePlayer will delete player with specified name.
func DeletePlayer(playerName string) {

	deleteSQL := "DELETE FROM player WHERE name=$1"
	tx := database.MustBegin()
	tx.MustExec(deleteSQL, playerName)
	tx.Commit()
}

// GetAnswer will get answer by playerName and quizNumber.
func GetAnswer(playerName string, quizNumber int) (*PlayerToQuiz, error) {

	getSQL := `
	SELECT p_q.player_id, p_q.quiz_id, p_q.correct
	FROM player_to_quiz p_q
	WHERE p_q.player_id = $1 and p_q.quiz_id = $2
	`
	playerFound, err := GetPlayer(playerName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "player %s not found"
			return nil, fmt.Errorf(tpl, playerName)
		}
		return nil, err
	}
	playerID := playerFound.ID
	quizFound, err := GetQuiz(quizNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "quiz number %d not found"
			return nil, fmt.Errorf(tpl, quizNumber)
		}
		return nil, err
	}
	quizID := quizFound.ID
	var answer PlayerToQuiz
	err = database.Get(&answer, getSQL, playerID, quizID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, fmt.Errorf("answer not found")
		}
		return nil, err
	}
	return &answer, nil
}

// ListAnswers will list all answers from all players to all quizzes.
func ListAnswers() ([]PlayerToQuiz, error) {

	listSQL := `
	SELECT p_q.player_id, p_q.quiz_id, p_q.correct
	FROM player_to_quiz p_q
	`
	var answers []PlayerToQuiz
	err := database.Select(&answers, listSQL)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

// QueryPlayerAnswersByQuiz query all players have answered the specified quiz.
func QueryPlayerAnswersByQuiz(quizNumber int) ([]PlayerToQuiz, error) {

	querySQL := `
	SELECT p_q.player_id, p_q.correct
	FROM player_to_quiz p_q
	WHERE p_q.quiz_id = $1
	`
	quizFound, err := GetQuiz(quizNumber)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "quiz number %d not found"
			return nil, fmt.Errorf(tpl, quizNumber)
		}
		return nil, err
	}
	quizID := quizFound.ID

	var playerToQuizes []PlayerToQuiz
	err = database.Select(&playerToQuizes, querySQL, quizID)
	if err != nil {
		return nil, err
	}
	return playerToQuizes, nil
}

// QueryQuizAnswersByPlayer query all quiz answered by specified player.
func QueryQuizAnswersByPlayer(playerName string) ([]PlayerToQuiz, error) {

	querySQL := `
	SELECT p_q.quiz_id, p_q.correct
	FROM player_to_quiz p_q
	WHERE p_q.player_id = $1
	`
	playerFound, err := GetPlayer(playerName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "player %s not found"
			return nil, fmt.Errorf(tpl, playerName)
		}
		return nil, err
	}
	playerID := playerFound.ID

	var playerToQuizes []PlayerToQuiz
	err = database.Select(&playerToQuizes, querySQL, playerID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "player %s have not answer any quiz."
			return nil, fmt.Errorf(tpl, playerName)
		}
		return nil, err
	}
	return playerToQuizes, nil
}
