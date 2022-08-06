// main.go

/*

 +----------------------+
 |  RedForest toolkit   |
 | created by: lae-laps |
 +----------------------+

 */

package main

import (
    "os"
)

func main() {
    
    if len(os.Args) == 2 {
        if os.Args[1] == "--cli" {
            CommandInterface()
        } else if os.Args[1] == "--menu" {
            //MenuInterface()
        }
    } else {
        CommandInterface()
    }
}

