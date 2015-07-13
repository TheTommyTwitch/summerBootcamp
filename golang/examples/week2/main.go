package main

import (
	"examples/week2/hello"
	"examples/week2/converters"
	"fmt"
	"github.com/ttacon/chalk"
)


func main() {
	hello.Hello()

	KM := converters.MilesToKM(43.2)
	fmt.Println(chalk.Red, KM, chalk.ResetColor)


}
