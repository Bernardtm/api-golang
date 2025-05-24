package utils

import "fmt"

func Pagination(page int, limit int, query *string, args *[]interface{}, argIndex *int) {
	if limit >= 0 && page > 0 {
		offset := (page - 1) * limit
		*query += fmt.Sprintf("\n\t LIMIT $%d ", *argIndex)
		*args = append(*args, limit)
		(*argIndex)++
		*query += fmt.Sprintf(" OFFSET $%d ", *argIndex)
		*args = append(*args, offset)
		(*argIndex)++
	}
}
