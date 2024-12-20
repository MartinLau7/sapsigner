//
//  IOServiceGetMatchingServices.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _IOSERVICEGETMATCHINGSERVICES_H
#define _IOSERVICEGETMATCHINGSERVICES_H

#include <IOKit/IOKitLib.h>

#ifdef IOServiceGetMatchingServices
#undef IOServiceGetMatchingServices
#endif	/* IOServiceGetMatchingServices */

__attribute__ ((noinline)) kern_return_t IOServiceGetMatchingServices(mach_port_t mainPort, CFDictionaryRef matching, io_iterator_t *existing);

#endif	/* _IOSERVICEGETMATCHINGSERVICES_H */
