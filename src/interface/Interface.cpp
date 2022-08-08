#include <string>
#include <sstream>
#include <iostream>
#include "ascii.hpp"
#include "../lib/Target.hpp"
#include "Interface.hpp"
#include "commands/Sock.hpp"

using namespace std;

// getters / setters

bool Interface::getShowBannerOnStartup() {
    return showBannerOnStartup;
}

void Interface::setShowBannerOnStartup(bool value) {
    showBannerOnStartup = value;
}

bool Interface::getUseAnsiEscapes() {
    return useAnsiEscapes;
}

void Interface::setUseAnsiEscapes(bool value) {
    useAnsiEscapes = value;
}

void Interface::setPromtp(string value) {
    prompt = value;
}

string Interface::getPrompt() {
    return prompt;
}

// constructor

Interface::Interface() {
    setPromtp("[PURPLE]>> [END]");
    setUseAnsiEscapes(true);
    setShowBannerOnStartup(false);
}

// main

void Interface::banner() {
    cout << "REDFOREST" << endl; 
}

void Interface::parse(string expression) {

    string arr[10];
    int i = 0;
    stringstream ssin(expression);
    while (ssin.good() && i < 4){
        ssin >> arr[i];
        ++i;
    }
    
    if (arr[0] == "sock") {
        Sock sock;
        sock.run(expression); 

    } else if (arr[0] == "exit") {
        exit(0);
    } else {
        ascii::printUserError("Invalid expression");
    }
}

void Interface::run() {
    if (showBannerOnStartup) {
        banner();
    }
    for (;;) {
        string expression;
        
        ascii::print(prompt); 

        getline(cin, expression);

        parse(expression);
    }
}
