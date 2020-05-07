#include <stdio.h>
#include <string>
using namespace std;

string babilonico(int n) {
    string num = "";
    int u = n%10;
    n /= 10;
    while (n > 0) {
        num += "L";
        n--;
    }
    while (u > 0) {
        num += "I";
        u--;
    }

    return num;
}

void camino(int n) {
    string num = babilonico(n%60);
    n = (int)n/60;

    while (n > 0) {
        num += "." + babilonico(n%60);
        n = (int)n/60;
    }
    printf("%s", num.c_str());
}

int main() {
    int n;
    scanf("%d", &n);

    camino(n);
    return 0;
}