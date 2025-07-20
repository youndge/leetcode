/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-28 21:42:45
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-28 21:48:18
 * @FilePath: /workspace/leetcode/cpp/src/ltc8.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc8.h"

void Ltc08::run()
{
    std::string s {"-1233ttt"};
    std::cout << "-123ttt -> " << strToInt(s) << std::endl;

    std::string s1 {"fffttt"};
    std::cout << "fffttt -> " << strToInt(s1) << std::endl;

    std::string s2 {"023ttt"};
    std::cout << "023ttt -> " << strToInt(s2) << std::endl;

    std::string s3 {"0tt34t"};
    std::cout << "0tt34t -> " << strToInt(s3) << std::endl;
}

int Ltc08::strToInt(std::string s)
{
    Automation automation;
    for (char c : s)
    {
        automation.get(c);
    }
    return automation.sign * automation.ans;
}