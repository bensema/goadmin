package service

import (
	"fmt"
	"github.com/bensema/gcurd"
	"github.com/bensema/goadmin/utils"
	"github.com/fatih/structs"
	"strings"
)

const (
	AdminSession = "admin-session"
)

func logFieldTemp(column string, newVal interface{}, oldVal interface{}, isNew bool, mosaics bool) string {
	if mosaics {
		oldVal, newVal = "*", "*"
	}
	if isNew {
		return fmt.Sprintf("[%s]为[%v]", column, newVal)
	} else {
		return fmt.Sprintf("[%s]从[%v]变成[%v]", column, oldVal, newVal)
	}

}

func logFieldChange[T gcurd.Model](newObj T, oldObj T, isNew bool, mosaicsColumns []string) string {
	var s []string
	n := structs.New(newObj)
	n.TagName = "json"
	_new := n.Map()

	o := structs.New(oldObj)
	o.TagName = "json"
	_old := o.Map()

	for _, col := range newObj.Columns() {
		if isNew {
			s = append(s, logFieldTemp(col, _new[col], nil, isNew, utils.CheckIn(mosaicsColumns, col)))
			continue
		}
		if _new[col] != _old[col] {
			s = append(s, logFieldTemp(col, _new[col], _old[col], isNew, utils.CheckIn(mosaicsColumns, col)))
		}
	}
	return strings.Join(s, ";")
}

func ContentNew[T gcurd.Model](newObj T, mosaicsColumns []string) string {
	return logFieldChange[gcurd.Model](newObj, nil, true, mosaicsColumns)
}

func ContentDiff[T gcurd.Model](newObj T, oldObj T, mosaicsColumns []string) string {
	return logFieldChange(newObj, oldObj, false, mosaicsColumns)
}
