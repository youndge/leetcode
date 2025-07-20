/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-07-01 21:32:53
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-07-01 21:38:56
 * @FilePath: /workspace/leetcode/cpp/src/ltc7.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc7.h"

void Ltc7::run()
{
    int x {-321};
    std::cout<<" IntChange: " << IntChange(x) << std::endl;
    int x1 {-120};
    std::cout<<" IntChange: " << IntChange(x1) << std::endl;
}
/**
 * @brief 
 * std::to_string
 * std::swap
 * std::stoll
 * @param x 
 * @return int 
 */
int Ltc7::IntChange(int x)
{
    std::string chars = std::to_string(x);
    int l = (chars.at(0) == '-') ? 1 : 0;
    int r = chars.size() - 1;

    while (l<r)
    {
        std::swap(chars.at(1), chars.at(r));
        ++l;--r;
    }
    long long reversed = std::stoll(chars);
    if (reversed > INT_MAX || reversed < INT_MIN)
    {
        return 0;
    }
    return static_cast<int>(reversed);
}