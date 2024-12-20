//
//  IORegistryEntryCreateCFProperty.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _IOREGISTRYENTRYCREATECFPROPERTY_H
#define _IOREGISTRYENTRYCREATECFPROPERTY_H

#include <IOKit/IOKitLib.h>

#ifdef IORegistryEntryCreateCFProperty
#undef IORegistryEntryCreateCFProperty
#endif	/* IORegistryEntryCreateCFProperty */

__attribute__ ((noinline)) CFTypeRef IORegistryEntryCreateCFProperty(io_registry_entry_t entry, CFStringRef key, CFAllocatorRef allocator, IOOptionBits options);

#endif	/* _IOREGISTRYENTRYCREATECFPROPERTY_H */
