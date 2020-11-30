#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main()
{
    srand(time(NULL));
    int target = rand() % 100 + 1;

    printf("Let's play a guessing game.\nI've picked a number between 1 and 100, try to guess it. You have 6 tries.\n");

    for (int x = 0; x < 6; x++) {
        printf("This is try %d\n", x+1);
        int guess;
        int check = scanf("%d", &guess);
        if (check != 1) {
            printf("Please enter only one number.");
            return 0;
        }
        if (guess < target) {
            printf("Too low!\n");
        } else if (guess > target) {
            printf("Too high!\n");
        } else {
            printf("You're right!\n");
            return 0;
        }
    }

    printf("Sorry, you failed. It was %d.\n", target);
}
