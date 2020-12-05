#include <stdlib.h>
#include <ctype.h>
#include <stddef.h>
#include <string.h>
#include "acronym.h"

char *abbreviate(const char *phrase)
{
    if (phrase == NULL || *phrase == '\0') {
        return NULL;
    }

    char *ret = malloc(strlen(phrase) * sizeof(char));
    char c, prev = ' ';
    int retp = 0;

    while ((c = *phrase++) != '\0') {
        if (prev == ' ' || prev == '-') {
            ret[retp++] = (char) toupper(c);
        }
        prev = c;
    }

    ret[retp] = '\0';
    return ret;
}
