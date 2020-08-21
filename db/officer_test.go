package db

import (
	"fmt"
	"testing"

	"github.com/ccns/quiz-server/config"
	"github.com/stretchr/testify/assert"
)

var (
	officers []Officer = []Officer{
		Officer{1, "rain", "0114"},
		Officer{2, "!*&T*GQ9qw8gd971", "@)()#)I@BKQHJ"},
	}
)

func TestCreateOfficer(t *testing.T) {

	err := CreateOfficer(officers[0].Username, officers[0].Password)
	assert.Nil(t, err)
	result, err := GetOfficer(officers[0].Username)
	assert.Nil(t, err)
	assert.Equal(t, result.Password, officers[0].Password)

	err = CreateOfficer(officers[1].Username, officers[1].Password)
	assert.Nil(t, err)
	result, err = GetOfficer(officers[1].Username)
	assert.Nil(t, err)
	assert.Equal(t, result.Password, officers[1].Password)

	err = CreateOfficer(officers[1].Username, officers[1].Password)
	assert.Equal(
		t,
		err,
		fmt.Errorf("officer %s already existed", officers[1].Username),
		"Should alert while creating duplicated officer.",
	)
}

func TestGetOfficer(t *testing.T) {

	result, err := GetOfficer(officers[1].Username)
	assert.Nil(t, err)
	assert.Equal(t, result.Password, officers[1].Password)
}

func TestListOfficers(t *testing.T) {

	result, err := ListOfficers()
	assert.Nil(t, err)
	for i := range result {
		assert.Equal(t, result[i].Username, officers[i].Username, "Username not match.")
		assert.Equal(t, result[i].Password, officers[i].Password, "Password not match.")
	}
}

func TestGetRole(t *testing.T) {

	var result *Role
	for _, r := range config.Config.Officer.DefaultRoles {
		result, _ = GetRole(r)
		assert.NotNil(t, result, fmt.Sprintf("role %s not found.", r))
	}
}

func TestListRoles(t *testing.T) {

	result, _ := ListRoles()
	for i, r := range config.Config.Officer.DefaultRoles {
		assert.Equal(t, result[i].Name, r, "role not match.")
	}
}

func TestRegisterRole(t *testing.T) {

	roles := config.Config.Officer.DefaultRoles
	RegisterRole(officers[0].Username, roles[0])
	RegisterRole(officers[0].Username, roles[1])
	RegisterRole(officers[1].Username, roles[0])

	resultOfficers, err := QueryOfficers(roles[0])
	assert.Nil(t, err)
	for i, o := range resultOfficers {
		assert.Equal(t, officers[i].Username, o.Username)
		assert.Equal(t, officers[i].Password, o.Password)
	}

	resultRoles, err := QueryRoles(officers[0].Username)
	assert.Nil(t, err)
	for i, r := range resultRoles {
		assert.Equal(t, roles[i], r.Name)
	}
}
