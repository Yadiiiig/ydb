package lib

type TableQuery struct {
	Conn  *Connection
	Table string
}

func (s Connection) Table(name string) *TableQuery {
	return &TableQuery{
		Conn:  &s,
		Table: name,
	}
}
