/**
 * 
 */
package leetcode;

/**
 * @author hzlinqien
 *
 */

class ListNode {
	int val;
	ListNode next;

	ListNode(int x) {
		val = x;
	}

	ListNode(int val, ListNode next) {
		this.val = val;
		this.next = next;
	}

	@Override
	public String toString() {
		StringBuilder sb = new StringBuilder();
		for (ListNode cur = this; cur != null; cur = cur.next) {
			sb.append("->").append(cur.val);
		}
		return sb.toString();
	}
}

public class ReverseLinkedList {

	public ListNode reverseList(ListNode head) {
		if (head == null || head.next == null) {
			return head;
		}
		int count = 0;
		for (ListNode cur = head; cur != null; cur = cur.next) {
			count++;
		}
		ListNode[] nodeArray = new ListNode[count];
		count--;
		for (ListNode cur = head; cur != null; cur = cur.next) {
			nodeArray[count--] = cur;
		}
		for (int i = 0; i < nodeArray.length; i++) {
			if (i == nodeArray.length - 1) {
				if (nodeArray[i].next == nodeArray[i - 1]) {
					nodeArray[i].next = null;
				}
			} else {
				nodeArray[i].next = nodeArray[i + 1];
			}
		}
		return nodeArray[0];
	}

	public static void main(String[] args) {
		ListNode head = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, null))));
		System.out.println(head);
		System.out.println(new ReverseLinkedList().reverseList(head));
	}
}