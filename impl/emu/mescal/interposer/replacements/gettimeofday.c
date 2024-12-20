#include "gettimeofday.h"

int gettimeofday(struct timeval *restrict tp, void *restrict tzp) {
    if (tp) {
        tp->tv_sec = 1717000000;
        tp->tv_usec = 0;
    }

    if (tzp) {
        ((struct timezone *) tzp)->tz_minuteswest = 0;
        ((struct timezone *) tzp)->tz_dsttime = 0;
    }

    return 0;
}
