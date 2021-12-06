package leetcode.algorithms;

import java.util.HashSet;

/**
 * For each email address, convert it to the canonical address that actually receives the mail. This involves a few steps:
 *
 * Separate the email address into a local part and the rest of the address.
 *
 * If the local part has a '+' character, remove it and everything beyond it from the local part.
 *
 * Remove all the zeros from the local part.
 *
 * The canonical address is local + rest.
 */
public class UniqueEmails {
    public static void main(String[] args) {
        String[] emails = {"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"};
        new UniqueEmails().numUniqueEmails(emails);
    }
    public int numUniqueEmails(String[] emails) {
        HashSet set = new HashSet();
        for (String email : emails) {
            String[] emailSplit = email.split("@");
            set.add(emailSplit[0].split("\\+")[0].replaceAll("\\.", "") + emailSplit[1]);
        }
        return set.size();
    }
}
