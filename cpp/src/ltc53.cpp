/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-28 21:06:56
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-28 21:12:26
 * @FilePath: /workspace/leetcode/cpp/src/ltc53.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc53.h"

void Ltc53::run()
{
    std::vector<int> nums {-2,1,-3,4,-1,2,1,-5,4}; //[4,-1,2,1]  6
    // std::vector<int> nums {6,4,-1,7,8}; //23
    std::cout << "maxSubArray is " << maxSubArray(nums) << std::endl; 
}

int Ltc53::maxSubArray(std::vector<int>& nums)
{
    int pre = 0, maxValue = nums.at(0);
    for (auto it : nums)
    {
        pre = std::max(it, pre+it);
        maxValue = std::max(pre, maxValue);
    }
    return maxValue;
}