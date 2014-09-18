package assignment2Package;

import java.util.*;

public class assignment2Class {

	public static void main(String[] args) {

	}

}

class Trail {
	private List<Integer> markers = new ArrayList<Integer>();
	
	public Trail() {
		markers.addAll(Arrays.asList( 100, 150, 105, 120, 90, 80, 50, 75, 75, 70, 80, 90, 100 ));
	}
	
	public Trail(Integer[] intArray) {
		markers.addAll(Arrays.asList(intArray));
	}

	public boolean isLevelTrailSegment(int start, int end) {
		return ((markers.get(start) - markers.get(end - 1)) <= 10) ? true : false;
	}
	
	public int isDifficult() {
		int difficultyCount = 0;
		for (int i = 0; i < markers.size(); i++) {
			if (Math.abs(markers.get(i) - markers.get(i+1)) >= 30) difficultyCount++;
		}
		// return (difficultyCount >= 3) ? true : false;
		return difficultyCount;
	}
	
	public String toString() {
		return markers.toString();
	}

}

class  MyTrailInventory {
	Trail[] trails;
	
	public MyTrailInventory(Trail[] trails) {
		trails = this.trails;
	}
	
	public int getMostDifficultTrail() {
		int index = 0;
		int lastTrail = trails[trails.length].isDifficult();
		
		for(int i = 0; i < trails.length; i++) {
			if (trails.length <= (i+1))
				if (trails[i].isDifficult() > lastTrailDifficulty) index = i;
		}
		return index;
	}
}