#include <stdio.h>
#include <math.h>

// Fonksiyon tanımı: f(x) = x^3 - 5*x^2 - 2*x + 10
double f(double x) {
    return x*x*x - 5*x*x - 2*x + 10;
}

double calc(double a, double b, double tol) {
    double err, x0;

    while (1) {
        x0 = (a * f(b) - b * f(a)) / (f(b) - f(a));
        err = fabs(f(x0)); // Hata miktarı (f(x0) değerinin mutlak değeri)

        // Eğer hata toleransın altına indiyse işlemi bitir
        if (err <= tol) {
            break;
        }
        // Kök, [a, x0] aralığındaysa, b'yi güncelle
        else if (f(a) * f(x0) < 0) {
            b = x0;
        }
        // Kök, [x0, b] aralığındaysa, a'yı güncelle
        else {
            a = x0;
        }
    }
    return x0;
}

int main() {
    double a = 1.0, b = 3.0; 
    double tol = 0.01;        

    double root = calc(a, b, tol);

    printf("Kok: %.4f\n", root);
    return 0;
}
