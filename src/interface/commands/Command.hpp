#ifndef COMMAND

#include <iostream>

class Command {
    private:
        bool replaceBoldTextInHelpMessage;
        std::string helpMessage;

    protected:
        void setHelpMessage(std::string message);
        std::string getHelpMessage();

    public:
        Command();
        void displayHelpMessage();
};

#endif

