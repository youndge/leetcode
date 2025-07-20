/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-22 21:07:16
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-30 21:46:59
 * @FilePath: /workspace/leetcode/cpp/src/ltc49.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc49.h"

using STRVEC = std::vector<std::vector<std::string>>;

void Ltc49::run()
{
    std::vector<std::string> strs {"eat", "tea", "tan", "ate", "nat", "bat"};
    //{[bat], ["nat", "tan"], ["ate", "eat", "tea"]}
    std::vector<std::vector<std::string>> print;
    print = groupAnagram(strs);
    std::cout << "{ }";
    for (auto subVec:print)
    {
        std::cout << "[";
        for (auto elem: subVec)
        {
            std::cout << elem << ((elem != subVec.back())? ", ": "");
        }
        std::cout << ((subVec != print.back())? "], ": "]");
    }
    std::cout << "}" << std::endl;
}

std::vector<std::vector<std::string>> Ltc49::groupAnagram(std::vector<std::string> strs)
{
    std::vector<std::vector<std::string>> ans;
    std::vector<std::string> temps {strs};
    for (auto temp:temps)
    {
        std::sort(temp.begin(), temp.end());
    }
}