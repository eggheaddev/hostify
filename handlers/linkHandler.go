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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("%v \n enter your token key: ",  "\u001b[32m" + welcome)
	scanner.Scan()
	input := scanner.Text()
	usr, err := user.Current()

	if err != nil {
		log.Fatal("The path you looking for do not exist")
	}

	path := filepath.FromSlash(usr.HomeDir + "/hostify.key")
	file, _ := os.Create(path)

	bitesWriter, err := file.WriteString(input)

	if err == nil {
		file.Close()
		done := fmt.Sprintf("%v bites writes", bitesWriter)
		io.SuccessMessage(done)
		os.Exit(0)
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(input, file)
}
