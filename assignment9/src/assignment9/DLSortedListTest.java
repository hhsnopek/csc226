package assignment9;

public class DLSortedListTest {
  public static void main(String[] args) {
    System.out.println("Hello World");
    // list tests
    DLSortedList<Integer> list = new DLSortedList<Integer>();
    list.add(1);
    list.add(5);
    //list.add(2);
    System.out.println(list);
    // Element test
    Element<Integer> item = new Element(null);
    System.out.println(item.getValue());
    System.out.println(item.getNext());
    System.out.println(item.getPrev());
  }
}
