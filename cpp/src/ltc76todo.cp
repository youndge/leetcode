/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-30 21:50:34
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-30 22:32:03
 * @FilePath: /workspace/leetcode/cpp/src/ltc76.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc76.h"

void Ltc76::run()
{
    std::string s {"AD0BEC0DEBANC"}, t {"ABC"}; //{"BANC"}
    std::cout<<"minWindow is "<<minWindow(s,t) << std::endl;
}

std::string Ltc76::minWindow(std::string s, std::string t)
{
    std::string ans {""}, temp {s};
    std::vector<int> idx;
    int eraseCounter {0};
    for (int i = 0; i < t.size(); ++i)
    {
        if (auto pos = s.find(t.at(i));std::string::npos != pos)
        {
            if (auto posInVec = std::find(idx.begin(), idx.end(), ))
            idx.push_back(pos);
            std::cout<<pos<<std::endl;
        }
        else
        {
            return ans;
        }
    }
    std::sort(idx.begin(),idx.end());
    auto print = [](int v){std::cout<<v<<" ";};
    std::for_each(idx.begin(), idx.end(), print);
    ans=s.substr(s.at(idx.front()), (s.at(idx.back()) - s.at(idx.front())));
    return ans;
 }