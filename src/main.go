package main

import (
	"app/view"
)

func main() {
	for {
		menu := view.PrintMenu()
		navigate(menu)
	}
}

func navigate(menu int64) {
	switch menu {
	case 1:
		doInsert()
	case 2:
		doGetIdList()
	case 3:
		doGetData()
	case 4:
		view.PrintToExit()
	default:
		view.PrintWrongInput()
	}

}

func doInsert() {

}

func doGetIdList() {

}

func doGetData() {

}
