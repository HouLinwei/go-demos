package call_func_by_name

import (
	"fmt"
	"reflect"
)

var funcs map[string]interface{}

func init() {
	funcs = make(map[string]interface{})
	fs := []string{"sayHi"}
	for _, name := range fs {
		if _, ok := funcs[name]; !ok {
			funcs[name] = sayHi // fixme
		}
	}

}

func CallFuncByName(fName string) {
	params := []string{"Hou", "Zhao"}
	if v, ok := funcs[fName]; !ok {
		fmt.Println("not exist")
	} else {
		f := reflect.ValueOf(v)
		if f.Kind() == reflect.Func {
			var ps []reflect.Value
			for _, p := range params {
				ps = append(ps, reflect.ValueOf(p))
			}
			f.Call(ps)
		} else {
			fmt.Println(fName, "is not a callable function.")
		}
	}

}

func sayHi(names ...string) {
	for _, name := range names {
		fmt.Println("hello,", name)
	}
}

func allFuncsIn(pkgName string) {
	// read package file
	// get all exported functions
	// assign them to a local map

	// !!!failed!!!
	// It cannot work with a static compiled language.
}
