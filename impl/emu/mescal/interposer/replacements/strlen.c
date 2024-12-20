#include <stdbool.h>

#include "strlen.h"

size_t strlen(const char *s) {
    for (size_t i = 0; true; ++i) {
        if (s[i] == '\0') {
            return i;
        }
    }
}
