#include "malloc_good_size.h"

#include "malloc_size.h"

size_t malloc_size(const void *ptr) {
    size_t aligned_size_t_size = malloc_good_size(sizeof(size_t));
    return *((const size_t *) ((const unsigned char *) ptr - aligned_size_t_size));
}
