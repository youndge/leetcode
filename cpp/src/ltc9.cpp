/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-29 20:43:01
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-29 20:47:49
 * @FilePath: /workspace/leetcode/cpp/src/ltc9.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc9.h"

void Ltc09::run()
{
    int x1 {121};
    std::cout << x1 << " is " << IsHuiwen(x1) << std::endl;
    int x2 {-121};
    std::cout << x2 << " is " << IsHuiwen(x2) << std::endl;
    int x3 {10};
    std::cout << x3 << " is " << IsHuiwen(x3) << std::endl;
}

bool Ltc09::IsHuiwen(int x)
{
    if ((x < 0) || (x % 10 == 0 && x!=0))
        return false;
    int revertedNumber {0};
    while (x > revertedNumber)
    {
        revertedNumber = revertedNumber * 10 + x % 10;
        x /= 10;
    }
    return x == revertedNumber || x == revertedNumber /10;
}