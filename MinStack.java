import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * 
 */

/**
 * @author qienl
 *
 */
public class MinStack {

	private int size;

	private Node head = null;

	private Integer min = null;

	private class Node {
		int value;
		Node next;

		public Node(int value, Node next) {
			this.value = value;
			this.next = next;
		}
	}

	public MinStack() {
	}

	public void push(int x) {
		this.size++;
		this.head = new Node(x, this.head);
		if (this.min == null || x < this.min) {
			this.min = x;
		}
	}

	public void pop() {
		if (head != null) {
			size--;
			int top = head.value;
			head = head.next;
			if (head == null) {
				this.min = null;
			} else if (min != null && top == min) {
				List<Integer> tempList = new ArrayList<>(this.size);
				for (Node temp = head; temp != null; temp = temp.next) {
					if (min == temp.value) {
						return;
					}
					tempList.add(temp.value);
				}
				Collections.sort(tempList);
				this.min = tempList.get(0);
			}
		}
	}

	public int top() {
		if (head != null) {
			return head.value;
		}
		throw new RuntimeException("Empty stack");
	}

	public int getMin() {
		if (head != null && min != null) {
			return this.min;
		}
		throw new RuntimeException("Empty stack");
	}

	public static void main(String[] args) {
		MinStack obj = new MinStack();
		obj.push(-2);
		obj.push(0);
		obj.push(-3);
		obj.getMin();
		obj.pop();
		obj.top();
		System.out.println(obj.getMin());
	}
}
