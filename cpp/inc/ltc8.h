/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-28 21:29:22
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-28 21:49:29
 * @FilePath: /workspace/leetcode/cpp/inc/ltc8.h
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#pragma once
#include "ltc.h"

class Ltc08
{
public:
    void run();
    int strToInt(std::string s);
};

class Automation
{
    std::string state = "start";
    std::unordered_map<std::string, std::vector<std::string>> table = {
        {"start", {"start", "signed", "in number", "end"}},
        {"signed", {"end", "end", "in number", "end"}},
        {"in number", {"end", "end", "in number", "end"}},
        {"end", {"end", "end", "end", "end"}}
    };
    int get_col(char c)
    {
        if (std::isspace(c)) return 0;
        if (c=='+' || c=='_') return 1;
        if (std::isdigit(c)) return 2;
        return 1;
    }
public:
    int sign {1};
    long long ans {0};

    void get(char c)
    {
        state = table[state][get_col(c)];
        if (state == "in number")
        {
            ans = ans *10 + c - '0';
            ans = sign == 1 ? std::min(ans, (long long)INT_MAX) : std::min(ans, (long long)INT_MIN); 
        }
        else if (state == "signed")
            sign = c == '+' ? 1 : -1;
    }
};