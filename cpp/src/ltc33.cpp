/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-28 20:48:08
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-28 21:01:44
 * @FilePath: /workspace/leetcode/cpp/src/ltc33.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#pragma once
#include "ltc33.h"

void Ltc33::run()
{
    // std::vector<int> intervals {3,4,5,1,2}; //1
    std::vector<int> intervals {4,5,6,7,0,1,2}; //0
    std::cout << "findMin" << " is " << search(intervals, 0) << std::endl;
}

int Ltc33::search(std::vector<int>& nums, int target)
{
    auto mid = nums.size() / 2;
    for (;(mid < nums.size() && nums.at(mid) > nums.at(mid + 1))
        || (nums.at(mid) > nums.at(mid - 1) && mid > 0);)
    {
        int idx = (mid < nums.size() && nums.at(mid) > nums.at(mid + 1)) ? (nums.size() - 1): 0;
        mid = (mid + idx) / 2;
        if (mid == 0 || mid == nums.size() - 1)
        {
            return -1;
        }
    }
    return mid;
}