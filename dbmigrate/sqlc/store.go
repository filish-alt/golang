package dbmigrate

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}



func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error{
       tx,err := store.db.BeginTx(ctx,nil)

	   if err !=nil{
		return nil
	   }

q:=New(tx)
 err =fn(q)

 if err != nil{
	if rbErr := tx.Rollback(); rbErr !=nil{
		return fmt.Errorf("txt: %v rberr: %v",err,rbErr)
	}
	return err
}
return tx.Commit()
}

type TransferParam struct{
   Username string
}

type TransferResult struct{
    User User
}
func(store *Store) TransferMoney(context context.Context,transferParam TransferParam) (TransferResult, error){
    var result TransferResult

	err := store.execTx(context, func(q *Queries) error {
		var err error
		result.User, err = q.CreateUsers(context, CreateUsersParams{
               Username: transferParam.Username,
		})
		if err!=nil{
			return err
		}
		return nil
	})

	return result, err
}