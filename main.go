package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gookit/color/colorp"
	"github.com/manifoldco/promptui"
)

func main() {
	if _, err := os.Stat("./vite.config.js"); err == nil {
		if _, err := os.Stat("./package.json"); err == nil {

			prompt := promptui.Select{
				Label: "Select your package manager:",
				Items: []string{"npm", "yarn", "pnpm", "bun"},
			}

			_, selectedPackge, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			stdout, err := exec.Command("node", "-v").Output()
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			colorp.Bluef("your are using node->%s and %s.\n", strings.TrimSpace(string(stdout)), selectedPackge)
			colorp.Grayln("Installing packages...")
			stdou, _ := exec.Command(selectedPackge, "install -D tailwindcss postcss autoprefixer").Output()
			fmt.Println(string(stdou))
		} else {
			colorp.Redp("Couldn't find package.json.")
		}
	} else {
		colorp.Redp("This is not vite project!")
	}

}
