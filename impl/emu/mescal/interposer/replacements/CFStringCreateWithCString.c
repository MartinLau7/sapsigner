// #include "strcmp.h"

#include "CFStringCreateWithCString.h"

CFStringRef CFStringCreateWithCString(CFAllocatorRef alloc, const char *cStr, CFStringEncoding encoding) {
    if (
        false
        // || !strcmp(cStr, "IOPlatformSerialNumber")
        || (cStr[0] == 'I' && cStr[1] == 'O' && cStr[2] == 'P' && cStr[3] == 'l' && cStr[4] == 'a' && cStr[5] == 't' && cStr[6] == 'f' && cStr[7] == 'o' && cStr[8] == 'r' && cStr[9] == 'm' && cStr[10] == 'S' && cStr[11] == 'e' && cStr[12] == 'r' && cStr[13] == 'i' && cStr[14] == 'a' && cStr[15] == 'l' && cStr[16] == 'N' && cStr[17] == 'u' && cStr[18] == 'm' && cStr[19] == 'b' && cStr[20] == 'e' && cStr[21] == 'r' && cStr[22] == '\0')
        // || !strcmp(cStr, "IOPlatformUUID")
        || (cStr[0] == 'I' && cStr[1] == 'O' && cStr[2] == 'P' && cStr[3] == 'l' && cStr[4] == 'a' && cStr[5] == 't' && cStr[6] == 'f' && cStr[7] == 'o' && cStr[8] == 'r' && cStr[9] == 'm' && cStr[10] == 'U' && cStr[11] == 'U' && cStr[12] == 'I' && cStr[13] == 'D' && cStr[14] == '\0')
        // || !strcmp(cStr, "board-id")
        || (cStr[0] == 'b' && cStr[1] == 'o' && cStr[2] == 'a' && cStr[3] == 'r' && cStr[4] == 'd' && cStr[5] == '-' && cStr[6] == 'i' && cStr[7] == 'd' && cStr[8] == '\0')
    ) {
        return (void *) ~0;
    }

    return NULL;

    (void) alloc;
    (void) encoding;
}
