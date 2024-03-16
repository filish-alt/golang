package dbmigrate

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/utils"
)

 var testqueries *Queries
func TestingMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
   
	conn,err:= sql.Open(config.DBDriver,config.DBSource)
    if err !=nil{
		log.Fatal("unable to connect",err)
	}
testqueries = New(conn)

os.Exit(m.Run())
}

