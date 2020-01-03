#include <ctype.h>

#include "pangram.h"

bool is_pangram(const char* sentence)
{
    if (!sentence)
        return false;

    bool counts[26] = {0};

    for (; *sentence != '\0'; ++sentence) {
        if (isalpha(*sentence)) {
            counts[tolower(*sentence) - 'a'] += 1;
        }
    }

    for (size_t i = 0; i < 26; ++i) {
        if (counts[i] == 0)
            return false;
    }
    return true;
}
