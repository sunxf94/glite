package mysql

import "github.com/sunxf94/glite/pkg/logger"

func Get(dest interface{}, query string, args ...interface{}) error {
	err := db.Get(dest, query, args...)
	lg := logger.Info()
	if err != nil {
		lg = logger.Error().Err(err)
	}
	lg.Interface("sql", query).Interface("args", args).Msg("mysql.get")

	return err
}

func IsNoRowsErr(err error) bool {
	return err.Error() == "sql: no rows in result set"
}
