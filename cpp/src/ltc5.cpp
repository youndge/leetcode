#include "ltc5.h"

void Ltc5::run()
{
    //std::string s {"cbbd"}; //bd
    std::string s {"babad"}; //bab
    std::cout << "huiwen: " << huiwen(s) << std::endl;
}

std::string Ltc5::huiwen(std::string s)
{
    std::string res {""}, temp {""};
    for (int i = 0; i < s.length(); ++i)
    {
        for (int j = i; j <s.length(); ++j)
        {
            temp += s.at(j);
            std::string tem {temp};
            std::reverse(tem.begin(), tem.end());
            if (temp == tem)
                res = res.length() > temp.length() ? res : temp;
        }
        temp = "";
    }
    return res;
}