(define assertEquals (actual expected) 
    (cond 
        (= expected actual) 
            "OK" 
        (true) 
            "FAIL"))

(+ "Expect OK: " (assertEquals 1 1))
(+ "Expect FAIL: " (assertEquals true nil))

(define add (a b) (+ a b))

(define lessThanTen (n) (< n 10))

(set x 5)

(assertEquals (and? (< x 10) (> x 3)) true)

(assertEquals (add x 3) 8)

(define factorial (n) 
    (cond 
        (= n 1) 
            1 
        true 
            (* n (factorial (- n 1)))))

(assertEquals (factorial 5) 120)

(assertEquals (lessThanTen x) true)

""
"Test arithmetic operations"
(assertEquals (+ 3 4) 7)
(assertEquals (- 5 2) 3)
(assertEquals (* 6 2) 12)
(assertEquals (/ 8 2) 4)

""
"Test comparison operations"
(assertEquals (= 3 3) true)
(assertEquals (< 3 4) true)
(assertEquals (> 5 2) true)

""
"Test car and cdr on lists"
(define firstElement (lst) (car lst))
(define restOfList (lst) (cdr lst))

""
"Assuming (list 1 2 3) creates a list [1, 2, 3]"
(set myList (1 2 3))

(assertEquals (firstElement myList) 1)
(assertEquals (firstElement(restOfList myList)) 2)

""
"Test type checking functions"
(assertEquals (number? 5) true)
(assertEquals (symbol? x) true)
(assertEquals (list? myList) true)
(assertEquals (nil? nil) true)

""
"Logical operations"
(assertEquals (and? (< 3 4) (> 5 2)) true)
(assertEquals (or? (< 3 2) (> 5 6)) nil)

""
"Function with multiple arguments"
(define sumThree (a b c) (+ a (+ b c)))
(assertEquals (sumThree 1 2 3) 6)

""
"More complex recursive function - Fibonacci"
(define fib (n) 
  (cond 
    (= n 0) 0
    (= n 1) 1
    true (+ (fib (- n 1)) (fib (- n 2)))
  )
)

(assertEquals (fib 5) 5)

""
"Testing global variable assignment and usage"
(set globalVar 10)
(assertEquals globalVar 10)