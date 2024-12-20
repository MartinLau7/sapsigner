#include "malloc.h"

#include "calloc.h"

void *calloc(size_t count, size_t size) {
    return malloc(count * size);
}
