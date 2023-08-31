package tests

import (
	"context"
	db "github.com/Bakhram74/advertisement-server.git/db/sqlc"
	"github.com/Bakhram74/advertisement-server.git/pkg/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func randomUser(t *testing.T) db.User {
	arg := db.CreateUserParams{
		Username:    utils.RandomString(6),
		PhoneNumber: utils.RandomNumbers(9),
		Password:    utils.RandomString(8),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.Role)
	require.Equal(t, false, user.IsBanned)
	require.Equal(t, "user", user.Role)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)
	return user
}

func TestCreateUser(t *testing.T) {
	randomUser(t)
}

func TestUpdateUser(t *testing.T) {
	user := randomUser(t)
	params := db.PartialUpdateUserParams{
		ID:                user.ID,
		Username:          "Alex",
		UpdateUsername:    true,
		PhoneNumber:       utils.RandomNumbers(7),
		UpdatePhoneNumber: true,
		Password:          utils.RandomString(8),
		UpdatePassword:    true,
	}
	updateUser, err := testQueries.PartialUpdateUser(context.Background(), params)
	require.NoError(t, err)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.Role)
	require.Equal(t, false, user.IsBanned)

	require.NotEqual(t, user.Username, updateUser.Username)
	require.NotEqual(t, user.PhoneNumber, updateUser.PhoneNumber)
	require.NotEqual(t, user.Password, updateUser.Password)
	require.Equal(t, "Alex", updateUser.Username)
	require.Equal(t, 7, len(updateUser.PhoneNumber))
}

func TestUpdateUserName(t *testing.T) {
	user := randomUser(t)
	params := db.PartialUpdateUserParams{
		ID:             user.ID,
		Username:       "Pedro",
		UpdateUsername: true,
	}
	updateUser, err := testQueries.PartialUpdateUser(context.Background(), params)
	require.NoError(t, err)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.Role)
	require.Equal(t, false, user.IsBanned)
	require.NotEqual(t, user.Username, updateUser.Username)
	require.Equal(t, user.PhoneNumber, updateUser.PhoneNumber)
	require.Equal(t, user.Password, updateUser.Password)
	require.Equal(t, "Pedro", updateUser.Username)
	require.Equal(t, user.PhoneNumber, updateUser.PhoneNumber)
}
func TestUpdateUserPhoneNumber(t *testing.T) {
	user := randomUser(t)
	params := db.PartialUpdateUserParams{
		ID:                user.ID,
		PhoneNumber:       utils.RandomNumbers(4),
		UpdatePhoneNumber: true,
	}
	updateUser, err := testQueries.PartialUpdateUser(context.Background(), params)
	require.NoError(t, err)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.Role)
	require.Equal(t, false, user.IsBanned)

	require.Equal(t, user.Username, updateUser.Username)
	require.NotEqual(t, user.PhoneNumber, updateUser.PhoneNumber)
	require.Equal(t, 4, len(updateUser.PhoneNumber))
	require.Equal(t, user.Password, updateUser.Password)
	require.Equal(t, user.Username, updateUser.Username)
}

func TestUpdateUserPassword(t *testing.T) {
	user := randomUser(t)
	params := db.PartialUpdateUserParams{
		ID:             user.ID,
		Password:       utils.RandomString(12),
		UpdatePassword: true,
	}
	updateUser, err := testQueries.PartialUpdateUser(context.Background(), params)
	require.NoError(t, err)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.Role)
	require.Equal(t, false, user.IsBanned)

	require.Equal(t, user.Username, updateUser.Username)
	require.NotEqual(t, user.Password, updateUser.Password)
	require.Equal(t, 12, len(updateUser.Password))
	require.Equal(t, user.PhoneNumber, updateUser.PhoneNumber)
	require.Equal(t, user.Username, updateUser.Username)
}
