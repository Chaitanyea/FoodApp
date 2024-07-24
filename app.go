package main

import "GoSQLReactProject/src/util"

func main() {
	x := util.GetDb()
	x.Get("")
}
