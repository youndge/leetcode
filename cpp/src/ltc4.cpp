/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-14 17:47:34
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-14 17:58:08
 * @FilePath: /workspace/leetcode/cpp/src/ltc4.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc4.h"

void ltc4::run()
{
    std::string s {"abcabcbb"};
    std::cout << "lengthOfLongestSubString: " << lengthOfLongestSubString(s) << std::endl;
}

int ltc4::lengthOfLongestSubString(std::string s)
{
    std::unordered_map<char, int> dic;
    int i {-1}, res {0}, len = s.size();
    for (int j {0}; j <len; ++j)
    {
        if (dic.find(s.at(j)) != dic.end())
        {
            i = std::max(i, dic.find(s.at(j))->second);
        }
        dic[s.at(j)] = j;
        res = std::max(res, j - i);
    }
    return res;
}