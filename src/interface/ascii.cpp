// ascii.cpp

#include <iostream>
#include "ascii.hpp"

using namespace std;

void ascii::print(string text) {
    cout << text << endl;
}

void ascii::printUserError(string err) {
    ascii::print("[RED][!][END] ");
}

