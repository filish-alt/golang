package dbmigrate

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T){
    args := CreateUsersParams{
		ID: 1,
		Username: pgtype.Text{String: "default_username"}, // Initialize Username with a default value
		Role:     pgtype.Text{String: "default_role"},

	}
user, err :=testqueries.CreateUsers(context.Background(),args)

require.NoError(t,err)
require.NotEmpty(t,user)

require.Equal(t,args.ID,user.ID)
require.Equal(t,args.Username,user.Username)


}