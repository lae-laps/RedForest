// Sock.hpp

#ifndef SOCK
#define SOCK

#include "Command.hpp"
#include "../../lib/Target.hpp"

class Sock : public Command {
    public:
        void run(string);
        bool parseExpression(string);
        void setTarget(Target);
        Target getTarget();
        void setMessage(string);
        string getMessage();

    private:
        Target target;
        string message;
};

#endif

