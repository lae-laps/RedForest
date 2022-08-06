#ifndef MAIN_INTERFACE

#include <iostream>

using namespace std;

class MainInterface {
    public:
        MainInterface();
        void run();
        void banner();
        void setShowBannerOnStartup(bool);
        bool getShowBannerOnStartup();
        void setUseAnsiEscapes(bool);
        bool getUseAnsiEscapes();
        void setPromtp(string);
        string getPrompt();
        void setExpression(string);
        string getExpression();
       
        void parse(); 

    private:
        bool showBannerOnStartup;
        bool useAnsiEscapes;
        string prompt;
        string expression;
};

#endif

