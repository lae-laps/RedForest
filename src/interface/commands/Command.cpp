/* Command class - parent to all of the child command classes 
 * Contains global blueprint for the command classes
 */

#include <regex>
#include <string>
#include <iostream>

#include "Command.h"

bool replaceBoldTextInHelpMessage;
std::string helpMessage;

void Command::setHelpMessage(std::string message) {
    helpMessage = message;
}

std::string Command::getHelpMessage() {
    return helpMessage;
}

Command::Command() {
    replaceBoldTextInHelpMessage = false;
}

void Command::displayHelpMessage() {

    //temportal pleeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaseeeeeeeeeeeeeeeeeeeeeeeeee
    setHelpMessage("test");


    if (replaceBoldTextInHelpMessage) {
        helpMessage = std::regex_replace(helpMessage, std::regex("**"), ""); // TODO: Make this work
    } 
    std::cout << helpMessage << std::endl;
} 

