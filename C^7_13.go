package main
import "fmt"

func next(prev bool, a int, b int) (int, int) {
    add := 0
    if (b > 0) {
        second, second_count := next(false, a, b-1)
        if prev {
            add = 1
        }
        if (a == 0) {
            return second + add * second_count, second_count
        }
        first, first_count := next(true, a-1, b)
        return first + second + (1 - add) * first_count + add * second_count,
            first_count + second_count
    }
    if (a > 0) {
        first, first_count := next(true, a-1, b)
        if !prev {
            add = 1
        }
        return first + add * first_count, first_count
    }
    return 0, 1
}

func main() {
    A := 7
    B := 8
    //sum, co := next(true, A-1, B) + next(false, A, B-1);
    afirst, afirst_count := next(true, A-1, B)
    asecond, asecond_count := next(false, A, B-1)
    //fmt.Println(sum, co)
    fmt.Println(afirst + asecond, afirst_count + asecond_count)
    fmt.Println((afirst + asecond) / (afirst_count + asecond_count))
}
