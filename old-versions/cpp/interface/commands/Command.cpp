/* Command class - parent to all of the child command classes 
 * Contains global blueprint for the command classes
 */

#include <regex>
#include <string>
#include <iostream>

#include "Command.hpp"

void Command::setHelpMessage(std::string message) {
    helpMessage = message;
}

std::string Command::getHelpMessage() {
    return helpMessage;
}

Command::Command() {
    replaceBoldTextInHelpMessage = false;
    setHelpMessage("testing");
}

void Command::displayHelpMessage() {
    /*if (replaceBoldTextInHelpMessage) {
        helpMessage = std::regex_replace(helpMessage, std::regex("**"), ""); // TODO: Make this work
    } */
    std::cout << helpMessage << std::endl;
} 

int Command::parseExpression(string expression) {
    /* STRUCTURE *
    parse()
    if (invalid) {
        printErrorMessage();
        return 1;
    } else {
        execute();
        return 0;
    }
    */
    return 0;
}

void Command::run() {
    // run 
}

