/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-14 17:08:21
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-14 17:18:45
 * @FilePath: /workspace/leetcode/cpp/src/ltc2.cpp
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#include "ltc2.h"

void ListNode::traverse()
{
    ListNode* current = this;
    while (current)
    {
        std::cout << current->val << " ";
        current = current->next;
    }
    std::cout << std::endl;
}

ListNode* ListNode::convertListNode(const std::vector<int>& array)
{
    if (array.size() == 0)
    {
        return nullptr;
    }
    ListNode* current = new ListNode(array.at(0));
    ListNode* head = current;
    for (auto i {1}; i < array.size(); i++)
    {
        ListNode* node = new ListNode(array.at(i));
        current->next = node;
        current = current->next;
    }
    return head;
}

void ltc2::run()
{
    std::vector<int> v1 {2,4,3}, v2 {5,6,4};
    ListNode* l1 = new ListNode(v1);
    ListNode* l2 = new ListNode(v2);
    l1->traverse();
    l2->traverse();

    ListNode* l3 = addTwoNumbers(l1,l2);
    l3->traverse();
}