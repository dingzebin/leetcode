package leetcode.algorithms;

/**
 * @author dzeb
 * @version 1.0
 * @Description 21. Merge Two Sorted Lists
 * @createTime 2020/10/26 23:35
 */
public class MergeTwoSortedLists {

    public static void main(String[] args) {
        printListNode(new MergeTwoSortedLists().mergeTwoLists(createList(new int[] {1, 2, 4}), createList(new int[] {1, 3, 4})));
    }

    public static ListNode createList(int[] args) {
        ListNode sentinel = new ListNode();
        ListNode cur = sentinel;
        for (int i : args) {
            cur.next = new ListNode(i);
            cur = cur.next;
        }
        return sentinel.next;
    }

    public static void printListNode(ListNode l1) {
        while (l1 != null) {
            System.out.println(l1.val);
            l1 = l1.next;
        }
    }

    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode sentinel = new ListNode();
        ListNode cur = sentinel;
        while (l1 != null || l2 != null) {
            if (l1 != null && l2 != null) {
                if (l1.val < l2.val) {
                    cur.next = l1;
                    l1 = l1.next;
                } else {
                    cur.next = l2;
                    l2 = l2.next;
                }
            } else if (l1 != null) {
                cur.next = l1;
                return sentinel.next;
            } else if (l2 != null) {
                cur.next = l2;
                return sentinel.next;
            }
            cur = cur.next;
        }
        return sentinel.next;
    }

    public static class ListNode {
      int val;
      ListNode next;
      ListNode() {}
      ListNode(int val) { this.val = val; }
      ListNode(int val, ListNode next) { this.val = val; this.next = next; }
    }
}
