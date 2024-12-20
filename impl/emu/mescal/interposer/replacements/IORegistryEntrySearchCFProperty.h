//
//  IORegistryEntrySearchCFProperty.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _IOREGISTRYENTRYSEARCHCFPROPERTY_H
#define _IOREGISTRYENTRYSEARCHCFPROPERTY_H

#include <IOKit/IOKitLib.h>

#ifdef IORegistryEntrySearchCFProperty
#undef IORegistryEntrySearchCFProperty
#endif	/* IORegistryEntrySearchCFProperty */

__attribute__ ((noinline)) CFTypeRef IORegistryEntrySearchCFProperty(io_registry_entry_t entry, const io_name_t plane, CFStringRef key, CFAllocatorRef allocator, IOOptionBits options);

#endif	/* _IOREGISTRYENTRYSEARCHCFPROPERTY_H */
