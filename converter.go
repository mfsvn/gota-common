package gota_common

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func ToString(x interface{}) string {
	switch x.(type) {
	case int:
		return fmt.Sprintf("%d", x)
	case int64:
		return fmt.Sprintf("%d", x)
	case int32:
		return fmt.Sprintf("%d", x)
	case sql.NullTime:
		y, _ := x.(sql.NullTime)
		if y.Valid {
			return y.Time.Format(DateTimeLayout)
		} else {
			return "null"
		}
	case time.Time:
		y, _ := x.(time.Time)
		return y.String()
	case sql.NullInt32:
		y, _ := x.(sql.NullInt32)
		if y.Valid {
			return fmt.Sprintf("%d", y.Int32)
		} else {
			return "null"
		}
	case sql.NullInt64:
		y, _ := x.(sql.NullInt64)
		if y.Valid {
			return fmt.Sprintf("%d", y.Int64)
		} else {
			return "null"
		}
	case sql.NullString:
		y, _ := x.(sql.NullString)
		if y.Valid {
			return fmt.Sprintf("%s", y.String)
		} else {
			return "null"
		}
	case sql.NullFloat64:
		y, _ := x.(sql.NullFloat64)
		if y.Valid {
			return fmt.Sprintf("%f", y.Float64)
		} else {
			return "null"
		}
	case string:
		y, _ := x.(string)
		return y
	case []byte:
		y, _ := x.([]byte)
		return string(y)
	default:
		return "???"
	}
	return "???"
}

func CloseRows(rows *sql.Rows) {
	err := rows.Close()
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		log.Panicln()
	}
}
