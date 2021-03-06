package strings

import (
	"github.com/yuin/gopher-lua"
	gostrings "strings"
)

func split(L *lua.LState) int {

	s := L.CheckString(1)
	sep := L.CheckString(2)
	as := gostrings.Split(s, sep)

	var rvStr []string
	for _, v := range as {
		if len(gostrings.TrimSpace(v)) == 0 {
			continue
		}
		rvStr = append(rvStr, v)
	}

	arr := L.CreateTable(len(as), 0)
	for _, v := range rvStr {
		arr.Append(lua.LString(v))
	}

	L.Push(arr)
	return 1
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		"split": split,
	})

	L.Push(mod)
	return 1
}
