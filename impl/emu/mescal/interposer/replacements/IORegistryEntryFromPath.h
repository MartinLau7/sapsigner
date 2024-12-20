//
//  IORegistryEntryFromPath.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _IOREGISTRYENTRYFROMPATH_H
#define _IOREGISTRYENTRYFROMPATH_H

#include <IOKit/IOKitLib.h>

#ifdef IORegistryEntryFromPath
#undef IORegistryEntryFromPath
#endif	/* IORegistryEntryFromPath */

__attribute__ ((noinline)) io_registry_entry_t IORegistryEntryFromPath(mach_port_t mainPort, const io_string_t path);

#endif	/* _IOREGISTRYENTRYFROMPATH_H */
