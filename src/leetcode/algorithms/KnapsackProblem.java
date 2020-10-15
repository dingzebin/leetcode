package leetcode.algorithms;

import java.util.ArrayList;
import java.util.List;

/**
 * @description:
 * @author: dzeb
 * @date: 2019/7/24 15:29
 **/
public class KnapsackProblem {
    public static void main(String[] args) {
        List<Article> articles = new ArrayList<>();
        articles.add(new Article(12, 4));
        articles.add(new Article(2, 2));
        articles.add(new Article(1, 2));
        articles.add(new Article(1, 1));
        articles.add(new Article(4, 10));

        System.out.println(fun(articles, 0, 15));
    }

    public static int fun(List<Article> articles, int idx, int weight) {
        int maxValue = 0;
        for (int i = idx; i < articles.size(); i++) {
            if (articles.get(i).weight < weight) {
                maxValue = Math.max(maxValue, articles.get(i).price + fun(articles, i + 1, weight - articles.get(i).weight));
            }
        }
        return maxValue;
    }

    private static class Article {
        private int weight;
        private int price;
        public Article(int weight, int price) {
            this.weight = weight;
            this.price = price;
        }
    }
}
