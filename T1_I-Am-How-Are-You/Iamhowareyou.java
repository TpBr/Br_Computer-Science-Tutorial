import java.io.*;

//reference: https://www.geeksforgeeks.org/console-readline-method-in-java-with-examples/ and https://www.quantamagazine.org/how-godels-incompleteness-theorems-work-20200714/

public class Iamhowareyou {
    public static void main(final String[] args) throws IOException { // need to convert to basic io

        Console standIn = System.console();
                      
        String temp = standIn.readLine("I am Ajavaprogam.  How are you (good/any other typed input)? \n");

        if (temp.contains("good")) {
            System.out.println("Great!");
        } else {
            System.out.println("Sounds like you are definitely not good right now.");
        }

        System.out.println("Say check this out. Let '0' be 6, let the successor function be 7, let '(' be 8, let ')' be 9, let '=' be 5, and let '+' be 11.");

    System.out.println("Then...");

    System.out.println("Wait for it...");

    standIn.readLine("The formula '1 + 2 = 3' has the Godel number 263871010580431552941101982536173915134895149480003604169911130273946997941207917298748114520726903950820819280916133192695469786920156944682314047557861444082586688365002070884826660983439423268888302848351517888024116980474108961030912541130206000000 and is true (if I did my basic math right). What do you think about that! \n");
	
    System.out.println("Well it is now my favorite number and encoding of truth.");

    }
}
