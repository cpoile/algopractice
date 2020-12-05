#include "acronym.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

char* abbreviate(const char* phrase)
{
    if (phrase == NULL || strlen(phrase) == 0) {
        return NULL;
    }

    char* ret = malloc(strlen(phrase) * sizeof(char));
    char prev = ' ';
    char c;
    int cur = 0;
    while ((c = *phrase++) != '\0') {
        if ((prev == ' ' || prev == '-' || prev == '_') && isalpha(c)) {
            ret[cur++] = (char) toupper(c);
        }
        prev = c;
    }

    ret[cur] = 0;
    return ret;
}
