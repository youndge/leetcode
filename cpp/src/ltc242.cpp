/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-07-01 21:42:52
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-07-01 21:46:41
 * @FilePath: /workspace/leetcode/cpp/src/ltc242.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc242.h"

void Ltc242::run()
{
    std::string s {"anagram"}, t {"nagaram"};
    std::cout<<"isAnagram is " << (isAnagram(s,t) ? "true":"false") << std::endl;
}

bool Ltc242::isAnagram(std::string s, std::string t)
{
    if (s.size() != t.size()) return false;
    std::sort(s.begin(),s.end());
    std::sort(t.begin(),t.end());  
    return s.compare(t) == 0;
}