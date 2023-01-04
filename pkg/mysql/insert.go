package mysql

import "strings"

func Insert(sql string, args ...interface{}) (int, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func GetPlaceholderN(n int) string {
	list := make([]string, 0, n)
	for i := 0; i < n; i++ {
		list = append(list, "?")
	}

	return strings.Join(list, ",")
}

func GetPlaceholderByFields(fields string) string {
	return GetPlaceholderN(len(strings.Split(fields, ",")))
}
