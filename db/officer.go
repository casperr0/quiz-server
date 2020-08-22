package db

import (
	"fmt"

	// posetgreSQL databse driver required by sqlx
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreateOfficer will create a new officer account by username and password.
func CreateOfficer(username, password string) error {

	officer, _ := GetOfficer(username)
	if officer != nil {
		tpl := "officer %s already existed"
		return fmt.Errorf(tpl, username)
	}

	createSQL := `
	INSERT INTO officer (username, password)
	SELECT $1::VARCHAR, $2::VARCHAR
	WHERE NOT EXISTS (
		SELECT 1 FROM officer
		WHERE officer.username = $1
	);
	`
	tx := database.MustBegin()
	tx.MustExec(createSQL, username, password)
	tx.Commit()
	return nil
}

// GetOfficer will get the officer account with specified username.
func GetOfficer(username string) (*Officer, error) {

	getSQL := "SELECT * FROM officer WHERE username=$1"
	officer := Officer{}
	err := database.Get(&officer, getSQL, username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "officer %s not found"
			return nil, fmt.Errorf(tpl, username)
		}
		return nil, err
	}
	return &officer, nil
}

// ListOfficers will list all current officers.
func ListOfficers() ([]Officer, error) {

	listSQL := "SELECT * FROM officer"
	var officers []Officer
	err := database.Select(&officers, listSQL)
	if err != nil {
		return nil, err
	}
	return officers, nil
}

// QueryOfficers query all officers with the role.
func QueryOfficers(roleName string) ([]Officer, error) {

	querySQL := `
	SELECT o.id, o.username, o.password
	FROM   officer o
	JOIN officer_to_role o_r ON o.id = o_r.officer_id
	JOIN role r ON r.id = o_r.role_id
    WHERE r.name = $1
	`
	var officers []Officer
	err := database.Select(&officers, querySQL, roleName)
	if err != nil {
		return nil, err
	}
	return officers, nil
}

// DeleteOfficer will delete officer with specified username.
func DeleteOfficer(username string) {

	deleteSQL := "DELETE FROM officer WHERE username=$1"
	tx := database.MustBegin()
	tx.MustExec(deleteSQL, username)
	tx.Commit()
}

// CreateRole will create a access role account.
func CreateRole(name string) {

	createSQL := `
	INSERT INTO role (name)
	SELECT $1::VARCHAR
	WHERE NOT EXISTS (
		SELECT 1 FROM role
		WHERE role.name = $1
	);
	`
	tx := database.MustBegin()
	tx.MustExec(createSQL, name)
	tx.Commit()
}

// GetRole will get the role with specified name.
func GetRole(roleName string) (*Role, error) {

	getSQL := "SELECT * FROM role WHERE name=$1"
	role := Role{}
	err := database.Get(&role, getSQL, roleName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "role %s not found"
			return nil, fmt.Errorf(tpl, roleName)
		}
		return nil, err
	}
	return &role, nil
}

// ListRoles will list all current roles.
func ListRoles() ([]Role, error) {

	listSQL := "SELECT * FROM role"
	var roles []Role
	err := database.Select(&roles, listSQL)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// QueryRoles query all roles of an officer.
func QueryRoles(officerName string) ([]Role, error) {

	querySQL := `
	SELECT r.id, r.name
	FROM   role r
	JOIN officer_to_role o_r ON r.id = o_r.role_id
	JOIN officer o ON o.id = o_r.officer_id
    WHERE o.username = $1
	`
	var roles []Role
	err := database.Select(&roles, querySQL, officerName)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// DeleteRole will delete role with specified name.
func DeleteRole(name string) {

	deleteSQL := "DELETE FROM role WHERE name=$1"
	tx := database.MustBegin()
	tx.MustExec(deleteSQL, name)
	tx.Commit()
}

// RegisterRole will register an officer with an access role.
func RegisterRole(officerName, roleName string) error {

	registerSQL := `
	INSERT INTO officer_to_role (officer_id, role_id)
	SELECT $1::INT, $2::INT
	WHERE NOT EXISTS(
		SELECT * FROM officer_to_role
		WHERE officer_id = $1 AND role_id = $2
	);
	`
	officerFound, err := GetOfficer(officerName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "officer %s not found"
			return fmt.Errorf(tpl, officerName)
		}
		return err
	}
	officerID := officerFound.ID
	roleFound, err := GetRole(roleName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "role %s not found"
			return fmt.Errorf(tpl, roleName)
		}
		return err
	}
	roleID := roleFound.ID
	tx := database.MustBegin()
	tx.MustExec(registerSQL, officerID, roleID)
	tx.Commit()
	return nil
}

// DeregisterRole will deregister an officer with an access role.
func DeregisterRole(officerName, roleName string) error {

	deregisterSQL := `
	DELETE FROM officer_to_role
	WHERE officer_id = $1 AND role_id = $2;
	`
	officerFound, err := GetOfficer(officerName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "officer %s not found"
			return fmt.Errorf(tpl, officerName)
		}
		return err
	}
	officerID := officerFound.ID
	roleFound, err := GetRole(roleName)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			tpl := "role %s not found"
			return fmt.Errorf(tpl, roleName)
		}
		return err
	}
	roleID := roleFound.ID
	tx := database.MustBegin()
	tx.MustExec(deregisterSQL, officerID, roleID)
	tx.Commit()
	return nil
}
