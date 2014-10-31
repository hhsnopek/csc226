package assignment9;

class Element<T> {
  private T value;
  private Element<T> next, prev;

  public Element(T element) {
    value = element;
    next = prev = null;
  }

  public T getValue() {
    return value;
  }

  public void setValue(T value) {
    this.value = value;
  }

  public Element<T> getNext() {
    return next;
  }

  public void setNext(Element<T> next) {
    this.next = next;
  }

  public Element<T> getPrev() {
    return prev;
  }

  public void setPrev(Element<T> prev) {
    this.prev = prev;
  }
}

public class DLSortedList<T> {
  Element<T> root;
  int len;

  public DLSortedList() {
    root = null;
    len = 0;
  }

  public int len() {
    return len;
  }

  public Element<T> Front() {
    if (len == 0) { return null; }
    return root.getNext();
  }

  public Element<T> Back() {
    if (len == 0) { return null; }
    return root.getPrev();
  }

  // inserts element T after at, returns element
  public void insert(Element<T> ele, Element<T> at) {
    Element<T> next = at.getNext();
    at.setNext(ele);
    ele.setPrev(at);
    ele.setNext(next);
    next.setPrev(ele);
    len++;
  }

  public void insertValue (T element, Element<T> at) {
    insert(new Element<T>(element), at);
  }

  public void add(T element) {
    Element<T> newElement  = new Element<T>(element);

    if (root == null) {
      root = newElement;
      len++;
      return;
    }

    // iterator
    for (Element<T> e = root; e != null; e = e.getNext()) {
      // figure out order (negative, zero, positive)
      int order = ((Comparable<T>) e.getValue()).compareTo(element);
      if (order < 0) {
        // insert before e
        insertValue(element, e.getPrev());
      }
      else if (order >= 0) {
        // insert after
        insertValue(element, e);
      }
      else {
        System.out.printf("!Error: Cannot add value(%s) to list.", element);
      }
    }
  }

  public String toString() {
	  String items = "";
	  for (Element<T> e = root; e != null; e = e.getNext())  {
		  items += e.getValue() + ", ";
	  }
	  return items;
  }
}
