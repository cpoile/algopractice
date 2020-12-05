#include <stdio.h>
#include "word_count.h"

int main2() {
    //char tmp[MAX_WORD_LENGTH+1] = "'test'\0";
    //char* t = strip(tmp, "'");

    word_count_word_t words[MAX_WORDS];
    int c = word_count("one,two,three", words);

    printf("Num words: %d\n", c);

    return 0;
}
