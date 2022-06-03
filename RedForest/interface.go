// interface.go

/*
 * In the event of a new module being added either by the original devs or by contributors to
 * RedForest, an entry should be added here, in the corresponding menu. Any further CLI interfaces
 * for specific tools should not be included here. This file should not be extended beyond the
 * point of including launchers for each module or new sections -> This allows for all the modules
 * to be decentralized and the modules to be written each in their own language. It is recomended
 * to separe your own modules into files and to have a main or interface file to call with your
 * launcher here.
 */

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
        return
    }
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
            DnsLookupCaller(sliced)
        case "banner":
            banner()
        case "map":
            PortScannerCaller(sliced)
        case "webscan":
//            WebScanCaller(sliced) // TODO: Implement this 
        case "usearch": // user searcher
            userSearcherCaller(sliced)
        default:
            Printf("[#] ~ Invalid expression -> use help to view usage of tool\n", 9)
    }
}

func CommandInterface() {
//    banner()
    for {
        expression := Input("redforest > ", 197)
        parseExpression(expression)
    }
}
