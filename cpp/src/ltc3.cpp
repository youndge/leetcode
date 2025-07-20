#include "ltc3.h"

void Ltc3::run()
{
    //std::string s {"abcabcbb"};
    //std::string s {"pwwkew"};
    std::string s {"bbbbbb"};
    std::cout << "lengthOfSubString: " << lengthOfLongestSubString(s) << std::endl;
}

int Ltc3::lengthOfLongestSubString(std::string s)
{
    std::unordered_map<char, int> dic;
    int i = -1, res = 0, len = s.size();
    for (int j = 0; j < len; ++j)
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