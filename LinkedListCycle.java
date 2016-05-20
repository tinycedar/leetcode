/**
 * 
 */
package leetcode;

import java.util.HashSet;
import java.util.Set;

/**
 * @author hzlinqien
 *
 */

public class LinkedListCycle {

	public boolean hasCycle(ListNode head) {
		Set<ListNode> nodeMap = new HashSet<>();
		for (ListNode temp = head; temp != null; temp = temp.next) {
			if (nodeMap.contains(temp)) {
				return true;
			}
			nodeMap.add(temp);
		}
		return false;
	}

	class ListNode {
		int val;
		ListNode next;

		ListNode(int x) {
			val = x;
			next = null;
		}
	}
}
