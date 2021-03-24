package model

import "fmt"

//LogFieldTemp 指定模板和字段，组建变更日志字段
func LogFieldTemp(field string, nw interface{}, old interface{}, diffTemp bool) (s string) {
	var (
		desc string
		//exist bool
		temp string
	)
	//if desc, exist = LogFieldDesc[field]; !exist {
	//	desc = field
	//}
	if diffTemp {
		temp = "[%s]从[%v]变成[%v]"
		s = fmt.Sprintf(temp, desc, old, nw)
	} else {
		temp = "[%s]为[%v]"
		s = fmt.Sprintf(temp, desc, nw)
	}

	return
}
