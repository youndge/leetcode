/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-30 22:33:42
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-30 22:43:06
 * @FilePath: /workspace/leetcode/cpp/src/ltc13.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc13.h"

void Ltc13::run()
{
    std::string v1 {"MCMXCIV"}; //1994
    std::cout<<"v1 is "<< RomanToNum(v1)<<std::endl;
    std::string v2 {"LVIII"}; //58
    std::cout<<"v2 is "<< RomanToNum(v2)<<std::endl;
    std::string v3 {"IV"}; //4
    std::cout<<"v3 is "<< RomanToNum(v3)<<std::endl;
}

int Ltc13::RomanToNum(const std::string roman)
{
    const std::unordered_map<char, int> symbolValue = {
        {'I',1},
        {'V',5},
        {'X',10},
        {'L',50},
        {'C',100},
        {'D',500},
        {'M',1000},
    };
    int num {0};
    for (auto i =0; i<roman.length();++i)
    {
        int value = symbolValue.at(roman.at(i));
        if (i<roman.length()-1&&symbolValue.at(roman.at(i))<symbolValue.at(roman.at(i+1)))
        {
            num -= value;
        }
        else
        {
            num += value;
        }
    }
    return num;
}