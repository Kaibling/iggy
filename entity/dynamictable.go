package entity

type DynamicTable struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Variables []DynamicTableVariable `json:"variables"`
	Meta      MetaData               `json:"meta"`
}

type DynamicTableVariable struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	VariableType string     `json:"variable_type"`
	DynamicTable Identifier `json:"dynamic_table,omitempty"`
	Meta         MetaData   `json:"meta"`
}

type NewDynamicTable struct {
	ID   string
	Name string `json:"name"`
}

type NewDynamicTableVariable struct {
	ID             string
	Name           string `json:"name"`
	VariableType   string `json:"variable_type"`
	DynamicTableID string `json:"dynamic_table_id"`
}
