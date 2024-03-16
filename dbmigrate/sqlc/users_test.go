package dbmigrate

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestUser(t *testing.T){
    args := CreateUsersParams{
		ID: 1,
		Username: "gfhgfh", // Initialize Username with a default value
		Role:     "",

	}
user, err :=testqueries.CreateUsers(context.Background(),args)

require.NoError(t,err)
require.NotEmpty(t,user)

require.Equal(t,args.ID,user.ID)
require.Equal(t,args.Username,user.Username)


}