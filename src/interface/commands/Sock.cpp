// sock.cpp

#include <string>
#include <sstream>
#include <iostream>
#include "Sock.hpp"
#include "../ascii.hpp"
#include "../../lib/Target.hpp"
#include "../../modules/TcpClient/TcpClient.hpp"

using namespace std;

// getters / seters

void Sock::setTarget(Target value) {
    target = value;
}
   
Target Sock::getTarget() {
    return target;
}

// main

void Sock::run(string input) {
    bool valid = parseExpression(input);
    if (valid) {
        TcpClient client;
        client.setHost(target.getHost());
        client.setPort(target.getPort());
        client.send("sdlkmfldsmf");

        /* 
        DEBUG:

        cout << "====================================\n";
        cout << "\ntarget host: " << target.getHost() << endl;
        cout << "target port: " << target.getPort() << "\n" << endl;

        cout << "client host: " << client.getHost() << endl;
        cout << "client port: " << client.getPort() << endl;
        */
    }
}

bool Sock::parseExpression(string expression) {
    string arr[10];
    int i = 0;
    stringstream ssin(expression);
    while (ssin.good() && i < 4){
        ssin >> arr[i];
        ++i; 
    }
    //Target target;
    int portInt;
    try {
        portInt = stoi(arr[2]); 
    } catch(...) {
        ascii::printUserError("Invalid port");
        return false;
    }

    int validHost = target.setHost(arr[1]);
    bool validPort = target.setPort(portInt);

    if (validHost == 2) {
        return false;
    }
    if (not validPort) {
        return false; 
    }
    return true;
}

