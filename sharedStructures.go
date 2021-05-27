package connect

type CompareOperator string

const (
	Eq          CompareOperator = "Eq"          // '='  - equal
	NotEq       CompareOperator = "NotEq"       // '!=' - not equal
	LessThan    CompareOperator = "LessThan"    // '<'  - lower that
	GreaterThan CompareOperator = "GreaterThan" // '>'  - greater that
	LessEq      CompareOperator = "LessEq"      // '<=' - lower or equal
	GreaterEq   CompareOperator = "GreaterEq"   // '>=' - greater or equal
	Like        CompareOperator = "NotEq"       // contains substring, % is wild character
)

// NamedValue - all fields must be assigned if used in set methods
type NamedValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NamedValue - all fields must be assigned if used in set methods
type NamedMultiValue struct {
	Name  string   `json:"name"`
	Value []string `json:"value"`
}

// SubCondition - A Part of a Condition
type SubCondition struct {
	FieldName  string `json:"fieldName"`  //left side of condition
	Comparator string `json:"comparator"` //middle of condition
	Value      string `json:"value"`      //right side of condition
}

// SortOrder - Sorting Order
type SortOrder struct {
	ColumnName    string `json:"columnName"`
	Direction     string `json:"direction"`
	CaseSensitive bool   `json:"caseSensitive"`
}
