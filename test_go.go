package test_mod_go

import "fmt"

func TestGo() string {
	some_var := "something2"
	return fmt.Sprintf("Print something, %s", some_var)
}
