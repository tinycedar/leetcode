package leetcode;

public class PowerOfTwo {

	public boolean isPowerOfTwo(int n) {
		if (n <= 0) {
			return false;
		}
		while (n % 2 == 0) {
			n /= 2;
		}
		return n == 1;
	}

	public static void main(String[] args) {
		System.out.println(new PowerOfTwo().isPowerOfTwo(6));
	}
}
