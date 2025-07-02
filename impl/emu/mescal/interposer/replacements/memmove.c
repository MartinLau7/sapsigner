#include "memmove.h"

void *memmove(void *restrict dst, const void *restrict src, size_t n) {
    if ((size_t) dst < (size_t) src) {
        unsigned char *it_dst = dst;
        const unsigned char *it_src = src;
        while (it_src != (const unsigned char *) src + n) {
            *it_dst++ = *it_src++;
        }
    } else {
        unsigned char *it_dst = (unsigned char *) dst + n;
        const unsigned char *it_src = (const unsigned char *) src + n;
        while (it_src != src) {
            *--it_dst = *--it_src;
        }
    }
    
    return dst;
}
