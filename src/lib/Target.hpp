#ifndef TARGET
#define TARGET

#include <iostream>

using namespace std;

class Target {
    public:
        Target();
        int setTarget(string);
        string getTarget();
        void setTargetValidity(int);
        int getTargetValidity();
        void setTargetType(int);
        int getTargetType();

    private:
        string target;
        unsigned short targetType;
        unsigned short targetValidity;

/* TARGET VALIDITY TABLE                TARGET TYPE TABLE
  -------------------------             -------------------    
  | 0 | valid ipv4 / ipv6 |             | 0 | ipv4 / ipv6 |
  | 1 | valid domain name |             | 1 | domain name |
  | 2 | invalid           |             | 2 | http(s)     |
  -------------------------             -------------------
*/

};

#endif

