package odbc

import (
	"strings"

	"github.com/alexbrainman/odbc" // DRIVER: odbc
	"github.com/xo/usql/drivers"
)

func init() {
	drivers.Register("odbc", drivers.Driver{
		LexerName: "tsql",
		IsPasswordErr: func(err error) bool {
			if e, ok := err.(*odbc.Error); ok {
				msg := strings.ToLower(e.Error())
				return strings.Contains(msg, "failed") &&
					(strings.Contains(msg, "login") ||
						strings.Contains(msg, "authentication") ||
						strings.Contains(msg, "password"))
			}
			return false
		},
	})
}
