#include "strncmp.h"

int strncmp(const char *s1, const char *s2, size_t n) {
    while (n--) {
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
