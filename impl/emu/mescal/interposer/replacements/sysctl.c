#include "sysctl.h"

int sysctl(int *name, u_int namelen, void *oldp, size_t *oldlenp, void *newp, size_t newlen) {
    return -1;

    (void) name;
    (void) namelen;
    (void) oldp;
    (void) oldlenp;
    (void) newp;
    (void) newlen;
}
