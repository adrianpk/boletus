package kabestan

import "fmt"

func SQLComma(upc bool) string {
	if upc {
		return ", "
	}
	return " "
}

// strUpdCol build an update colum fragment of type string.
func SQLStrUpd(colName, fieldName string) string {
	return fmt.Sprintf("%s = :%s", colName, fieldName)
}

// whereID build an SQL where clause for ID.
func SQLWhereID(id string) string {
	return fmt.Sprintf("WHERE id = '%s' AND (is_deleted IS NULL OR NOT is_deleted);", id)
}
