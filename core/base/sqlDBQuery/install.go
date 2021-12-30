package sqlDBQuery

import (
	"database/sql"
	"fmt"
	"github.com/HXSecurity/DongTai-agent-go/model"
	"github.com/brahma-adshonor/gohook"
)

func init() {
	model.HookMap["sqlDBQuery"] = new(SqlDBQuery)
}

type SqlDBQuery struct {
}

func (h *SqlDBQuery) Hook() {
	var db *sql.DB
	err := gohook.HookMethod(db, "Query", Query, QueryT)
	if err != nil {
		fmt.Println(err, "SqlDBQuery")
	} else {
		fmt.Println("SqlDBQuery")
	}
}

func (h *SqlDBQuery) UnHook() {
	var db *sql.DB
	gohook.UnHookMethod(db, "Query")
}
