/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-29 19:49:00
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-29 20:01:48
 * @FilePath: /workspace/leetcode/cpp/src/ltc217.cc
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc217.h"

void Ltc217::run()
{
    std::vector<int> nums {1,2,3,1}; //1
    std::cout << "containDuplicate is " << containDuplicate(nums) << std::endl;
    std::vector<int> nums2 {1,2,3,4}; //0
    std::cout << "containDuplicate is " << containDuplicate(nums2) << std::endl;
}

bool Ltc217::containDuplicate(std::vector<int>& nums)
{
    std::map<int, int> m;
    for (auto it : nums)
    {
        if (m.count(it) > 0)
        {
            return true;
        }
        m[it] = 1;
    }
    return false;
}