#include <stddef.h>
#include "isogram.h"

bool is_isogram(const char phrase[])
{
    if (phrase == NULL) {
        return false;
    }

    int counts[26] = {0};
    char c;
    while ((c = *phrase++) != '\0') {
        if (c >= 'A' && c <= 'Z') {
            counts[c - 'A'] += 1;
        } else if (c >= 'a' && c <= 'z') {
            counts[c - 'a'] += 1;
        }
    }
    for (size_t i = 0; i < 26; i++) {
        if (counts[i] > 1) {
            return false;
        }
    }
    return true;
}
