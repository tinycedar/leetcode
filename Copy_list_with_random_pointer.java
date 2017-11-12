/**
 * Definition for singly-linked list with a random pointer.
 * class RandomListNode {
 *     int label;
 *     RandomListNode next, random;
 *     RandomListNode(int x) { this.label = x; }
 * };
 */
public class Copy_list_with_random_pointer {
    public RandomListNode copyRandomList(RandomListNode head) {
        if (head == null) {
            return null;
        }
        Map<RandomListNode, RandomListNode> map = new HashMap<>();
        for (RandomListNode n = head; n != null; n = n.next) {
            map.put(n, new RandomListNode(n.label));
        }
        for (RandomListNode n = head; n != null; n = n.next) {
            map.get(n).next = map.get(n.next);
            map.get(n).random = map.get(n.random);
        }
        return map.get(head);
    }
}