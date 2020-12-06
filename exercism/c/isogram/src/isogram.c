#include "isogram.h"
#include <stddef.h>
#include <ctype.h>

bool is_isogram(const char phrase[]) {
    if (phrase == NULL) {
        return false;
    }

    int chars[26] = {0};
    const char* l = phrase;

    for (; *l != 0; l++) {
        if (!isalpha(*l)) {
            continue;
        }

        int c = tolower(*l) - 'a';
        if (++chars[c] > 1) {
            return false;
        }
    }

    return true;
}
