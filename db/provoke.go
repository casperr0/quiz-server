package db

// CreateProvoke create a new provoke with correct and message.
func CreateProvoke(correct bool, message string) error {

	createSQL := `
	INSERT INTO provoke (correct, message)
	SELECT $1::BOOLEAN, $2::VARCHAR
	`
	tx := database.MustBegin()
	tx.MustExec(createSQL, correct, message)
	tx.Commit()
	return nil
}

// ListProvokes list all provokes.
func ListProvokes() ([]Provoke, error) {

	listSQL := `
	SELECT p.id, p.correct, p.message
	FROM provoke p
	`
	var provokes []Provoke
	err := database.Select(&provokes, listSQL)
	if err != nil {
		return nil, err
	}
	return provokes, nil
}

// QueryProvokes query all provokes with the correctness.
func QueryProvokes(correct bool) ([]Provoke, error) {

	querySQL := `
	SELECT p.id, p.correct, p.message
	FROM provoke p
	WHERE p.correct = $1
	`
	var provokes []Provoke
	err := database.Select(&provokes, querySQL, correct)
	if err != nil {
		return nil, err
	}
	return provokes, nil
}
