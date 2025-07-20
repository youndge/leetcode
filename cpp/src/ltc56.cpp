/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-29 21:07:56
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-29 21:22:33
 * @FilePath: /workspace/leetcode/cpp/src/ltc56.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc56.h"

void Ltc56::run()
{
    std::vector<std::vector<int>> intervals {{1,3},{2,6},{8,10},{15,18}}; //{{1,6},{8,10},{15,18}}
    std::cout << "intervals is {";
    for (auto it:merge(intervals))
    {
        std::cout << "[";
        for (auto i:it)
        {
            std::cout<<i<<(i!=it.back()?",":"");
        }
        std::cout<<(it==intervals.back()?"]":"], ");
    }
    std::cout<<"]"<<std::endl;
}


std::vector<std::vector<int>> Ltc56::merge(std::vector<std::vector<int>>& intervals)
{
    if (intervals.size() < 2) return intervals;
    std::vector<std::vector<int>> ans;
    std::sort(intervals.begin(), intervals.end());
    int pre = intervals.front().at(0), suf = intervals.front().at(1);
    ans.push_back(std::vector<int> {pre, suf});
    for (auto it = intervals.begin()+1; it != intervals.end(); ++it)
    {
        if ((*it).at(0) <= suf)
        {
            ans.size() ? ans.pop_back():(void)0;
            ans.push_back(std::vector<int> {std::min(pre, (*it).at(0)), std::max(suf, (*it).at(1))});
        }
        else
        {
            ans.push_back(*it);
        }
        pre = ans.back().at(0);
        suf = ans.back().at(1);
    }
    return ans;
}