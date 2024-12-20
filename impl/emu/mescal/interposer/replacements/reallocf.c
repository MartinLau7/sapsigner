#include "realloc.h"

#include "reallocf.h"

void *reallocf(void *ptr, size_t size) {
    return realloc(ptr, size);
}
