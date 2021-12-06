package leetcode.algorithms;

import java.util.List;

/**
 * You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.
 *
 * You may assume the two numbers do not contain any leading zero, except the number 0 itself.
 */
public class AddTwoNumbers {
    public static void main(String[] args) {
        System.out.println(7 / 10);
    }

    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        int num = 0;
        ListNode result = new ListNode(0);
        ListNode cur = result;
        while (l1 != null || l2 != null) {
            if (l1 != null) {
                num += l1.val;
                l1 = l1.next;
            }
            if (l2 != null) {
                num += l2.val;
                l2 = l2.next;
            }
            cur.next = new ListNode(num % 10);
            cur = cur.next;
            num /= 10;
        }
        if (num > 0) {
            cur.next = new ListNode(num);
        }
        return result.next;
    }

    private class ListNode {
        int val;
        ListNode next;
        ListNode(int x) {
            val = x;
        }
    }
}
