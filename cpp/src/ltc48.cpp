/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-29 21:27:13
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-29 21:34:52
 * @FilePath: /workspace/leetcode/cpp/src/ltc48.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc48.h"

void Ltc48::run()
{
    std::vector<std::vector<int>> intervals {{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}; 
    //{{15,13,2,5},{14,3,4,1},{12,6,8,9},{16,7,10,11}}
    std::cout << "matrix is {";
    rotate(intervals);
    for (auto it:intervals)
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

void Ltc48::rotate(std::vector<std::vector<int>>& matrix)
{
    std::vector<int> temp;
    for (auto i=0; i < matrix.size(); ++i)
    {
        for (auto j=0;j<matrix.size();++j)
        {
            temp.push_back(matrix.at(i).at(j));
        }
    }
    for (int i = 0; i<matrix.size();++i)
    {
        int j = matrix.size() - 1;
        for (;j>=0;--j)
        {
            matrix.at(j).at(i) = temp.back();
            temp.pop_back();
        }
    }
}