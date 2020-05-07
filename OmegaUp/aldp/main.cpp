#include <stdio.h>
#include <map>
using namespace std;

void haveKeys(map<int, int> doors, int keys[], int size) {
    for (int i = 0; i < size; i++) {
        if (doors.find(keys[i]) != doors.end()) {
            printf("%d ", doors[keys[i]]);
        } else {
            printf("0 ");
        }
    }
}

int main() {
    int n, m, item;
    map<int, int> doors;

    scanf("%d\n", &n);
    for (int i = 0; i < n; i++) {
        scanf("%d", &item);
        doors[item] = i+1;
    }

    scanf("%d\n", &m);
    int keys[m];
    for (int i = 0; i < m; i++) {
        scanf("%d", &keys[i]);
    }

    haveKeys(doors, keys, m);
}