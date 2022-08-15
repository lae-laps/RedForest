#include <iostream>
#include <boost/regex.hpp>
#include <boost/algorithm/string.hpp>

#include "Target.hpp"
#include "../interface/ascii.hpp"

using namespace std;
using namespace boost::algorithm;

Target::Target() {
}

bool Target::setPort(int value) {
    //try {
   
    cout << "port: " << value << "\n";

    if (value <= 65535 && value >= 0) { 
        port = value;
    } else {
        ascii::printUserError("Invalid port");
        return false;
    }
    return true;
    
    /*} catch(...) {
        ascii::printUserError("Invalid port"); 
        return false;
    }*/
}

int Target::getPort() {
    return port;
}

int Target::setHost(string value) {

    boost::regex ipv4{"^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)(\\.(?!$)|$)){4}$"};
    boost::regex ipv6{"(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))"};
    boost::regex domain{"^(((?!\\-))(xn\\-\\-)?[a-z0-9\\-_]{0,61}[a-z0-9]{1,1}\\.)*(xn\\-\\-)?([a-z0-9\\-]{1,61}|[a-z0-9\\-]{1,30})\\.[a-z]{2,}$"};

    int return_value;

    if (boost::regex_match(value, ipv4) || boost::regex_match(value, ipv6)) {
        return_value = 0;
    } else if (boost::regex_match(value, domain)) {
        return_value = 1;
    } else {
        return_value = 2;
    }
    if (return_value == 2) {
        return_value = 3;
        ascii::printUserError("Invalid target");
    }
    setTargetValidity(return_value);
    return getTargetValidity();
}

string Target::getHost() {
    return host;
}

void Target::setTargetValidity(int value) {
    targetValidity = value;
}

int Target::getTargetValidity() {
    return targetValidity;
}

void Target::setTargetType(int value) {
    targetType = value;
}

int Target::getTargetType() {
    return targetType;
}

