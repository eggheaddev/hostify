package handlers

import (
	"bufio"
	"fmt"
	"hostify/io"
	"log"
	"os"
	"os/user"
	"path/filepath"
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
	usr, errorGetpath := user.Current()

	hostifyPath := filepath.Join(usr.HomeDir, "hostify.key")

	_, errorFile := os.Stat(hostifyPath)

	fmt.Println(hostifyPath)

	// * if hostify.key exist show error
	if errorFile == nil {

		io.ErrorMessage(fmt.Sprintf("hostify.key is ready exist in Path:\n%v", hostifyPath))
		os.Exit(1)
	} else {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("%v \n enter your token key: ", io.Green+welcome+io.Reset)
		scanner.Scan()
		input := scanner.Text()

		if errorGetpath != nil {
			log.Fatal("The path you looking for do not exist")
			os.Exit(1)
		}

		// path := filepath.FromSlash(hostifyPath)
		fmt.Println(hostifyPath)
		createFile, _ := os.Create(hostifyPath)

		bitesWriter, errorWrite := createFile.WriteString(input)

		if errorWrite == nil {

			createFile.Close()
			done := fmt.Sprintf("%v bytes written", bitesWriter)
			io.SuccessMessage(done)
			os.Exit(0)
		} else {

			fmt.Println(errorWrite)
			os.Exit(1)
		}
		fmt.Println(input, createFile)
	}
}
