package handlers

import (
	"bufio"
	"fmt"
	"hostify/io"
	"log"
	"os"
)

var welcome string = `
 __    __                        __      __   ______                   ______   __        ______
/  |  /  |                      /  |    /  | /      \                 /      \ /  |      /      |
$$ |  $$ |  ______    _______  _$$ |_   $$/ /$$$$$$  |__    __       /$$$$$$  |$$ |      $$$$$$/
$$ |__$$ | /      \  /       |/ $$   |  /  |$$ |_ $$//  |  /  |      $$ |  $$/ $$ |        $$ |
$$    $$ |/$$$$$$  |/$$$$$$$/ $$$$$$/   $$ |$$   |   $$ |  $$ |      $$ |      $$ |        $$ |
$$$$$$$$ |$$ |  $$ |$$      \   $$ | __ $$ |$$$$/    $$ |  $$ |      $$ |   __ $$ |        $$ |
$$ |  $$ |$$ \__$$ | $$$$$$  |  $$ |/  |$$ |$$ |     $$ \__$$ |      $$ \__/  |$$ |_____  _$$ |_
$$ |  $$ |$$    $$/ /     $$/   $$  $$/ $$ |$$ |     $$    $$ |      $$    $$/ $$       |/ $$   |
$$/   $$/  $$$$$$/  $$$$$$$/     $$$$/  $$/ $$/       $$$$$$$ |       $$$$$$/  $$$$$$$$/ $$$$$$/
                                                     /  \__$$ |
                                                     $$    $$/
                                                      $$$$$$/
`

// LinkKeyHandler connect your token key to the cli
func LinkKeyHandler() {

	hostifyPath := TokenPath()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%v \n enter your token key: ", io.Green+welcome+io.Reset)
	scanner.Scan()
	input := scanner.Text()

	createFile, createError := os.Create(hostifyPath)

	if createError != nil {
		io.ErrorMessage("creating hostify.key file")
		log.Fatal(createError)
	}

	bitesWriter, errorWrite := createFile.WriteString(input)

	if errorWrite == nil {
		// * show done message
		createFile.Close()
		done := fmt.Sprintf("%v bytes written", bitesWriter)
		io.SuccessMessage(done)
	} else {
		log.Fatal(errorWrite)
	}
}
