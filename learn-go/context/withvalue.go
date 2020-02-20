package main

/*
func WithValue(parent Context, key, val interface{}) Context

*/


import (
	"fmt"
	"context"
)

func fromContext(ctx context.Context, k interface{}) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("fount value: ", v)
		return
	}
	fmt.Println("k not fount: ", k)
}

func withvalue() {

	type favContextKey string


	language := favContextKey("language")
	color := favContextKey("color")

	root := context.Background()

	ctx := context.WithValue(root, language, "Go")
	ctx = context.WithValue(ctx, color, "Read")

	fromContext(ctx, language)
	fromContext(ctx, color)


}
/*
func withvalue2() {

	//1. 往context中存储的数据类型
	type User struct{
		name string
		age int
	}

	//2. User 绑定的key
	type userkey string

	//3. 给key赋值
	k := userkey("user")


	var user := User{"xiaoming",18}

	ctx := context.WithValue(context.Background(), k, &user)


}
*/
func main() {
	withvalue()
}
