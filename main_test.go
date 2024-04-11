package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestSelectClientWhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1
	cl, err := selectClient(db, clientID)

	// напиши тест здесь
	require.NoError(t, err)
	assert.Equal(t, clientID, cl.ID)
	assert.NotEmpty(t, cl.Birthday, cl.Email, cl.FIO, cl.Login)
}

func TestSelectClientWhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1
	cl, err := selectClient(db, clientID)

	// напиши тест здесь
	require.Equal(t, sql.ErrNoRows, err)
	assert.Empty(t, cl.Birthday, cl.Email, cl.FIO, cl.Login, cl.ID)
}

func TestInsertClientThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	cl.ID, err = insertClient(db, cl)

	// напиши тест здесь
	require.NoError(t, err)
	require.NotEmpty(t, cl.ID)

	stored, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	assert.Equal(t, cl, stored)
}

func TestInsertClientDeleteClientThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}
	id, err := insertClient(db, cl)

	// напиши тест здесь
	require.NoError(t, err)
	require.NotEmpty(t, id)

	_, err = selectClient(db, id)
	require.NoError(t, err)

	err = deleteClient(db, id)
	require.NoError(t, err)

	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err)
}
