#include "matching_brackets.h"

#define MAGIC_STACK_SIZE 20

bool is_paired(const char* input)
{
    char stack[MAGIC_STACK_SIZE] = {'\0'};
    int p = -1;
    char c;

    while ((c = *input++) != '\0') {
        switch (c) {
            case '{':
                stack[++p] = '}';
                break;
            case '[':
                stack[++p] = ']';
                break;
            case '(':
                stack[++p] = ')';
                break;
            case '}':
            case ']':
            case ')':
                if (p >= 0 && stack[p] == c) {
                    p--;
                } else {
                    return false;
                }
        }
    }

    return p == -1;
}
