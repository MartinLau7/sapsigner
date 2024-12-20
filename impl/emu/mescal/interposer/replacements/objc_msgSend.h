//
//  objc_msgSend.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef _OBJC_MSGSEND_H
#define _OBJC_MSGSEND_H

#ifdef OBJC_OLD_DISPATCH_PROTOTYPES
#undef OBJC_OLD_DISPATCH_PROTOTYPES
#endif	/* OBJC_OLD_DISPATCH_PROTOTYPES */
#define OBJC_OLD_DISPATCH_PROTOTYPES 1

#include <objc/message.h>

#ifdef objc_msgSend
#undef objc_msgSend
#endif	/* objc_msgSend */

__attribute__ ((noinline)) id objc_msgSend(id self, SEL op, ...);

#endif	/* _OBJC_MSGSEND_H */
