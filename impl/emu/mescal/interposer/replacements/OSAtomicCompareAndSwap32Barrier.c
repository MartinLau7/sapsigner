#include "OSAtomicCompareAndSwap32Barrier.h"

bool OSAtomicCompareAndSwap32Barrier(int32_t oldValue, int32_t newValue, volatile int32_t *theValue) {
    if (*theValue != oldValue)
        return false;

    *theValue = newValue;
    return true;
}
