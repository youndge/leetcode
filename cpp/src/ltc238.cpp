/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-29 20:53:51
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-29 21:02:43
 * @FilePath: /workspace/leetcode/cpp/src/ltc238.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%
 */
#include "ltc238.h"

void Ltc238::run()
{
    std::vector<int> nums {1,2,3,1}; //5
    for (int i = 0; i < nums.size(); ++i)
    {
        std::cout << "productExceptSelf is " << productExceptSelf(nums)[i] << std::endl;
    }
}

std::vector<int> Ltc238::productExceptSelf(std::vector<int>& nums)
{
    size_t n {nums.size()};
    int pre = 1, suf =1;
    std::vector<int> ans (n,1);
    for (auto i = 1; i < n; ++i)
    {
        size_t right = n -1 -i;
        ans.at(i) *= pre;
        ans.at(right) *= suf;
        pre *= nums.at(i);
        suf *= nums.at(right);

        std::cout << "ans[" << i << "]=" << ans.at(i) << ";ans[" << right << "]=" << ans.at(right) << "; pre=" << pre << "; suf=" << suf << std::endl;
    }
    return ans;
}