// #include "strcmp.h"

#include "dlopen.h"

void *dlopen(const char* path, int mode) {
    if (
        false
        // || !strcmp(path, "/System/Library/PrivateFrameworks/CoreFP.framework/CoreFP")
        || (path[0] == '/' && path[1] == 'S' && path[2] == 'y' && path[3] == 's' && path[4] == 't' && path[5] == 'e' && path[6] == 'm' && path[7] == '/' && path[8] == 'L' && path[9] == 'i' && path[10] == 'b' && path[11] == 'r' && path[12] == 'a' && path[13] == 'r' && path[14] == 'y' && path[15] == '/' && path[16] == 'P' && path[17] == 'r' && path[18] == 'i' && path[19] == 'v' && path[20] == 'a' && path[21] == 't' && path[22] == 'e' && path[23] == 'F' && path[24] == 'r' && path[25] == 'a' && path[26] == 'm' && path[27] == 'e' && path[28] == 'w' && path[29] == 'o' && path[30] == 'r' && path[31] == 'k' && path[32] == 's' && path[33] == '/' && path[34] == 'C' && path[35] == 'o' && path[36] == 'r' && path[37] == 'e' && path[38] == 'F' && path[39] == 'P' && path[40] == '.' && path[41] == 'f' && path[42] == 'r' && path[43] == 'a' && path[44] == 'm' && path[45] == 'e' && path[46] == 'w' && path[47] == 'o' && path[48] == 'r' && path[49] == 'k' && path[50] == '/' && path[51] == 'C' && path[52] == 'o' && path[53] == 'r' && path[54] == 'e' && path[55] == 'F' && path[56] == 'P' && path[57] == '\0')
    ) {
        return (void *) ~0;
    }

    return NULL;

    (void) mode;
}
