#ifndef INTERFACE
#define INTERFACE

#include <iostream>

using namespace std;

class Interface {
    public:
        Interface();
        void run();
        void banner();
        void setShowBannerOnStartup(bool);
        bool getShowBannerOnStartup();
        void setUseAnsiEscapes(bool);
        bool getUseAnsiEscapes();
        void setPromtp(string);
        string getPrompt();
       
        void parse(string); 

    private:
        bool showBannerOnStartup;
        bool useAnsiEscapes;
        string prompt;
};

#endif

