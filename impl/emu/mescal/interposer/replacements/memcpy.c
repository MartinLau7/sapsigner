#include "memmove.h"

#include "memcpy.h"

void *memcpy(void *restrict dst, const void *restrict src, size_t n) {
    return memmove(dst, src, n);
}
