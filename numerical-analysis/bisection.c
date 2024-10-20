#include <stdio.h>
#include <math.h>

// Hesaplamak istediğimiz fonksiyon f(x) = 2^x - 5x + 2 = 0
double f(double x) {
    return pow(2, x) - 5 * x + 2;
}

// Bisection yöntemi ile kök bulan fonksiyon
void bisection(double a, double b, double epsilon) {
    double c; // Orta nokta
    int iteration = 0; // Kaç iterasyon yapıldığını takip etmek için

    // f(a) ve f(b) işaretlerinin farklı olması gerekir
    if (f(a) * f(b) >= 0) {
        printf("Hata: f(a) ve f(b) işaretleri aynı olamaz.\n");
        return;
    }

    // Bisection işlemi, |b - a| > epsilon olana kadar devam eder
    while ((b - a) >= epsilon) {
        // Yeni orta noktayı hesapla
        c = (a + b) / 2;

        // İterasyonlar
        printf("Iterasyon %d: c = %.4lf, f(c) = %.4lf\n", ++iteration, c, f(c));

        // f(c) = 0 ise, c köktür
        if (f(c) == 0.0) {
            break;
        }

        // f(c) ve f(a) işaretleri aynı ise, a'yı güncelle
        else if (f(c) * f(a) < 0) {
            b = c;
        }

        // f(c) ve f(b) işaretleri aynı ise, b'yi güncelle
        else {
            a = c;
        }
    }

    printf("Yaklasik kok: %.4lf\n", c);
}

int main() {
    double a = 0, b = 1;  // Başlangıç aralığı [0, 1]
    double epsilon = 0.1; // Hedef doğruluk

    bisection(a, b, epsilon);

    return 0;
}
