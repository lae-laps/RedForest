// sock.cpp

#include <sstream>
#include <iostream>
#include "Sock.hpp"
#include "../ascii.hpp"
#include "../../lib/Target.hpp"

void Sock::run(string input) {
    parseExpression(input);
}

int Sock::parseExpression(string expression) {
    string arr[10];
    int i = 0;
    stringstream ssin(expression);
    while (ssin.good() && i < 4){
        ssin >> arr[i];
        ++i; 
    }
    
}

