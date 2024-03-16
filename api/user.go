package api

import (
	"net/http"
    "github.com/gin-gonic/gin"
	
   dbmigrate "github.com/techschool/simplebank/dbmigrate/sqlc"
)

type createUserRequest struct {
	Username string  `json:"username" binding:"required,alphanum"`
	Role     string `json:"role"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	arg := dbmigrate.CreateUsersParams{
		Username: req.Username,
		Role: req.Role,
	}

	user, err := server.store.CreateUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
