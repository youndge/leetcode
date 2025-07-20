/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-22 21:24:09
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-22 21:35:39
 * @FilePath: /workspace/leetcode/cpp/src/ltc121.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc121.h"

void Ltc121::run()
{
    std::vector<int> prices {7,1,5,3,6,4}; //5
    std::cout << "max profit" << " is " << maxProfit(prices) << std::endl;
    std::vector<int> prices2 {7,6,4,3,1}; //0
    std::cout << "max profit" << " is " << maxProfit(prices2) << std::endl;    
}

int Ltc121::maxProfit(std::vector<int>& prices)
{
    int minPrice = 1e9;
    int maxProfit {0};
    for (auto price : prices)
    {
        maxProfit = std::max(maxProfit, price - minPrice);
        minPrice = std::min(price, minPrice);
    }
    return maxProfit;
}