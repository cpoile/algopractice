#include "word_count.h"
#include <ctype.h>
#include <string.h>


int addToList(word_count_word_t* words, char word[])
{
    for (int i = 0; i < MAX_WORDS; i++) {
        if (strcmp(word, words[i].text) == 0) {
            words[i].count++;
            return 0;
        } else if (words[i].count == 0) {
            strcpy_s(words[i].text, MAX_WORD_LENGTH + 1, word);
            words[i].count++;
            return 1;
        }
    }
    return EXCESSIVE_NUMBER_OF_WORDS;
}

int word_count(const char* input_text, word_count_word_t* words)
{
    char c;
    char cur[MAX_WORD_LENGTH + 1] = {0};
    int pos = 0;
    int numWords = 0;
    for (;;) {
        c = *input_text++;
        if (isalnum(c) || (c == '\'' && pos != 0)) {
            if (pos == MAX_WORD_LENGTH) {
                return EXCESSIVE_LENGTH_WORD;
            }
            cur[pos++] = (char) tolower(c);
        } else if (pos > 0) {
            if (numWords == 20) {
                return EXCESSIVE_NUMBER_OF_WORDS;
            }
            if (cur[pos - 1] == '\'') {
                cur[pos - 1] = '\0';
            } else {
                cur[pos] = '\0';
            }
            numWords += addToList(words, cur);
            pos = 0;
        }
        if (c == '\0') {
            break;
        }
    }

    return numWords;
}
