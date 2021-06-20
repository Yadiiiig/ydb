package main

type TableQuery struct {
	Conn  *Connection
	Table string
}

func (s Connection) table(name string) *TableQuery {
	return &TableQuery{
		Conn:  &s,
		Table: name,
	}
}
