//
//  DADiskCopyDescription.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _DADISKCOPYDESCRIPTION_H
#define _DADISKCOPYDESCRIPTION_H

#include <DiskArbitration/DADisk.h>

#ifdef DADiskCopyDescription
#undef DADiskCopyDescription
#endif	/* DADiskCopyDescription */

__attribute__ ((noinline)) CFDictionaryRef DADiskCopyDescription(DADiskRef disk);

#endif	/* _DADISKCOPYDESCRIPTION_H */
