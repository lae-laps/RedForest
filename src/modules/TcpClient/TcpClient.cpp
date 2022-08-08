// TcpClient.cpp

#include <iostream>

#include "../../interface/ascii.hpp"
#include "TcpClient.hpp"

using namespace std;

TcpClient::TcpClient() {
    setBufferLen(1024);                                 // set buffer len standard
}

void TcpClient::setBufferLen(int value) {
    bufferLen = value;
}

int TcpClient::getBufferLen() {
    return bufferLen;
}

void TcpClient::setHost(string value) {
    host = value;
}

string TcpClient::getHost() {
    return host;
}

void TcpClient::setPort(int value) {
    port = value;
}

int TcpClient::getPort() {
    return port;
}

int TcpClient::send(string message) {
    try {
        return true;
    } catch (...) {
        return false;
    }
}

