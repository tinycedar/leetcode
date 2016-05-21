import java.text.MessageFormat;

/**
 * 
 */

/**
 * @author qienl
 *
 */
public class FirstBadVersion {

	public int firstBadVersion(int n) {
		return doFind(1, n);
	}

	private int doFind(int begin, int end) {
		int middle = Long.valueOf((Long.valueOf(end) + Long.valueOf(begin)) / 2).intValue();
//		System.out.println(MessageFormat.format("[{0}, {1}]\t{2}", begin, end, middle));
		if (end - begin <= 1) {
			return isBadVersion(begin) ? begin : end;
		}
		if (isBadVersion(middle)) {
			return doFind(begin, middle);
		} else {
			return doFind(middle + 1, end);
		}
	}

	private boolean isBadVersion(int n) {
		return n >= 1702766719;
	}

	public static void main(String[] args) {
		System.out.println(new FirstBadVersion().firstBadVersion(2126753390));
	}
}
