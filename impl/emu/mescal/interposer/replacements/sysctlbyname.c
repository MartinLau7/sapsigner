#include "sysctlbyname.h"

int sysctlbyname(const char *name, void *oldp, size_t *oldlenp, void *newp, size_t newlen) {
    *oldlenp = 0;

    return 0;

    (void) name;
    (void) oldp;
    (void) newp;
    (void) newlen;
}
