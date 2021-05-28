package connect

// KId - global object identification
type KId string

// KIdList - list of global object identifiers
type KIdList []KId

// SortDirection - Sorting Direction
type SortDirection string

const (
	Asc  SortDirection = "Asc"
	Desc SortDirection = "Desc"
)

// CompareOperator - Simple Query Operator
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

// LogicalOperator - Compound Operator
type LogicalOperator string

const (
	Or  LogicalOperator = "Or"
	And LogicalOperator = "And"
)

// NamedValue - all fields must be assigned if used in set methods
type NamedValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NamedMultiValue - Name-multivalue pair.
// Note: all fields must be assigned if used in set methods
type NamedMultiValue struct {
	Name  string   `json:"name"`
	Value []string `json:"value"`
}

// SubCondition - A Part of a Condition
type SubCondition struct {
	FieldName  string          `json:"fieldName"`  // left side of condition
	Comparator CompareOperator `json:"comparator"` // middle of condition
	Value      string          `json:"value"`      // right side of condition
}

type SubConditionList []SubCondition

// SortOrder - Sorting Order
type SortOrder struct {
	ColumnName    string        `json:"columnName"`
	Direction     SortDirection `json:"direction"`
	CaseSensitive bool          `json:"caseSensitive"`
}

type SortOrderList []SortOrder

// SearchQuery - General Query for Searching
// Query substitution (quicksearch):
// SearchQuery doesn't support complex queries, only queries
// with all AND operators (or all OR operators) are supported.
// Combination of AND and OR is not allowed. This limitation is for special cases solved by using
// substitution of complicated query-part by simple condition.
// Only the quicksearch is currently implemented and only in "UsersGet()" method.
// Behavior of quicksearch in Users::get():
// QUICKSEACH    = "x"   is equal to:   (loginName    = "x")  OR (fullName    = "x")
// QUICKSEACH LIKE "x*"  is equal to:   (loginName LIKE "x*") OR (fullName LIKE "x*")
// QUICKSEACH   <> "x"   is equal to:   (loginName   <> "x") AND (fullName   <> "x")
type SearchQuery struct {
	Fields     []string         `json:"fields"`
	Conditions SubConditionList `json:"conditions"`
	Combining  LogicalOperator  `json:"combining"`
	Start      int              `json:"start"`
	Limit      int              `json:"limit"`
	OrderBy    SortOrderList    `json:"orderBy,omitempty"`
}

// LocalizableMessage can contain replacement marks:  { "User %1 cannot be deleted.", ["jsmith"], 1 }
type LocalizableMessage struct {
	Message              string   `json:"message"`              // text with placeholders %1, %2, etc., e.g. "User %1 cannot be deleted."
	PositionalParameters []string `json:"positionalParameters"` // additional strings to replace the placeholders in message (first string replaces %1 etc.)
	Plurality            int      `json:"plurality"`            // count of items, used to distinguish among singular/paucal/plural; 1 for messages with no counted items
}

type LocalizableMessageList []LocalizableMessage

// ManipulationError - error structure to be used when manipulating with globally addressable list items
type ManipulationError struct {
	ID           KId                `json:"id"` // entity KId, can be user, group, alias, ML...
	ErrorMessage LocalizableMessage `json:"errorMessage"`
}

type ManipulationErrorList []ManipulationError

// RestrictionKind - A kind of restriction
type RestrictionKind string

const (
	Regex                  RestrictionKind = "Regex"                  // regular expression
	ByteLength             RestrictionKind = "ByteLength"             // maximal length in Bytes
	ForbiddenNameList      RestrictionKind = "ForbiddenNameList"      // list of denied exact names due to filesystem or KMS store
	ForbiddenPrefixList    RestrictionKind = "ForbiddenPrefixList"    // list of denied prefixes due to filesystem or KMS store
	ForbiddenSuffixList    RestrictionKind = "ForbiddenSuffixList"    // list of denied suffixes due to filesystem or KMS store
	ForbiddenCharacterList RestrictionKind = "ForbiddenCharacterList" // list of denied characters
)

// ItemName - Item of the Entity; used in restrictions
type ItemName string

const (
	Name        ItemName = "Name"        // Entity Name
	Description ItemName = "Description" // Entity Description
	Email       ItemName = "Email"       // Entity Email Address
	FullName    ItemName = "FullName"    // Entity Full Name
	TimeItem    ItemName = "TimeItem"    // Entity Time - it cannot be simply Time because of C++ conflict - see bug 34684 comment #3
	DateItem    ItemName = "DateItem"    // Entity Date - I expect same problem with Date as with Time
	DomainName  ItemName = "DomainName"  ///< differs from name (eg. cannot contains underscore)
)

// ByteUnits - Units used for handling large values of bytes.
type ByteUnits string

const (
	Bytes     ByteUnits = "Bytes"
	KiloBytes ByteUnits = "KiloBytes"
	MegaBytes ByteUnits = "MegaBytes"
	GigaBytes ByteUnits = "GigaBytes"
	TeraBytes ByteUnits = "TeraBytes"
	PetaBytes ByteUnits = "PetaBytes"
)

// ByteValueWithUnits - Stores size of very large values of bytes e.g. for user quota
// Note: all fields must be assigned if used in set methods
type ByteValueWithUnits struct {
	Value int       `json:"value"`
	Units ByteUnits `json:"units"`
}

// SizeLimit - Settings of size limit
// Note: all fields must be assigned if used in set methods
type SizeLimit struct {
	IsActive bool               `json:"isActive"`
	Limit    ByteValueWithUnits `json:"limit"`
}

// AddResult - Result of the add operation
type AddResult struct {
	ID           KId                `json:"id"`           // purposely not id - loginName is shown
	Success      bool               `json:"success"`      // was operation successful? if yes so id is new id for this item else errorMessage tells why it failed
	ErrorMessage LocalizableMessage `json:"errorMessage"` // contains number of recovered user messages or error message
}

// AddResultList - list of add operation results
type AddResultList []AddResult

type IpAddress string

type IpAddressList []IpAddress

// StoreStatus - Status of entry in persistent manager
type StoreStatus string

const (
	StoreStatusClean    StoreStatus = "StoreStatusClean"    // already present in configuration store
	StoreStatusModified StoreStatus = "StoreStatusModified" // update waiting for apply()
	StoreStatusNew      StoreStatus = "StoreStatusNew"      // added to manager but not synced to configuration store
)

/*
 When using start and limit to only get a part of all results
 (e.g. only 20 users, skipping the first 40 users),
 use this special limit value for unlimited count
 (of course the service still respects the value of start).

 Note that each service is allowed to use its safety limit
 (such as 50,000) to prevent useless overload.
 The limits are documented per-service or per-method.

 Implementation note: Some source code transformations may lead to signed long, i.e. 4294967295.
 But the correct value is -1.
*/

const Unlimited int = -1

// Time - should be used instead of time_t, where time zones can affect time interpretation
// Note: all fields must be assigned if used in set methods
type Time struct {
	Hour int `json:"hour"`
	Min  int `json:"min"`
}

// Date - should be used instead of time_t, where time zones can affect time interpretation
// Note: all fields must be assigned if used in set methods
type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type OptionalString struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type OptionalLong struct {
	Enabled bool `json:"enabled"`
	Value   int  `json:"value"`
}

type OptionalEntity struct {
	Enabled bool `json:"enabled"`
	ID      KId  `json:"id"`
	Name    int  `json:"name"`
}

type LocalizableMessageParameters struct {
	PositionalParameters []string `json:"positionalParameters"` // additional strings to replace the placeholders in message (first string replaces %1 etc.)
	Plurality            int      `json:"plurality"`            // count of items, used to distinguish among singular/paucal/plural; 1 for messages with no counted items
}

type Error struct {
	InputIndex        int                          `json:"inputIndex"`        // 0-based index to input array, e.g. 3 means that the relates to the 4th element of the input parameter array
	Code              int                          `json:"code"`              //-32767..-1 (JSON-RPC) or 1..32767 (application)
	Message           string                       `json:"message"`           // text with placeholders %1, %2, etc., e.g. "User %1 cannot be deleted."
	MessageParameters LocalizableMessageParameters `json:"messageParameters"` // strings to replace placeholders in message, and message plurality.
}

type ErrorList []Error

type Download struct {
	URL    string `json:"url"`    // download url
	Name   string `json:"name"`   // filename
	Length int    `json:"length"` // file size in bytes
}

// CreateResult - Details about a particular item created.
type CreateResult struct {
	InputIndex int    `json:"inputIndex"` // 0-based index to input array, e.g. 3 means that the relates to the 4th element of the input parameter array
	ID         string `json:"id"`         // ID of created item
}

type CreateResultList []CreateResult

type DateTimeStamp int

type ApiApplication struct {
	Name    string
	Vendor  string
	Version string
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
