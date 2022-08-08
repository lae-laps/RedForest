// TcpClient.cpp

#include <iostream>
#include <boost/asio.hpp>
#include <boost/array.hpp>

#include "../../interface/ascii.hpp"
#include "TcpClient.hpp"

using namespace std;
using boost::asio::ip::tcp;

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
        boost::asio::io_service ios;                    // create object & endpoint
        boost::asio::ip::tcp::endpoint endpoint(boost::asio::ip::address::from_string(host), port);
        boost::asio::ip::tcp::socket socket(ios);       // initialize socket towards our object
        socket.connect(endpoint);                       // connect to endpoint

        boost::array<char, 1024> buf;
            std::copy(message.begin(),message.end(),buf.begin());
        boost::system::error_code error;
        socket.write_some(boost::asio::buffer(buf, message.size()), error);
            socket.close();
        return 0;
    } catch (...) {
        ascii::printRuntimeError("Undefined error while connecting to host");
        return 1;
    }
}

