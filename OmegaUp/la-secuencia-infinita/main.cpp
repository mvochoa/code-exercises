#include<stdio.h>
using namespace std;

int main() {
    int n = 0, op = -1;
    scanf("%d", &n);

    for (int cont = 0, serie = 1; cont < n; cont++, serie+=op) {
        printf("%d ", serie);
        if (serie == 5 || serie == 1) {
            op *= -1;
        }
    }
    return 0;
}
