// ANSI-utils package for CLI utils

/* Methods description
    * print -> print with ANSI escape code color -> ( content string, color int ) 
    * printBackgound -> print with ANSI escape code background ( content string, backgroundColor int )
    * printBold -> print with ANSI escape code color but in bold ( content string, color int )
 */

// ansi.go

/*
TODO: Use runtime.GOOS to check platform and use correct escapes
*/

package main 

import (
    "os"
    "fmt"    
    "time"
    "bufio"
    "strings"
    "strconv"
//    "golang.org/x/term"
)

type Tty struct {
        state *term.State
        line  string
}

func readPassword(tty chan Tty, fd int) {
        state, _ := term.MakeRaw(fd)
        t := term.NewTerminal(os.Stdin, "")
        f := func(s string, i int, r rune) (string, int, bool) {
                //"•"
                fmt.Print("*")
                return s, i, false
        }

        t.AutoCompleteCallback = f
        line, _ := t.ReadPassword("")
        tty <- Tty{state: state, line: line}
}

func inputPassword() string {
        tty := make(chan Tty)
        fd := int(os.Stdin.Fd())
        go readPassword(tty, fd)
        t := <-tty
        close(tty)
        term.Restore(fd, t.state)
        password := t.line
        return password
}

func input(prompt string, color int) string {
    print(prompt, color)
         
    in := bufio.NewReader(os.Stdin)                                                                                                                                                                          
    line, err := in.ReadString('\n')                                                                                                                                                                         
    if err != nil {                                                                                                                                                                                          
        printRuntimeError("Error reading line")
    }
    line = strings.Replace(line, "\n", "", -1) // strip chars from line
    return line
}

func printUserError(content string) {
    print("[!] ", 166)
    print(content, 7)
}

func printRuntimeError(content string) {
    print("[#] ", 1)
    print(content, 7)
}

func print(content string, color int) {
    colorStr := strconv.Itoa(color)
    escape := "\033[38;5;" + colorStr + "m" + content + "\033[m"
    fmt.Print(escape)
}

func printInfo(info string) {
    print("[+] ", 10)
    print(info, 7)
}

func printPrompt(prompt string) {
    print("[>] ", 4)
    print(prompt, 7)
}

func printBold(content string, color int) {
    colorStr := strconv.Itoa(color)
    escape := "\u001b[1m\033[38;5;" + colorStr + "m" + content + "\033[m\u001b[0m"
    fmt.Printf(escape)
}

func printBackground(content string, backgroundColor int) {
    colorStr := strconv.Itoa(backgroundColor)
    escape := "\033[48;5;" + colorStr + "m" + content + "\033[m"
    fmt.Println(escape)
}

func clearScreen() {
    fmt.Print("\u001b[2J")
    fmt.Printf("\u001b[1;1f")
}

// idk what this is
func chargingAnimation(interval int, color int) {
    for i := 0; i < 100; i++ {
        counter := strconv.Itoa(i + 1) 
        print("\u001b[1000D" + counter + "%", color)
        time.Sleep(time.Duration(interval) * time.Millisecond)
    }
}

