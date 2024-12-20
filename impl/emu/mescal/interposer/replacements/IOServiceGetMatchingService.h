//
//  IOServiceGetMatchingService.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _IOSERVICEGETMATCHINGSERVICE_H
#define _IOSERVICEGETMATCHINGSERVICE_H

#include <IOKit/IOKitLib.h>

#ifdef IOServiceGetMatchingService
#undef IOServiceGetMatchingService
#endif	/* IOServiceGetMatchingService */

__attribute__ ((noinline)) io_service_t IOServiceGetMatchingService(mach_port_t mainPort, CFDictionaryRef matching);

#endif	/* _IOSERVICEGETMATCHINGSERVICE_H */
