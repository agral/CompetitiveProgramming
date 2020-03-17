import java.io.*;
import java.util.*;

public class SimpleArraySum
{
  static int simpleArraySum(int[] array) {
    int sum = 0;
    for (int i = 0; i < array.length; ++i) {
      sum += array[i];
    }

    return sum;
  }

  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) throws IOException {
    int arrSize = Integer.parseInt(scanner.nextLine().trim());
    int[] array = new int[arrSize];

    String[] strArray = scanner.nextLine().split(" ");
    for (int i = 0; i < arrSize; ++i)
    {
      array[i] = Integer.parseInt(strArray[i].trim());
    }

    int sum = simpleArraySum(array);
    System.out.println(String.valueOf(sum));
  }
}
