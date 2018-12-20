package leetcode.algorithms.uniqueMorseRepresentations;

import java.util.HashSet;
import java.util.Set;

public class UniqueMorseRepresentations {
    public static void main(String[] args) {

    }
    public int uniqueMorseRepresentations(String[] words) {
        String[] chars = {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};
        StringBuilder sb = null;
        Set<String> set = new HashSet<>();
        for (String word : words) {
            sb = new StringBuilder();
            for (int i = 0; i < word.length(); i++) {
                sb.append(chars[word.charAt(i) - 'a']);
            }
            set.add(sb.toString());
        }
        return set.size();
    }
}
