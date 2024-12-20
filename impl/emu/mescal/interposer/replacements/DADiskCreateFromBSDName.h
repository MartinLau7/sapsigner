//
//  DADiskCreateFromBSDName.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _DADISKCREATEFROMBSDNAME_H
#define _DADISKCREATEFROMBSDNAME_H

#include <DiskArbitration/DADisk.h>

#ifdef DADiskCreateFromBSDName
#undef DADiskCreateFromBSDName
#endif	/* DADiskCreateFromBSDName */

__attribute__ ((noinline)) DADiskRef DADiskCreateFromBSDName(CFAllocatorRef allocator, DASessionRef session, const char *name);

#endif	/* _DADISKCREATEFROMBSDNAME_H */
