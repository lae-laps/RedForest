// ANSI-utils package for CLI utils

/* Methods description
    * Println -> print with ANSI escape code color -> ( content string, color int ) 
    * PrintBackgound -> print with ANSI escape code background ( content string, backgroundColor int )
    * PrintBold -> print with ANSI escape code color but in bold ( content string, color int )
 */

// ansi.go
package main 

import (
    "os"
    "fmt"
    "time"
    "bufio"
    "strconv"
)

type options_menu []string

func (opt options_menu) show() {
    for i, content := range opt {
        i_str := strconv.Itoa(i)
        fmt.Print("[" + i_str + "] ~ " + content + "\n")
    }
}

func Input(prompt string, color int) string {
    Printf(prompt, color)
    in := bufio.NewReader(os.Stdin)
    line, err := in.ReadString('\n')
    if err != nil {
        Printf("Error reading line\n", 9)
    }
    return line
}

func Printf(content string, color int) {
    colorStr := strconv.Itoa(color)
    escape := "\033[38;5;" + colorStr + "m" + content + "\033[m"
    fmt.Printf(escape)
}

func PrintBold(content string, color int) {
    colorStr := strconv.Itoa(color)
    escape := "\u001b[1m\033[38;5;" + colorStr + "m" + content + "\033[m\u001b[0m"
    fmt.Println(escape)
}

func PrintBackground(content string, backgroundColor int) {
    colorStr := strconv.Itoa(backgroundColor)
    escape := "\033[48;5;" + colorStr + "m" + content + "\033[m"
    fmt.Println(escape)
}

func ClearScreen() {
    fmt.Print("\u001b[2J")
    fmt.Printf("\u001b[1;1f")
}

// idk what this is
func Charging() {
    for i := 0; i < 100; i++ {
        counter := strconv.Itoa(i + 1) 
        fmt.Print("\u001b[1000D" + counter + "%")
        time.Sleep(30 * time.Millisecond)
    }
}
