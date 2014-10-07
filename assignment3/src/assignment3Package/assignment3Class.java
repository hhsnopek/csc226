package assignment3Package;

interface IClub {
	public void addMembership(String id, int year);
	public void removeMembership(String id);
	public int searchForMembership(String id);
	public void removeMembershipsOfYear(int year);
	public int getLargestPercentageMembershipIncreaseSince(int year);
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
		int percent = 0;

		return percent;
	}
}
