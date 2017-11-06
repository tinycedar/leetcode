package com.netease.kaola.act.compose;

public class Intersection_of_two_linked_lists {

    public static class ListNode {
        int val;
        ListNode next;

        ListNode(int x) {
            val = x;
            next = null;
        }

        ListNode(int x, ListNode next) {
            this.val = x;
            this.next = next;
        }
    }

    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        if (headA == null || headB == null) {
            return null;
        }
        int sizeA = 1;
        ListNode tailA = headA;
        while (tailA != null && tailA.next != null) {
            sizeA++;
            tailA = tailA.next;
        }

        int sizeB = 1;
        ListNode tailB = headB;
        while (tailB != null && tailB.next != null) {
            sizeB++;
            tailB = tailB.next;
        }
        if (tailA.val != tailB.val) {
            return null;
        }

        int diff = sizeA - sizeB;
        ListNode startA = diff > 0 ? headA : headB;
        for (int i = Math.abs(diff); i > 0; i--) {
            startA = startA.next;
        }

        ListNode startB = diff > 0 ? headB : headA;
        while (startA != null && startB != null) {
            if (startA.val == startB.val) {
                return startA;
            }
            startA = startA.next;
            startB = startB.next;
        }
        return null;
    }

    public static void main(String[] args) {
        Intersection_of_two_linked_lists i = new Intersection_of_two_linked_lists();

        ListNode a = new ListNode(3, null);
        ListNode b = new ListNode(2, new ListNode(3, null));
        ListNode result = i.getIntersectionNode(a, b);
        System.err.println(result);
    }
}


