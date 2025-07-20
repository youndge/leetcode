/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2025-06-14 16:09:25
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-06-14 17:32:08
 * @FilePath: /workspace/leetcode/cpp/inc/ltc2.h
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
#pragma once
#include <iostream>
#include <vector>

class ListNode
{
public:
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(NULL) {}
    ListNode(std::vector<int>& arrary) : val(arrary.at(0)), next(NULL) {}
    void traverse();
    ListNode* convertListNode(const std::vector<int>& array);
};

class ltc2
{
public:
    void run();
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int len1 {1}, len2 {1};
        ListNode* p {l1};
        ListNode* q {l2};
        while (p->next) {len1++; p = p->next;}
        while (q->next) {len2++; q = q->next;}
        if (len1 > len2)
        {
            for (int i = 1; i <= len1 - len2; i++)
            {
                q->next = new ListNode(0);
                q = q-> next;
            }
        }
        else
        {
            for (int i = 1; i <= len2 - len1; i++)
            {
                p->next = new ListNode(0);
                p = p->next;
            }
        }
        p = l1; q = l2;
        bool count {false};
        ListNode* l3 = new ListNode(-1);
        ListNode* w = l3;
        int i {0};
        while(p&&q)
        {
            i = count + p->val + q->val;
            w->next = new ListNode(i%10);
            count = i >= 10 ? true : false;
            w = w->next;
            p = p->next;
            q = q->next;
        }
        if (count)
        {
            w->next = new ListNode(1);
            w = w->next;
        }
        return l3->next;
    }
};
