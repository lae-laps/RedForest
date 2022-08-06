#include <string>
#include <sstream>
#include <iostream>

#include "MainInterface.hpp"
#include "commands/Sock.hpp"

using namespace std;

// getters / setters

bool MainInterface::getShowBannerOnStartup() {
    return showBannerOnStartup;
}

void MainInterface::setShowBannerOnStartup(bool value) {
    showBannerOnStartup = value;
}

bool MainInterface::getUseAnsiEscapes() {
    return useAnsiEscapes;
}

void MainInterface::setUseAnsiEscapes(bool value) {
    useAnsiEscapes = value;
}

void MainInterface::setPromtp(string value) {
    prompt = value;
}

string MainInterface::getPrompt() {
    return prompt;
}

void MainInterface::setExpression(string value) {
    expression = value;
}

string MainInterface::getExpression() {
    return expression;
}

// constructor

MainInterface::MainInterface() {
    setPromtp(">> ");
    setUseAnsiEscapes(true);
    setShowBannerOnStartup(false);
}

// main

void MainInterface::banner() {
    cout << "REDFOREST" << endl; 
}

void MainInterface::parse() {

    string arr[10];
    int i = 0;
    stringstream ssin(getExpression());
    while (ssin.good() && i < 4){
        ssin >> arr[i];
        ++i;
    }
    
    if (arr[0] == "sock") {
        Sock sock;
        int isValid = sock.setTarget(arr[1]);
        cout << isValid << endl;
        sock.connection(); 

    } else if (arr[0] == "exit") {
        exit(0);
    }
}

void MainInterface::run() {
    if (showBannerOnStartup) {
        banner();
    }
    for (;;) {
        cout << prompt; 

        getline(cin, expression);

        parse();
    }
}

