#include "memcpy.h"

#include "read.h"

ssize_t read(int fildes, void *buf, size_t nbyte) {
    if (fildes == 0) { // ./../CoreFP.icxs
#pragma clang diagnostic push
#pragma clang diagnostic ignored "-Wdollar-in-identifier-extension"
        static unsigned char *volatile CoreFP$ICXS_data = NULL;
        static volatile size_t CoreFP$ICXS_size = 0;

        ssize_t ret = CoreFP$ICXS_size < nbyte ? CoreFP$ICXS_size : nbyte;
        memcpy(buf, CoreFP$ICXS_data, ret);

        CoreFP$ICXS_size -= ret;
        CoreFP$ICXS_data += ret;
#pragma clang diagnostic pop

        return ret;
    }

    return -1;
}
