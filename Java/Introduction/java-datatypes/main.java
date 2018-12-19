import java.util.Scanner;

class Main {
    public static String dataType(String num) {
        try {
            String answer = num + " can be fitted in:\n";
            long l = Long.parseLong(num);

            if (l >= -128 && l <= 127) {
                answer += "* byte\n";
            }
            if (l >= -32768 && l <= 32767) {
                answer += "* short\n";
            }
            if (l >= -2147483648 && l <= 2147483647) {
                answer += "* int\n";
            }
            if (l >= (-1 * Math.pow(2, 63)) && l <= (Math.pow(2, 63) - 1)) {
                answer += "* long\n";
            }

            return answer;
        } catch(Exception e) {
            return num + " can't be fitted anywhere.\n";
        }
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String[] nums = new String[sc.nextInt()];

        sc.nextLine();
        for (int i = 0; i < nums.length; i++) {
            nums[i] = sc.nextLine();
        }

        for (String s : nums) {
            System.out.print(dataType(s));
        }
    }
}
