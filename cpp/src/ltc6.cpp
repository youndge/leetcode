/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-07-01 21:50:16
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-07-01 22:06:50
 * @FilePath: /workspace/leetcode/cpp/src/ltc6.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc6.h"

void Ltc6::run()
{
    std::string s {"PAYPALISHIRING"}; int numRows {4}; //"PINALSIGYAHRPI"
    std::cout<<"ZshapeChange: "<<ZshapeChange(s,numRows)<<std::endl;
}

std::string Ltc6::ZshapeChange(std::string s, int n)
{
    std::string res {""};
    auto convert = [=,&res] () {
        int t = n*2-2;
        for (int i = 0; i < n;++i)
        {
            for (int j = 0; j + i < s.length();j+=t)
            {
                res +=s.at(j+i);
                if (i>0&&i<n-1&&j+t-i<s.length())
                {
                    res+= s.at(j+t-i);
                }
            }
        }
        return res;
    };
    return (n==1||n>=s.length())?s:convert();
}