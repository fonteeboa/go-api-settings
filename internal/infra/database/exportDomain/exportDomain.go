package exportDomain

// QueryParams é uma estrutura para representar os parâmetros de consulta dinâmica.
type QueryParams struct {
	Filters []Filter
	Joins   []Join
}

// Filter representa um filtro para a consulta.
type Filter struct {
	Field    string
	Operator string
	Value    interface{}
}

// Join representa uma cláusula JOIN para a consulta.
type Join struct {
	Table     string
	Condition string
	JoinType  string // "INNER JOIN", "LEFT JOIN", etc.
}

// ForeignKeyConfig representa a configuração de uma chave estrangeira
type ForeignKeyConfig struct {
	Model          interface{}
	ForeignKeyName string
	References     string
	OnUpdate       string
	OnDelete       string
}
