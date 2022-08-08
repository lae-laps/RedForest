#ifndef SOCK
#define SOCK

#include "Command.hpp"
#include "../../lib/Target.hpp"

class Sock : public Command {
    public:
        void run(string);
        int parseExpression(string);
};

#endif

