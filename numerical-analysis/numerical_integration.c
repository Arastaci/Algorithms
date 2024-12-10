#include <stdio.h>
#include <math.h>

double f(double x) {
    if (fabs(x - 0.1) < 1e-10) return -0.99;
    if (fabs(x - 0.2) < 1e-10) return -0.96;
    if (fabs(x - 0.3) < 1e-10) return -0.91;
    if (fabs(x - 0.4) < 1e-10) return -0.84;
    if (fabs(x - 0.5) < 1e-10) return -0.75;
}

double numericIntegral(double a, double b, int n) {
    if (n % 2 != 0) {
        printf("Simpson yöntemi için n çift olmalıdır.\n");
        return -1;
    }
    
    double h = (b - a) / n;
    double sum = f(a) + f(b);   // İlk ve son nokta 1 kez toplanır
    
    // Ara noktalar için
    for (int i = 1; i < n; i++) {
        double x = a + i * h;
        if (i % 2 == 0)
            sum += 2 * f(x);  // Çift indeksli noktalar 2 ile çarpılırarak toplanır
        else
            sum += 4 * f(x);  // Tek indeksli noktalar 4 ile çarpılırarak toplanır
    }
    
    return (h / 3) * sum;   //Formül gereği h/3 ile çarpılır
}

int main() {
    double a = 0.1, b = 0.5;
    int n;

    printf("Integral sınırları: [%.1f, %.1f]\n", a, b);
    printf("Tablo değerleri:\n");
    printf("x     : 0.1  0.2  0.3  0.4  0.5\n");
    printf("f(x)  : -0.99 -0.96 -0.91 -0.84 -0.75\n\n");
    
    printf("n değerini giriniz (2 veya 4 olmalı): ");
    scanf("%d", &n);

    if (n % 2 != 0) {
        printf("Hata! n çift olmalı.\n");
        return 1;
    }

    double result = numericIntegral(a, b, n);
    if (result != -1) {
        printf("n = %d için integral değeri: %.6f\n", n, result);
    }

    return 0;
}
