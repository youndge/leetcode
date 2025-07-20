/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-03-04 22:57:34
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-14 11:35:36
 * @FilePath: /cpp/src/ltc01.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include <iostream>
#include "ltc01.h"
#include<map>

std::vector<int> Ltc01::sum(std::vector<int>& nums, int t)
{
    std::vector<int> res(2, -1);
    if (nums.size() < 2)
    {
        return res;
    }

    for (auto i = 0; i<nums.size() - 1; ++i)
    {
        for (auto j = 1; j < nums.size(); ++j)
        {
            if (nums.at(i) + nums.at(j) == t)
            {
                res.at(0) = i;
                res.at(1) = j;
            }
        }
    }
return res;
}

void Ltc01::run()
{
    std::vector<int> nums = {2,7,5,11,15};
    int target = 9;
    Ltc01 s;
    std::cout << s.sum(nums, target).at(0) << " " << s.sum(nums, target).at(1) << std::endl;
}