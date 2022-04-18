let max = 100;

let fizzBuzz = fn(x) {
    if(mod(x, 15) == 0) {
        puts(x, " FizzBuzz");
    } else {
        if(mod(x, 3) == 0) {
            puts(x, " Fizz");
        } else {
            if(mod(x, 5) == 0) {
                puts(x, " Buzz");
            }
            else {
                puts(x);
            }
        }
    }

    if(x < max) {
        fizzBuzz(x + 1)
    }
}

fizzBuzz(0)