// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/xhyonline/gf16.
//

package gdb

// MasterLink acts like function Master but with additional `schema` parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Master.
func (c *Core) MasterLink(schema ...string) (Link, error) {
	db, err := c.db.Master(schema...)
	if err != nil {
		return nil, err
	}
	return &dbLink{db}, nil
}

// SlaveLink acts like function Slave but with additional `schema` parameter specifying
// the schema for the connection. It is defined for internal usage.
// Also see Slave.
func (c *Core) SlaveLink(schema ...string) (Link, error) {
	db, err := c.db.Slave(schema...)
	if err != nil {
		return nil, err
	}
	return &dbLink{db}, nil
}

// QuoteWord checks given string `s` a word, if true quotes it with security chars of the database
// and returns the quoted string; or else return `s` without any change.
func (c *Core) QuoteWord(s string) string {
	charLeft, charRight := c.db.GetChars()
	return doQuoteWord(s, charLeft, charRight)
}

// QuoteString quotes string with quote chars. Strings like:
// "user", "user u", "user,user_detail", "user u, user_detail ut", "u.id asc".
func (c *Core) QuoteString(s string) string {
	charLeft, charRight := c.db.GetChars()
	return doQuoteString(s, charLeft, charRight)
}

// QuotePrefixTableName adds prefix string and quotes chars for the table.
// It handles table string like:
// "user", "user u",
// "user,user_detail",
// "user u, user_detail ut",
// "user as u, user_detail as ut".
//
// Note that, this will automatically checks the table prefix whether already added,
// if true it does nothing to the table name, or else adds the prefix to the table name.
func (c *Core) QuotePrefixTableName(table string) string {
	charLeft, charRight := c.db.GetChars()
	return doHandleTableName(table, c.db.GetPrefix(), charLeft, charRight)
}

// GetChars returns the security char for current database.
// It does nothing in default.
func (c *Core) GetChars() (charLeft string, charRight string) {
	return "", ""
}

// DoCommit is a hook function, which deals with the sql string before it's committed to underlying driver.
// The parameter `link` specifies the current database connection operation object. You can modify the sql
// string `sql` and its arguments `args` as you wish before they're committed to driver.
// Also see Core.DoCommit.
func (c *Core) DoCommit(sql string) string {
	return sql
}

// Tables retrieves and returns the tables of current schema.
// It's mainly used in cli tool chain for automatically generating the models.
//
// It does nothing in default.
func (c *Core) Tables(schema ...string) (tables []string, err error) {
	return
}

// TableFields retrieves and returns the fields information of specified table of current schema.
//
// Note that it returns a map containing the field name and its corresponding fields.
// As a map is unsorted, the TableField struct has a "Index" field marks its sequence in the fields.
//
// It's using cache feature to enhance the performance, which is never expired util the process restarts.
//
// It does nothing in default.
func (c *Core) TableFields(table string, schema ...string) (fields map[string]*TableField, err error) {
	return
}
