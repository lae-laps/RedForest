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
        ascii::printUserError("Invalid target");
        return;
    }

    string path;

    HttpClient client;
    client.connection(path, target.getHost(), target.getPort());

}

bool Curl::parseExpression(string expression) {
    
    string arr[10];
    int i = 0;
    stringstream ssin(expression);
    while (ssin.good() && i < 4) {
        ssin >> arr[i];
        i++;
    }

    bool https = starts_with(arr[1], "https://"); 

    if (not https) {
        target.setPort(80);                                                     // set target's port to 80 if the target doesn't start with https://
    } else {
        target.setPort(443);                                                    // set target's port to 443 otherwise
    }

    if () {
        return true;
    }
    return false;
}


