package hook

type Gorm struct {
}

func (g *Gorm) GetHook() []string {
	return []string{
		"gormDBOrder",
		"gormDBExec",
		"gormDBGroup",
		"gormDBHaving",
		"gormDBPluck",
		"gormDBRaw",
		"gormDBSelect",
	}
}

func (g *Gorm) HookAll() {
	Hook(g.GetHook())
}

func (g *Gorm) UnHookAll() {
	UnHook(g.GetHook())
}
