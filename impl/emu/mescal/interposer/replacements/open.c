// #include "strcmp.h"

#include "open.h"

int open(const char *path, int oflag, ...) {
    if (
        false
        // || !strcmp(path, "./../CoreFP.icxs")
        || (path[0] == '.' && path[1] == '/' && path[2] == '.' && path[3] == '.' && path[4] == '/' && path[5] == 'C' && path[6] == 'o' && path[7] == 'r' && path[8] == 'e' && path[9] == 'F' && path[10] == 'P' && path[11] == '.' && path[12] == 'i' && path[13] == 'c' && path[14] == 'x' && path[15] == 's' && path[16] == '\0')
    ) {
        return 0;
    }

    return -1;

    (void) path;
    (void) oflag;
}
