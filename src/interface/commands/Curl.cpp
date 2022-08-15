// Curl.cpp

#include "Curl.hpp"
#include "../ascii.hpp"
#include "../../modules/HTTPClient/HttpClient.hpp"

void Curl::run(string expression) {
    bool valid = parseExpression(expression);
    if (not valid) {
        ascii::printUserError("Invalid target");
        return;
    }

    

}

bool Curl::parseExpression(string expression) {
    if () {
        return true;
    }
    return false;
}


