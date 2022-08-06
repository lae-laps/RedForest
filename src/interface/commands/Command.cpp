/* Command class - parent to all of the child command classes 
 * Contains global blueprint for the command classes
 */

#include <regex>
#include <string>
#include <iostream>

class Command {
    private:
        bool replaceBoldTextInHelpMessage;
        std::string helpMessage;

    protected:
        void setHelpMessage(std::string message) {
            helpMessage = message;
        }

        std::string getHelpMessage() {
            return helpMessage;
        }

    public:
        Command() {
            replaceBoldTextInHelpMessage = false;
        }
    
        void displayHelpMessage() {
            if (replaceBoldTextInHelpMessage) {
                helpMessage = std::regex_replace(helpMessage, std::regex("**"), ""); // TODO: Make this work
            } 
            std::cout << helpMessage << std::endl;
        } 
};

