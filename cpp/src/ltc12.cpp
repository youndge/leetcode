/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-14 11:42:10
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-14 12:03:14
 * @FilePath: /workspace/leetcode/cpp/src/ltc12.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc12.h"
#include <utility>

void Ltc12::run()
{
    int v1 {3749};
    std::cout << "v1" << "is " << transferToRoma(v1) << std::endl;
    int v2 {58};
    std::cout << "v2" << "is " << transferToRoma(v2) << std::endl;
    int v3 {1994};
    std::cout << "v3" << "is " << transferToRoma(v3) << std::endl;
}

std::string Ltc12::transferToRoma(const int number)
{
    const std::pair<int, std::string> valueSymbols[] = {
        {1000, "M"}, {900, "CM"}, {500, "D"},
        {400, "CD"}, {100, "C"},  {90, "XC"},
        {50, "L"},   {40, "XL"},   {10, "X"},
        {9, "IX"}, {5, "V"}, {1, "I"},
    };
    std::string roman;
    int num {number};
    for (const auto &[value, symbol] : valueSymbols)
    {
        while (num >= value)
        {
            num -= value;
            roman += symbol;
        }
        if (num == 0)
        {
            break;
        }
    }
    return roman;
}