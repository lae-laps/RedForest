#include <iostream>
#include <asio.hpp>
#include <asio/ip/address.hpp>
#include <asio/ip/tcp.hpp>
#include <asio/ts/buffer.hpp>
#include <asio/ts/internet.hpp>

using namespace std;

int main() {
   
    asio::error_code ec;            // error handler

    asio::io_context context;       // unique instance of ASIO - hides platform specific requirements

    // an endpoint is something ASIO can connect TO
    // here we create a TCP style endpoint
    // the make_address() function converts a string version of an IP into ASIO ip
    // we also pass the error code, and if there is an error, the ec will contain the error
    asio::ip::tcp::endpoint endpoint(asio::ip::make_address("93.184.216.34", ec), 80);

    // we create a NETWORKING socket
    // we asociate it to our ASIO object

    asio::ip::tcp::socket socket(context);

    // we tell our socket object to connect to our endpoint
    socket.connect(endpoint, ec);

    // here we check the error code
    
    if (!ec) {              // no error code
        cout << "Succesfull Connection" << endl;
    } else {
        // the message() method in our error code returns a string with our error
        cout << "error: " << ec.message() << endl;  
    }

    // is_open() getter method in sockets to determine if the socket is actively open
    if (socket.is_open()) {

        // we create a raw string containing an HTTP request
        string httpRequest = 
            "GET /index.html HTTP/1.1\n"
            "Host: example.com\n"
            "Connection: close\n\n";

        // we write to the socket via a asio::buffer()
        // asio buffers take the content and lenght as shown below
        socket.write_some(asio::buffer(httpRequest.data(), httpRequest.size()), ec);

        socket.wait(socket.wait_read);

        size_t bytes = socket.available();
        cout << "Bytes avaliable: " << bytes << endl;

        if (bytes > 0) {
            std::vector<char> vBuffer(bytes);
            socket.read_some(asio::buffer(vBuffer.data(), vBuffer.size()), ec);
 
            for (auto c : vBuffer) {
                cout << c;
            }
        }
    }
    return 0;
}

