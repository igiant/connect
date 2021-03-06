package connect

import "encoding/json"

// @defgroup SUBGROUP4 Definitions

// TimeRangeType - @{
type TimeRangeType string

const (
	TimeRangeDaily      TimeRangeType = "TimeRangeDaily"
	TimeRangeWeekly     TimeRangeType = "TimeRangeWeekly"
	TimeRangeAbsolute   TimeRangeType = "TimeRangeAbsolute"
	TimeRangeChildGroup TimeRangeType = "TimeRangeChildGroup" // not supported in QT
)

type UtcDateTime string

type DayType string

const (
	Monday    DayType = "Monday"
	Tuesday   DayType = "Tuesday"
	Wednesday DayType = "Wednesday"
	Thursday  DayType = "Thursday"
	Friday    DayType = "Friday"
	Saturday  DayType = "Saturday"
	Sunday    DayType = "Sunday"
)

type DayList []DayType

type TimeRangeGroup struct {
	Id   KId    `json:"id"`
	Name string `json:"name"`
}

// Note: If type is changed, all fields representing the associated value must be also assigned, if used in set method.
//    And conversely, type must be assigned if value was changed.

// TimeRangeEntry -    type + (fromDay, toDay) + (fromDate, toDate) + childGroupId
type TimeRangeEntry struct {
	Id             KId           `json:"id"`
	GroupId        KId           `json:"groupId"`
	SharedId       KId           `json:"sharedId"` // read-only; filled when the item is shared in MyKerio
	GroupName      string        `json:"groupName"`
	Description    string        `json:"description"`
	Type           TimeRangeType `json:"type"`
	Enabled        bool          `json:"enabled"`
	Status         StoreStatus   `json:"status"`
	FromTime       Time          `json:"fromTime"` // This doesn't contain seconds, so we round data created by QT admin
	ToTime         Time          `json:"toTime"`   // This doesn't contain seconds, so we round data created by QT admin
	Days           DayList       `json:"days"`
	FromDay        DayType       `json:"fromDay"`
	ToDay          DayType       `json:"toDay"`
	FromDate       Date          `json:"fromDate"` // hour and min used from Time
	ToDate         Date          `json:"toDate"`   // hour and min used from Time
	ChildGroupId   KId           `json:"childGroupId"`
	ChildGroupName string        `json:"childGroupName"`
}

type TimeRangeEntryList []TimeRangeEntry

type TimeRangeGroupList []TimeRangeGroup

// TimeRangesCreate - Create new ranges.
//	ranges - details for new ranges. Field KiD is assigned by the manager to temporary value until apply() or reset().
// Return
//	errors - possible errors: - "This time range already exists!" duplicate name-value
//	result - list of IDs of created TimeRanges
func (s *ServerConnection) TimeRangesCreate(ranges TimeRangeEntryList) (ErrorList, CreateResultList, error) {
	params := struct {
		Ranges TimeRangeEntryList `json:"ranges"`
	}{ranges}
	data, err := s.CallRaw("TimeRanges.create", params)
	if err != nil {
		return nil, nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList        `json:"errors"`
			Result CreateResultList `json:"result"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, errors.Result.Result, err
}

// TimeRangesGet - Get the list of ranges.
//	query - conditions and limits. Included from weblib. KWF engine implementation notes: - LIKE matches substring (second argument) in a string (first argument). There are no wildcards. - sort and match are case insensitive. - column alias (first operand):
// Return
//  list - list of ranges and its details
//  totalItems - count of all ranges on the server (before the start/limit applied)
func (s *ServerConnection) TimeRangesGet(query SearchQuery) (TimeRangeEntryList, int, error) {
	query = addMissedParametersToSearchQuery(query)
	params := struct {
		Query SearchQuery `json:"query"`
	}{query}
	data, err := s.CallRaw("TimeRanges.get", params)
	if err != nil {
		return nil, 0, err
	}
	list := struct {
		Result struct {
			List       TimeRangeEntryList `json:"list"`
			TotalItems int                `json:"totalItems"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &list)
	return list.Result.List, list.Result.TotalItems, err
}

// TimeRangesGetGroupList - Get the list of groups, sorted in ascending order.
// Return
//	groups - list of all groups
func (s *ServerConnection) TimeRangesGetGroupList() (TimeRangeGroupList, error) {
	data, err := s.CallRaw("TimeRanges.getGroupList", nil)
	if err != nil {
		return nil, err
	}
	groups := struct {
		Result struct {
			Groups TimeRangeGroupList `json:"groups"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &groups)
	return groups.Result.Groups, err
}

// TimeRangesRemove - Remove ranges.
//	rangeIds - IDs of ranges that should be removed
// Return
//	errors - Errors by removing ranges
func (s *ServerConnection) TimeRangesRemove(rangeIds KIdList) (ErrorList, error) {
	params := struct {
		RangeIds KIdList `json:"rangeIds"`
	}{rangeIds}
	data, err := s.CallRaw("TimeRanges.remove", params)
	if err != nil {
		return nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList `json:"errors"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, err
}

// TimeRangesSet - Set existing ranges.
//	rangeIds - IDs of ranges to be updated.
//	details - details for update. Field "KId" is ignored. All other fields must be filled in and they are written to all ranges specified by rangeIds.
// Return
//	errors - possible errors: - "This time range already exists!" duplicate name-value
func (s *ServerConnection) TimeRangesSet(rangeIds KIdList, details TimeRangeEntry) (ErrorList, error) {
	params := struct {
		RangeIds KIdList        `json:"rangeIds"`
		Details  TimeRangeEntry `json:"details"`
	}{rangeIds, details}
	data, err := s.CallRaw("TimeRanges.set", params)
	if err != nil {
		return nil, err
	}
	errors := struct {
		Result struct {
			Errors ErrorList `json:"errors"`
		} `json:"result"`
	}{}
	err = json.Unmarshal(data, &errors)
	return errors.Result.Errors, err
}
