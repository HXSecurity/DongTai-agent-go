package sqlDBQuery

import (
	"database/sql"
	"github.com/brahma-adshonor/gohook"
	"go-agent/model"
)

func init() {
	model.HookMap["sqlDBQuery"] = new(SqlDBQuery)
}

type SqlDBQuery struct {
}

func (h *SqlDBQuery) Hook() {
	var db *sql.DB
	gohook.HookMethod(db, "Query", Query, QueryT)
}

func (h *SqlDBQuery) UnHook() {
	var db *sql.DB
	gohook.UnHookMethod(db, "Query")
}
