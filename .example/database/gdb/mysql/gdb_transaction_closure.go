package main

import (
	"github.com/xhyonline/gf16/database/gdb"
	"github.com/xhyonline/gf16/frame/g"
)

func main() {
	var (
		err   error
		db    = g.DB()
		table = "user"
	)
	if err = db.Transaction(func(tx *gdb.TX) error {
		// Nested transaction 1.
		if err = tx.Transaction(func(tx *gdb.TX) error {
			_, err = tx.Model(table).Data(g.Map{"id": 1, "name": "john"}).Insert()
			return err
		}); err != nil {
			return err
		}
		// Nested transaction 2, panic.
		if err = tx.Transaction(func(tx *gdb.TX) error {
			_, err = tx.Model(table).Data(g.Map{"id": 2, "name": "smith"}).Insert()
			// Create a panic that can make this transaction rollback automatically.
			panic("error")
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		panic(err)
	}
}
