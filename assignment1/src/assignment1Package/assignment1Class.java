package assignment1Package;

import java.util.*;

public class assignment1Class {

	/**
	 * @param args
	 */
	public static void main(String[] args) {
		System.out.println("Assignment 1 - Henry Snopek");
		
		
		System.out.println("\nProblem 1");
		System.out.println("TAN: \t\t" + scrambleWord("TAN"));
		System.out.println("ABRACADABRA: \t" + scrambleWord("ABRACADABRA"));
		System.out.println("WHOA: \t\t" + scrambleWord("WHOA"));
		System.out.println("AARDVRAK: \t" + scrambleWord("AARDVRAK"));
		System.out.println("EGGS: \t\t" + scrambleWord("EGGS"));
		System.out.println("A: \t\t" + scrambleWord("A"));
		System.out.println(": " + scrambleWord(""));
		
		System.out.println("\nProblem 2");
		List<String> wordList = Arrays.asList("TAN", "ABRACADABRA", "WHOA", "APPLE", "EGGS");
		System.out.println("wordList: \t" + wordList);
		scrambleOrRemove(wordList);
		
		System.out.println("\nProblem 3");
		System.out.println("wordList: \t" + wordList);
		System.out.println("total chars: \t" + getNumCharacters(wordList));
		
		System.out.println("\nProblem 4");
		String text = "BARCAD,ABARA";
		System.out.println("text: \t\t" + text);
		scrambleAndDisplay(text);
	}
	
	public static String scrambleWord(String word) {
		String scrambledString = "";
		int wordLen = word.length();
		word = word.toUpperCase();
		
		for (int i = 0; i < wordLen; i++) {
			if (word.charAt(i) == 'A') {
				// check to see if (i + 1) exists
				if ((i + 1) != wordLen) {
					// if next letter is not 'A', add it before the current letter
					if (word.charAt(i + 1) != 'A') {
						scrambledString += word.charAt(i + 1);
						scrambledString += word.charAt(i);
						i++;
					}
					else {
						scrambledString += word.charAt(i);
					}
				}
				else {
					scrambledString += word.charAt(i);
				}
			}
			else {
				// check to see if (i - 1) is not wordLen
				if ((i - 1) != wordLen)
					scrambledString += word.charAt(i);
			}
		}
		return scrambledString;
	}
	
	public static void scrambleOrRemove(List<String> wordList) {
		List<String> newWordList = new ArrayList<String>();
		
		for (String word : wordList) {
			String scrambledWord = scrambleWord(word);
			
			if (!word.equals(scrambledWord))
				newWordList.add(scrambledWord);
		}
		
		System.out.println("newWordList: \t" + newWordList);
	}
	
	public static int getNumCharacters(List<String> wordList) {
		int totalChar = 0;
		
		for (String word : wordList) {
			totalChar += word.length();
		}
		
		return totalChar;
	}
	
	public static void scrambleAndDisplay(String text) {
		String newText = scrambleWord(text);
		
		if (!text.equals(newText))
			System.out.println("newText: \t" + newText.replaceAll("[.,?!]+", ""));
	}

}
