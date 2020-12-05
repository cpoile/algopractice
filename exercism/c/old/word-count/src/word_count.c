#include <ctype.h>
#include <stdbool.h>
#include <string.h>

#include "word_count.h"

char* strip(char* text, char* sep);

int word_count(const char* input_text, word_count_word_t* words)
{
    int count = 0;

    // copy input so we can modify it
    char input[MAX_WORDS * (MAX_WORD_LENGTH + 1)];
    strcpy(input, input_text);

    // clear the allocated words structure
    memset(words, 0, sizeof(word_count_word_t) * MAX_WORDS);
    word_count_word_t* retp = words;

    // lowercase the string
    char* t = input;
    while ((*t = tolower(*t))) ++t;

    char* sep = " ,\n:!&@$%^&.";

    for (char* word = strtok(input, sep); word; word = strtok(NULL, sep)) {
        if (strlen(word) > MAX_WORD_LENGTH) {
            return EXCESSIVE_LENGTH_WORD;
        }

        bool recorded = false;
        word = strip(word, "'");
        for (int i = 0; i < count; i++) {
            if (strcmp(word, words[i].text) == 0) {
                words[i].count++;
                recorded = true;
                break;
            }
        }
        if (!recorded) {
            if (count == 20) {
                return EXCESSIVE_NUMBER_OF_WORDS;
            }

            strcpy(retp->text, word);
            retp->count = 1;
            retp++;
            count++;
        }
    }

    return count;
}

// cins returns true if c is in s
bool cins(char c, char* s)
{
    char* t = s;
    while (*t != '\0') {
        if (c == *t) {
            return true;
        }
        t++;
    }
    return false;
}

// strip returns a pointer to the text stripped of sep from prefix and postfix
// TODO: edge cases not thoroughly tested, don't depend on this generally
char* strip(char* text, char* sep)
{
    char* beg = text;
    while (cins(*beg, sep)) ++beg;

    // get the end char
    char* end = beg;
    while (*end != '\0') ++end;

    // strip the suffix chars, but don't go past beg
    while (end != beg && --end != beg) {
        if (!cins(*end, sep)) {  // c is not in separator, stop stripping.
            ++end;
            *end = '\0';
            break;
        }
    }

    return beg;
}
