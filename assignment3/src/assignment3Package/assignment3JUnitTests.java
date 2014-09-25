package assignment3Package;

import static org.junit.Assert.*;
import org.junit.Test;

public class assignment3JUnitTests {

	@Test
	public void testClassMembership_customInit_getID() {
		Membership newMembership = new Membership("JackBlack", 2001);
		assertEquals("JackBlack", newMembership.getID());
	}
	
	@Test
	public void testClassMembership_defaultInit_getEnrollmentYear() {
		Membership newMembership = new Membership();
		assertEquals(2000, newMembership.getEnrollmentYear());
	}
	
	@Test
	public void testClassMembership_customInit_getEnrollmentYear() {
		Membership newMembership = new Membership("JerryTom", 2014);
		assertEquals(2014, newMembership.getEnrollmentYear());
	}
	
	@Test(expected=IndexOutOfBoundsException.class)
	public void testClassClub_defaultInit_addMembership_shouldFail() {
		Club testClub1 = new Club();
		testClub1.addMembership("JackBlack", 2001);
		assertEquals(0, testClub1.searchForMembership("JackBlack"));
	}
	
	@Test
	public void testClassClub_customInit_addMembership() {
		Club testClub1 = new Club(5);
		testClub1.addMembership("JackBlack", 2001);
		testClub1.addMembership("JerryTom", 2014);
		assertEquals(1, testClub1.searchForMembership("JerryTom"));
	}
	
	@Test
	public void testClassClub_defaultInit_searchForMembership() {
		Club testClub1 = new Club();
		assertEquals(-1, testClub1.searchForMembership("John"));
	}
	
	@Test
	public void testClassClub_customInit_searchForMembership() {
		Club testClub1 = new Club(2);
		testClub1.addMembership("JackBlack", 2001);
		testClub1.addMembership("JerryTom", 2014);
		assertEquals(1, testClub1.searchForMembership("JerryTom"));
	}
	
	@Test
	public void testClassClub_removeMembership() {
		Club testClub1 = new Club(2);
		testClub1.addMembership("JackBlack", 2001);
		testClub1.addMembership("JerryTom", 2014);
		testClub1.removeMembership("JackBlack");
	}
	
	@Test
	public void testClassClub_removeMembershipsOfYear() {
		Club testClub1 = new Club(3);
		testClub1.addMembership("JackBlack", 2001);
		testClub1.addMembership("SamJohn", 2001);
		testClub1.addMembership("JerryTom", 2014);
		testClub1.removeMembershipsOfYear(2001);
	}
	
	@Test
	public void testClassClub_getLargestPercentageMembershipIncreaseSince() {
		Club testClub1 = new Club(7);
		testClub1.addMembership("JackBlack", 2001);
		testClub1.addMembership("SamJohn", 2001);
		testClub1.addMembership("Jacklack", 2002);
		testClub1.addMembership("SamJhn", 2002);
		testClub1.addMembership("JeryTom", 2007);
		testClub1.addMembership("JerryTom", 2014);
		
		testClub1.getLargestPercentageMembershipIncreaseSince(2002);
	}

}
