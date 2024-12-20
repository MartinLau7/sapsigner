#include "malloc_good_size.h"

size_t malloc_good_size(size_t size) {
    return size + ((_Alignof(max_align_t) - size) % _Alignof(max_align_t));
}
