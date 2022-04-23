// interface.go
package main

import (
    "os"
    "fmt"
    "strings"
)

func banner() {
    Printf(`
 ██▀███  ▓█████ ▓█████▄   █████▒▒█████   ██▀███  ▓█████   ██████ ▄▄▄█████▓
▓██ ▒ ██▒▓█   ▀ ▒██▀ ██▌▓██   ▒▒██▒  ██▒▓██ ▒ ██▒▓█   ▀ ▒██    ▒ ▓  ██▒ ▓▒
▓██ ░▄█ ▒▒███   ░██   █▌▒████ ░▒██░  ██▒▓██ ░▄█ ▒▒███   ░ ▓██▄   ▒ ▓██░ ▒░
▒██▀▀█▄  ▒▓█  ▄ ░▓█▄   ▌░▓█▒  ░▒██   ██░▒██▀▀█▄  ▒▓█  ▄   ▒   ██▒░ ▓██▓ ░
░██▓ ▒██▒░▒████▒░▒████▓ ░▒█░   ░ ████▓▒░░██▓ ▒██▒░▒████▒▒██████▒▒  ▒██▒ ░
░ ▒▓ ░▒▓░░░ ▒░ ░ ▒▒▓  ▒  ▒ ░   ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░░░ ▒░ ░▒ ▒▓▒ ▒ ░  ▒ ░░
  ░▒ ░ ▒░ ░ ░  ░ ░ ▒  ▒  ░       ░ ▒ ▒░   ░▒ ░ ▒░ ░ ░  ░░ ░▒  ░ ░    ░
  ░░   ░    ░    ░ ░  ░  ░ ░   ░ ░ ░ ▒    ░░   ░    ░   ░  ░  ░    ░
   ░        ░  ░   ░               ░ ░     ░        ░  ░      ░
                 ░
                           >>-------------------<<
                            created by: lae-laps
                           >>-------------------<<
`, 196)

}

func help(sliced []string) {
    if len(sliced) < 2 {
    PrintBold(`
USAGE OF REDFOREST
==================`, 8)
    Printf(`
help ....... shows this help message
clear ...... clears the screen
exit ....... exits the tool
list ....... lists the available tools
version .... shows the current version of the tool

`, 8)
    } else {
        var helpMessage string
        for i, content := range sliced {
            if i > 0 {
                switch content {
                    case "clear":
                        helpMessage = "Clears the screen\n"
                    case "dnslookup":
                        helpMessage = "dnslookup help here\n"
                    default:
                        helpMessage = "could not find help for this command\n"
                }
                header := "\nHelp for " + content + "\n=========" 
                for i := 0; i < len(content); i++ {
                    header += "="
                }
                PrintBold(header, 8)
                Printf(helpMessage, 8)
            }
        }
    fmt.Print("\n")
    }
}

func parseExpression (expression string) {
    if len(expression) < 2 {
        Printf("[#] ~ Cannot validate an empty string\n", 9)
    } else {
        sliced := strings.Fields(expression)
        switch sliced[0] {
            case "exit":
                Printf("Exiting...\n", 8)
                os.Exit(0)
            case "help":
                help(sliced)
            case "clear":
                ClearScreen()
            case "dnslookup":
                DnsLookupCaller(expression)
            case "banner":
                banner()
            default:
                Printf("[#] ~ Invalid expression -> use help to view usage of tool\n", 9)
        }
    }
}

func DynamicInterface() {
//    banner()
    for {
        expression := Input("redforest > ", 197)
        parseExpression(expression)
    }
}
