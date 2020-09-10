package db

import (
	"fmt"
	"log"
	"os"

	// posetgreSQL databse driver required by sqlx
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
	// load config for environment setup
)

// Officer describe the schema of event staff.
type Officer struct {
	ID       int    `db:"id" json:"_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

// Role describe the schema of role with different permissions.
type Role struct {
	ID   int    `db:"id" json:"_id"`
	Name string `db:"name" json:"name"`
}

// Player describe the schema of quiz player.
type Player struct {
	ID       int    `db:"id" json:"_id"`
	Name     string `db:"name" json:"name"`
	Nickname string `db:"nickname" json:"nickname"`
	Platform string `db:"platform" json:"platform"`
}

// Quiz describe the schema of quiz content.
type Quiz struct {
	ID          int    `db:"id" json:"_id"`
	Number      int    `db:"number" json:"number"`
	Description string `db:"description" json:"description"`
	Hint        string `db:"hint" json:"hint"`
	Score       int    `db:"score" json:"score"`
	OptionA     string `db:"option_a" json:"option_a"`
	OptionB     string `db:"option_b" json:"option_b"`
	OptionC     string `db:"option_c" json:"option_c"`
	OptionD     string `db:"option_d" json:"option_d"`
	Answer      string `db:"answer" json:"answer"`
}

// Tag describe the schema of category tags.
type Tag struct {
	ID   int    `db:"id" json:"_id"`
	Name string `db:"name" json:"name"`
}

// Provoke describe the schema of provoke mesaage.
type Provoke struct {
	ID      int    `db:"id" json:"_id"`
	Correct bool   `db:"correct" json:"correct"`
	Message string `db:"message" json:"message"`
}

// OfficerToRole describe the many-to-many relationships between officers and roles.
type OfficerToRole struct {
	ID        int `db:"id" json:"_id"`
	OfficerID int `db:"officer_id" json:"officer_id"`
	RoleID    int `db:"role_id" json:"role_id"`
}

// PlayerToQuiz describe the many-to-many relationships between players and quizes.
type PlayerToQuiz struct {
	ID       int  `db:"id" json:"_id"`
	PlayerID int  `db:"player_id" json:"player_id"`
	QuizID   int  `db:"quiz_id" json:"quiz_id"`
	Correct  bool `db:"correct" json:"correct"`
}

// QuizToTag describe the many-to-many relationships between quizes and tags.
type QuizToTag struct {
	ID     int `db:"id" json:"_id"`
	QuizID int `db:"quiz_id" json:"quiz_id"`
	TagID  int `db:"tag_id" json:"tag_id"`
}

const drop = `
DROP TABLE IF EXISTS officer_to_role;
DROP TABLE IF EXISTS player_to_quiz;
DROP TABLE IF EXISTS quiz_to_tag;
DROP TABLE IF EXISTS provoke;
DROP TABLE IF EXISTS tag;
DROP TABLE IF EXISTS quiz;
DROP TABLE IF EXISTS role;
DROP TABLE IF EXISTS player;
DROP TABLE IF EXISTS officer;
`

const schema = `
CREATE TABLE IF NOT EXISTS officer (
	id INT GENERATED ALWAYS AS IDENTITY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS role (
	id INT GENERATED ALWAYS AS IDENTITY,
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS player (
	id INT GENERATED ALWAYS AS IDENTITY,
	name VARCHAR(255) NOT NULL,
	nickname VARCHAR(255) NOT NULL,
	platform VARCHAR(255) NOT NULL,
	PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS quiz (
	id INT GENERATED ALWAYS AS IDENTITY,
	number INT NOT NULL,
	description VARCHAR(2048) NOT NULL,
	hint VARCHAR(512) NOT NULL,
	score INT NOT NULL,
	option_a VARCHAR(255) NOT NULL,
	option_b VARCHAR(255) NOT NULL,
	option_c VARCHAR(255) NOT NULL,
	option_d VARCHAR(255) NOT NULL,
	answer VARCHAR(255) NOT NULL,
	PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS tag (
	id INT GENERATED ALWAYS AS IDENTITY,
	name VARCHAR(255) NOT NULL,
	PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS provoke (
	id INT GENERATED ALWAYS AS IDENTITY,
	correct BOOLEAN NOT NULL,
	message VARCHAR(255) NOT NULL,
	PRIMARY KEY(id)
);
CREATE TABLE IF NOT EXISTS officer_to_role (
	id INT GENERATED ALWAYS AS IDENTITY,
	officer_id INT NOT NULL,
	role_id INT NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_officer
		FOREIGN KEY(officer_id)
			REFERENCES officer(id)
			ON DELETE CASCADE,
	CONSTRAINT fk_role
		FOREIGN KEY(role_id)
			REFERENCES role(id)
			ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS player_to_quiz (
	id INT GENERATED ALWAYS AS IDENTITY,
	player_id INT NOT NULL,
	quiz_id INT NOT NULL,
	correct BOOLEAN NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_player
		FOREIGN KEY(player_id)
			REFERENCES player(id)
			ON DELETE CASCADE,
	CONSTRAINT fk_quiz
		FOREIGN KEY(quiz_id)
			REFERENCES quiz(id)
			ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS quiz_to_tag (
	id INT GENERATED ALWAYS AS IDENTITY,
	quiz_id INT NOT NULL,
	tag_id INT NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_quiz
		FOREIGN KEY(quiz_id)
			REFERENCES quiz(id)
			ON DELETE CASCADE,
	CONSTRAINT fk_tag
		FOREIGN KEY(tag_id)
			REFERENCES tag(id)
			ON DELETE CASCADE
);
`

var (
	database *sqlx.DB
	err      interface{}
)

// ConnectDatabase build the connection with database.
func ConnectDatabase(reset bool) {
	connStr := "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s"
	connStr = fmt.Sprintf(
		connStr,
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USERNAME"),
		os.Getenv("PG_DBNAME"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_SSLMODE"),
	)

	database, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if reset {
		database.MustExec(drop)
		log.Print("database has been reset")
	}
	database.MustExec(schema)
}

// DisconnectDatabase break the connection with database.
func DisconnectDatabase() {

	database.Close()
}
