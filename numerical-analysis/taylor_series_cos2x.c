#include <stdio.h>
#include <math.h>

int factorial(int n) {
    if (n == 0 || n == 1) {
        return 1;
    }
    return n * factorial(n - 1);
}

int main() {
    double x0 = 0.0; // Maclaurin serisi x0 = 0 civarında yapılır
    double x = 0.2;  // f(0.2) değerini bulmak istiyoruz
    double f_x = 0.0; // Taylor serisinin toplamı burada tutulacak

    // 1. Terim: f(0) = cos(0) = 1
    f_x += 1;

    // 3. Terim: -(2^2 * x^2) / 2!
    f_x += -pow(2, 2) * pow(x, 2) / factorial(2);

    // 5. Terim: (2^4 * x^4) / 4!
    f_x += pow(2, 4) * pow(x, 4) / factorial(4);

    printf("cos(%.1lf * 2) ~ %.6lf\n", x, f_x);

    return 0;
}
