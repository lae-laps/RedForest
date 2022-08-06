#ifndef SOCK
#define SOCK

#include "Command.hpp"
#include "../../lib/Target.hpp"

class Sock : public Command {
    public:
        void dialog(string);
};

#endif

