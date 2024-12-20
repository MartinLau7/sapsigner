#include <stdbool.h>

#include "strcmp.h"

int strcmp(const char *s1, const char *s2) {
    while (true) {
        int diff = *s1 - *s2;
        if (diff) {
            return diff;
        }

        if (*s1 == '\0') {
            break;
        }

        ++s1;
        ++s2;
    }

    return 0;
}
