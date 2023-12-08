(define add (a b) (+ a b))

(define lessThanTen (n) (< n 10))

(set x 5)

(and? (< x 10) (> x 3))

(add x 3)

(define factorial (n) 
    (cond 
        (= n 1) 
            1 
        true 
            (* n (factorial (- n 1)))))

(factorial 35)

(lessThanTen x)

