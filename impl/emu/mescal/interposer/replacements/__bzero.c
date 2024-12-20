#include "__bzero.h"

void __bzero(void *s, size_t n) {
    for (unsigned char *it = s; it != (unsigned char *) s + n; ++it) {
        *it = 0;
    }
}
