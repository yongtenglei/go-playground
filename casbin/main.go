package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func KeyMatch(key1 string, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}
	return key1 == key2[:i]
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}
func main() {
	a, err := gormadapter.NewAdapter("mysql", "root:1@tcp(127.0.0.1:3306)/casbin", true)
	if err != nil {
		fmt.Println(err)
	}
	e, err := casbin.NewEnforcer("./model.conf", a)
	if err != nil {
		fmt.Println(err)
	}
	e.AddFunction("my_func", KeyMatchFunc)
	// Load the policy from DB.
	err = e.LoadPolicy()
	if err != nil {
		fmt.Println(err)
	}
	sub := "rey1"
	obj := "data1"
	act := "read"
	e.AddPolicy(sub, obj, act)
	//e.AddPolicy("charlotte", "data2", "write")
	//e.AddPolicy("charlotte1", "data3", "write")
	//e.AddPolicy("charlotte2", "data1", "write")
	//e.AddPolicy("charlotte3", "data5", "write")

	//e.RemoveFilteredPolicy(0, "charlotte")
	e.RemoveFilteredPolicy(0, "rey")

	fmt.Println(e.GetPolicy())
	ok, err := e.Enforce("rey*", obj, act)
	if err != nil {
		log.Fatal(err)
	}
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

	err = e.SavePolicy()
	if err != nil {
		fmt.Println(err)
	}

}
