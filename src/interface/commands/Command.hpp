#ifndef COMMAND
#define COMMAND

#include <iostream>

using namespace std;

class Command {
    private:
        bool replaceBoldTextInHelpMessage;
        string helpMessage;

    protected:
        void setHelpMessage(string message);
        string getHelpMessage();

    public:
        Command();
        void run();
        void displayHelpMessage();
        int parseExpression(string);     // polymorphic method 
};

#endif

