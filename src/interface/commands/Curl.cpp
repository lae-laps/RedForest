// Curl.cpp

#include <iostream>
#include "Curl.hpp"
#include "../ascii.hpp"
#include "../../modules/HTTPClient/HttpClient.hpp"
#include "../../lib/Target.hpp"

#include <boost/algorithm/string.hpp>

using namespace std;
using namespace boost::algorithm;

void Curl::run(string expression) {
    
    bool valid = parseExpression(expression);

    if (not valid) {
        return;
    }

    string path;

    HttpClient client;

    //cout << "host: " << target.getHost() << "\nport: " << target.getPort() << endl;

    client.connection(path, target.getHost(), target.getPort());
}

bool Curl::parseExpression(string expression) {

    string arr[10];                                                             // split the expression string into an array so that it is easier
    int i = 0;                                                                  // to parse it and separate the command line arguments
    stringstream ssin(expression);
    while (ssin.good() && i < 4) {
        ssin >> arr[i];
        i++;
    }

    bool https = starts_with(arr[1], "https://"); 

    // parse and set the port

    if (not https) {
        target.setPort(80);                                                     // set target's port to 80 if the target doesn't start with https://
    } else {
        target.setPort(443);                                                    // set target's port to 443 otherwise
    }
    
    // parse and set the address

    string address;

    try {
        address = replace_all_copy(address, "http://", "");
        address = replace_all_copy(address, "https://", "");
    } catch(...) {
        // no need to replace
    }

        vector<string> strs;
        boost::split(strs, address, boost::is_any_of("/"));

    int addressValidity = target.setHost(arr[1]); 

    if (addressValidity == 2) {
        return false;
    }
    return true;
}

