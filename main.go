import java.util.List;
import java.util.ArrayList;

public class SonarQubeExample {

    public static void main(String[] args) {
        List<String> names = new ArrayList<String>();
        
        names.add("John");
        names.add("Jane");
        names.add("Mary");
        
        // Code smell: unnecessary use of generics with raw type
        List rawList = new ArrayList();
        rawList.add("Raw List Item");

        // Possible bug: null pointer dereference
        String name = names.get(0);
        if (name != null) {
            System.out.println("Name length: " + name.length());
        }

        // Security vulnerability: hard-coded password
        String password = "12345"; // Hardcoded password: sensitive information
        System.out.println("Password: " + password);
    }

    // Code Smell: Unused method
    public void unusedMethod() {
        System.out.println("This method is never called.");
    }
}