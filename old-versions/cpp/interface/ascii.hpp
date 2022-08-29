#ifndef ASCII
#define ASCII

#include <iostream>

using namespace std;

class ascii {
    public:
        static void clearScreen();
        static void print(string);
        static void printInfo(string);
        static void printUserError(string);
        static void printRuntimeError(string);
};

#endif

