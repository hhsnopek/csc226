package assignment3Package;

interface IClub {
	public void addMembership(String id, int year);
	public void removeMembership(String id);	
	public int searchForMembership(String id);	
	public void removeMembershipsOfYear(int year);
	public int getLargestPercentageMembershipIncreaseSince(int year);
}

class point {
	public int x;
	public int y;
	public point() { this.x = 2000; this.y = 0; }
	public void setPoints(int x, int y) { this.x = x; this.y = y; }
	public int[] getPoints() { return new int[] { this.x, this.y }; }
	public void incrementY() { y++; }
}

class Membership {
	String ID;
	int enrollmentYear;
	
	public Membership() {
		ID = "JohnDoe";
		enrollmentYear = 2000;
	}
	
	public Membership(String ID, int year) {
		this.ID = ID;
		this.enrollmentYear = year;
	}
	
	public String getID() {
		return ID;
	}
	
	public int getEnrollmentYear() {
		return enrollmentYear;
	}
}

class Club implements IClub {
	Membership[] members;
	int numMemberships;
	
	public Club() {
		members = new Membership[0];
	}
	
	public Club(int memberLimit) {
		members = new Membership[memberLimit];
	}
	
	public void addMembership(String id, int year) {
		if ((id.equalsIgnoreCase("")) || (year < 2000)) return;
		members[this.numMemberships] = new Membership(id, year);
		this.numMemberships++;
	}
	
	public void removeMembership(String id) {
		if (id.equalsIgnoreCase("")) return;
		members[searchForMembership(id)] = null;
	}
	
	public int searchForMembership(String id) {
		int index = -1;
		
		if (id.equalsIgnoreCase("")) return -1;
		for (int j = 0; j < members.length; j++) {
			if ((members[j].getID()).equalsIgnoreCase(id)) {
				return j;
			}
		}
		return index;
	}
	
	public void removeMembershipsOfYear(int year) {
		if (year < 2000) return;
		for (int i = 0; i < members.length; i++) {
			if (members[i].getEnrollmentYear() == year) members[i] = null;
		}
	}
	
	public int getLargestPercentageMembershipIncreaseSince(int year) {
		int[] yearCollection = new int[members.length];
		point[] dataPoints = new point[members.length];
		int percent = 0;
		
		// gather all years
		for (int i = 0; i < members.length; i++) {
			yearCollection[i] = members[i].getEnrollmentYear();
		}
		
		// count amount of years and place in dataPoints (object array)
		for (int i = 0; i < yearCollection.length; i++) {
			// search for year
			for (int c = 0; c < dataPoints.length; c++) {
				if (dataPoints[c].getPoints()[1] == yearCollection[i]) {
					// add to year if found
					dataPoints[c].incrementY();
				}
			}
		}
		
		int[] initialPoints = dataPoints[0].getPoints();
		
		// perform derivative on data
		for (int i = 0; i < dataPoints.length; i++) {
			int[] currentPoints = dataPoints[i].getPoints();
			int midpoint = _midPoint(currentPoints[0], currentPoints[1], initialPoints[0], initialPoints[1]);

			// if percent is > than previous, resign percent
			if (midpoint > percent) {
				percent = midpoint;
				initialPoints = currentPoints;
			}
		}

		return percent;
	}
	
	private int _midPoint(int x2, int y2, int x1, int y1) { int percent = ((x2 - x1) / (y2 - y1)); return percent; }
}
