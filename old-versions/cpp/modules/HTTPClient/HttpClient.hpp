// HttpClient.cpp

#ifndef HTTP_CLIENT
#define HTTP_CLIENT

#include <iostream>
#include <asio/ip/tcp.hpp>

using namespace std;

class HttpClient {
    public:
        HttpClient();
        ~HttpClient();
        void connection(string, string, int);
    protected:
        void grabData(asio::ip::tcp::socket&);
        
};

#endif

