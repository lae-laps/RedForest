#ifndef TARGET

#include <iostream>

using namespace std;

class Target {
    public:
        Target();
        int setTarget(string);
        string getTarget();
        void setTargetValidity(int);
        int getTargetValidity();

    private:
        string target;
        int targetValidity;

/* TARGET VALIDITY TABLE
  -------------------------
  | 0 | valid ipv4        |
  | 1 | valid ipv6        |
  | 2 | valid domain name |
  | 3 | invalid           |
  -------------------------
 */

};

#endif

