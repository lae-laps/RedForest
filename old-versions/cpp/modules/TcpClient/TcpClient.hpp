#ifndef TCP_CLIENT
#define TCP_CLIENT

#include <iostream>

using namespace std;

class TcpClient {
    private:
        int bufferLen;
        int port;
        string host;
        
    public:
        TcpClient();
        void setBufferLen(int);
        int getBufferLen();
        void setHost(string);
        string getHost();
        void setPort(int);
        int getPort();
       
        int send(string);

};

#endif

