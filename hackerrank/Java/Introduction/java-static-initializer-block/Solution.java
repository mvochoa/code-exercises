import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Solution {

    public static boolean flag = funcFlag();
    public static int B;
    public static int H;

    private static boolean funcFlag() {
        Scanner sc = new Scanner(System.in);
        B = Integer.parseInt(sc.nextLine());
        H = Integer.parseInt(sc.nextLine());

        if (B <= 0 || H <= 0) {
            System.out.println("java.lang.Exception: Breadth and height must be positive");
            return false;
        }
        return true;
    }

    public static void main(String[] args) {
        if (flag) {
            int area = B * H;
            System.out.print(area);
        }

    }

}
