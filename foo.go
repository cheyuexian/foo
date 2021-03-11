package foo

import "fmt"
func Greet( name string ) string  {
	return fmt.Sprintf("%s,你好 v0.1.0",name)
}