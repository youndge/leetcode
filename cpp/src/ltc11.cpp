/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-21 18:18:03
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-21 18:25:43
 * @FilePath: /workspace/leetcode/cpp/src/ltc11.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc11.h"

void Ltc11::run()
{
    std::vector<int> v1 {1,8,6,2,5,4,8,3,7};
    std::cout << "v1" << " is " << getMaxWater(v1) << std::endl;
    std::vector<int> v2 {1,1};
    std::cout << "v2" << " is " << getMaxWater(v2) << std::endl;
}

int Ltc11::getMaxWater(const std::vector<int>& height)
{
    int l {0}, r = height.size() - 1, ans {0};
    while (l < r)
    {
        int area = std::min(height.at(l), height.at(r)) * (r - l);
        ans = std::max(ans, area);
        if (height.at(l) <= height.at(r))
        {
            ++l;
        }
        else
        {
            --r;
        }
    }
    return ans;
}