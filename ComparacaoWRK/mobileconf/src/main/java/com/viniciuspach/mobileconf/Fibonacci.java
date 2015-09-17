/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package com.viniciuspach.mobileconf;

import static spark.Spark.*;

/**
 *
 * @author vinicius
 */
public class Fibonacci {

    public static void main(String[] args) {
        get("/:number", (req, res) -> {
            int number = Integer.parseInt(req.params(":number"));
            return "Sparkjava: fib(" + number + ") = " + fib(number);
        });
    }

    private static int fib(int n) {
        if (n == 0) {
            return 0;
        } else if (n == 1) {
            return 1;
        } else {
            return fib(n - 1) + fib(n - 2);
        }
    }
}
