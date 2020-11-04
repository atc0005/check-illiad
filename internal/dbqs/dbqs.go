// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/check-illiad
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package dbqs

import (
	"database/sql"
	"fmt"
)

// RowsCount returns the number of rows for a specified table.
func RowsCount(db *sql.DB, table string) (int, error) {
	var rowsCount int
	rcQuery := fmt.Sprintf("SELECT COUNT(*) as count FROM %s", table)
	if err := db.QueryRow(rcQuery).Scan(&rowsCount); err != nil {
		return -1, fmt.Errorf(
			"failed to retrieve row count for table %s: %w",
			table,
			err,
		)
	}

	return rowsCount, nil
}
