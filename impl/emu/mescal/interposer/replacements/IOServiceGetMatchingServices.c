#include "IOServiceGetMatchingServices.h"

kern_return_t IOServiceGetMatchingServices(mach_port_t mainPort, CFDictionaryRef matching, io_iterator_t *existing) {
    return kIOReturnSuccess;

    (void) mainPort;
    (void) matching;
    (void) existing;
}
