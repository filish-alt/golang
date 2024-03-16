package api

import (
	"github.com/gin-gonic/gin"
	dbmigrate "github.com/techschool/simplebank/dbmigrate/sqlc"
)
type Server struct{
	store *dbmigrate.Store
	router *gin.Engine
}

func NewServer(store *dbmigrate.Store) *Server{
         server := &Server{store: store}
		 router := gin.Default()
router.POST("/account",server.CreateAccount)

       
		 server.router=router
		 return server
} 
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H{
	return gin.H{"error":err.Error()}
}