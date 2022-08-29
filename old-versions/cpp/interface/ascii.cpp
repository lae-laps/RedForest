// ascii.cpp

#include "ascii.hpp"

#include <string>
#include <sstream>
#include <iostream>
#include <boost/format.hpp>
#include <boost/algorithm/string.hpp>

using namespace std;
using namespace boost::algorithm;

void ascii::print(string text) {
    // \[\033[38;5;xm\] text \[\033[m\] 
    text = replace_all_copy(text, "[END]", "\033[m");
    text = replace_all_copy(text, "[RED]", "\033[38;5;1m");
    text = replace_all_copy(text, "[GREEN]", "\033[38;5;2m");
    text = replace_all_copy(text, "[YELLOW]", "\033[38;5;3m");
    text = replace_all_copy(text, "[BLUE]", "\033[38;5;4m");
    text = replace_all_copy(text, "[PURPLE]", "\033[38;5;5m");
    cout << text;
}

void ascii::printUserError(string err) {
    ascii::print("[YELLOW][-][END] " + err + "\n");
}

void ascii::printRuntimeError(string err) {
    ascii::print("[RED][!][END] " + err + "\n");
}

void ascii::printInfo(string info) {
    ascii::print("[BLUE][*][END] " + info + "\n");
}

void ascii::clearScreen() {
    try {
        cout << "\033[2J\033[1;1H";
    } catch(...) {
        system("clear");
        // printRuntimeError("Error clearing the screen normally");                     // DEBUG
    }
}

