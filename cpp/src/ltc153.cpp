/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-21 17:51:50
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-21 18:04:38
 * @FilePath: /workspace/leetcode/cpp/src/ltc153.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc153.h"

void Ltc153::run()
{
    // std::vector<int> intervals {3,4,5,1,2};
    std::vector<int> intervals {4,5,6,7,0,1,2};
    std::cout << "findMin" << " is " << findMin(intervals) << std::endl;
}

int Ltc153::findMin(std::vector<int>& nums)
{
    auto mid = nums.size() / 2;
    for (; (mid < nums.size() && nums.at(mid) > nums.at(mid + 1)) || (nums.at(mid) > nums.at(mid - 1) && mid > 0);)
    {
        int idx = (mid < nums.size() && nums.at(mid) > nums.at(mid + 1)) ? (nums.size()-1):0;
        mid = (mid + idx) / 2;
        if (mid == 0 || mid == nums.size() - 1)
        {
            break;
        }
    }
    return nums.at(mid);
}