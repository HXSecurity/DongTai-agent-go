/*
 *  @author: Breaker
 *  @date: 2022/3/28 6:28 PM
 */

package hook

type Chi struct {
}

func (c *Chi) GetHook() []string {
	return []string{
		"httpRequestCookie",
		"chiServerHTTP",
	}
}

func (c *Chi) HookAll() {
	Hook(c.GetHook())
}

func (c *Chi) UnHookAll() {
	UnHook(c.GetHook())
}
