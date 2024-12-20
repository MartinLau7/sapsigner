//
//  DASessionCreate.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-11-29.
//

#ifndef _DASESSIONCREATE_H
#define _DASESSIONCREATE_H

#include <DiskArbitration/DASession.h>

#ifdef DASessionCreate
#undef DASessionCreate
#endif	/* DASessionCreate */

__attribute__ ((noinline)) DASessionRef DASessionCreate(CFAllocatorRef allocator);

#endif	/* _DASESSIONCREATE_H */
