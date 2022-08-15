// Curl.hpp

#ifndef CURL
#define CURL

#include "Command.hpp"                      // inheritance from Command class
#include "../../lib/Target.hpp"

class Curl : public Command {
    public:
        void run(string);
        bool parseExpression(string);

    private:
        Target target;
};

#endif

