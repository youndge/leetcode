/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-30 22:12:56
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-30 22:25:39
 * @FilePath: /workspace/leetcode/cpp/src/ltc15.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc15.h"

void Ltc15::run()
{
    // std::vector<int> intervals {0,1,1}; //[]
    std::vector<int> intervals {-1,0,1,2,-1,-4}; //{{-1,-1,2},{-1,0,1}}
    // std::vector<int> intervals {11,13,15,17}; //11

    for (auto it:threeSum(intervals))
    {
        std::cout<<"[]";
        for (auto i:it)
        {
            std::cout<<i<<(i!=it.back()?",":"");
        }
        std::cout<<(it==threeSum(intervals).back()?"]":"], ");
    }
    std::cout<<"}"<<std::endl;
}

std::vector<std::vector<int>> Ltc15::threeSum(std::vector<int>& nums)
{
    std::sort(nums.begin(),nums.end());
    std::vector<std::vector<int>> ans;
    for (int i = 0, j = nums.size()-1;i<j;++i,++j)
    {
        auto it = std::find(nums.begin()+i,nums.begin()+j,-(nums.at(i)+nums.at(j)));
        if(nums.end()!=it)
        {
            std::vector<int> temp {nums.at(i), nums.at(j),*it};
            std::sort(temp.begin(),temp.end());
            if (std::find(ans.begin(), ans.end(),temp) == ans.end())
            {
                ans.push_back(temp);
            }
        }
    }
    return ans;
}