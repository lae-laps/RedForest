// HttpClient.cpp

#include "HttpClient.hpp"

#include <iostream>
#include <asio.hpp>
#include <asio/ip/address.hpp>
#include <asio/ip/tcp.hpp>
#include <asio/ts/buffer.hpp>
#include <asio/ts/internet.hpp>

using namespace std;

vector<char> vBuffer(20 * 1024);

// this function uses a lambda function
// it implements a buffer with a fixed size, and recursively grabs a buffer of data
// only if the recieved data is not enough for one buffer
void grabData(asio::ip::tcp::socket& socket) {
    socket.async_read_some(asio::buffer(vBuffer.data(), vBuffer.size()),
        [&](error_code ec, size_t length) {
        // this block of code contains what should happen with the retrieved data
            if (!ec) {
                cout << "\n\nRead " << length << "bytes\n" << endl;
                for (int i = 0; i < length; i++) {
                    cout << vBuffer[i];
                }
                grabData(socket);
            }   
        }
    );
}

int main() {
   
    // we create an Error Handler ( Error Code ) ASIO variable which we will use throuought the program to check for errors.
    // When an exception happens, the error code object defined bellow, will adopt that exception so we can handle or ignore it.
    asio::error_code ec;    

    // unique instance of ASIO - hides platform specific requirements
    // this object is the main ASIO context object, and will be run in a separate thread most of the time.
    asio::io_context context;       
    // since we are running our ASIO context in a separate thread to not block the main thread, 
    // the run instruction used in the context is then instructed to end the context when it has no
    // more instructions to do. The problem with this is that by the time we actually give real work
    // to the context, the context thread has already ended. The following line is ASIO's way of giving
    // some "fake" work to the context to keep it running while we dont give it our real work. 
    asio::io_context::work idleWork(context);

    // here we run ASIO in a separate thread
    // This makes ASIO not block the main thread while performing ASYNC calls
    // the run function of the context runs the context and returns as soon as 
    // the context runs out of work to do
    thread thrContext = thread([&]() { context.run(); });
