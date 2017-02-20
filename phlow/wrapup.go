package phlow

import "fmt"

//WrapUp ...
func WrapUp() {
	//Before check - status

	if err := Add(); err != nil {
		fmt.Println("project files could not be added" + err.Error())
		return
	}

	info, _ := Branch("current")

	if output, err := Commit("Closes#" + info.current); err == nil {
		fmt.Println(output)
		return
	}

}
