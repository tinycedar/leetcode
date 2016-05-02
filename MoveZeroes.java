public class MoveZeroes {

    public static void main(String[] args) {
        int[] nums = new int[]{0, 1, 0, 3, 12};
        new MoveZeroes().moveZeroes(nums);
        for (int x : nums) {
            System.out.print(x + " ");
        }
    }

    public void moveZeroes(int[] nums) {
        if (nums != null && nums.length > 0) {
            int length = nums.length;
            int[] zeroArray = new int[length + 1];
            int[] nonZeroArray = new int[length + 1];

            int zeroIndex = 0;
            int nonZeroIndex = 0;
            for (int i = 0, j = length -1; i < j; i++, j--) {
                if (nums[i] == 0) {
                    zeroArray[zeroIndex++] = i + 1;
                } else {
                    nonZeroArray[nonZeroIndex++] = i + 1;
                }
            }
            for (int i = 0; i < zeroArray.length; i++) {
                int zIndex = zeroArray[i];
                int nzIndex = nonZeroArray[i];
                if (zIndex > 0 && zIndex < nzIndex) {
                    nums[zIndex - 1] = nums[nzIndex - 1];
                    nums[nzIndex - 1] = 0;
                }
            }
        }
    }
}
