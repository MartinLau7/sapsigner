//
//  IORegistryEntryGetParentEntry.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _IOREGISTRYENTRYGETPARENTENTRY_H
#define _IOREGISTRYENTRYGETPARENTENTRY_H

#include <IOKit/IOKitLib.h>

#ifdef IORegistryEntryGetParentEntry
#undef IORegistryEntryGetParentEntry
#endif	/* IORegistryEntryGetParentEntry */

__attribute__ ((noinline)) kern_return_t IORegistryEntryGetParentEntry(io_registry_entry_t entry, const io_name_t plane, io_registry_entry_t *parent);

#endif	/* _IOREGISTRYENTRYGETPARENTENTRY_H */
