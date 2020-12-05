#include <stdint.h>
#include "pangram.h"
#include <ctype.h>

bool is_pangram(const char *sentence) {
    if (sentence == NULL) {
        return false;
    }
    bool chars[26] = {false};
    char c;

    while ((c = *sentence++) != '\0') {
        if (isalpha(c)) {
            int idx = tolower(c) - 'a';
            chars[idx] = true;
        }
    }

    for (int i = 0; i < 26; i++) {
        if (!chars[i]) {
            return false;
        }
    }

    return true;
}
