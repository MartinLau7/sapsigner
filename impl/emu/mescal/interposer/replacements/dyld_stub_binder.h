//
//  dyld_stub_binder.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-14.
//

#ifndef _DYLD_STUB_BINDER_H
#define _DYLD_STUB_BINDER_H

#ifdef dyld_stub_binder
#undef dyld_stub_binder
#endif	/* dyld_stub_binder */

__attribute__ ((noinline)) void dyld_stub_binder(void) __asm__ ("dyld_stub_binder");

#endif	/* _DYLD_STUB_BINDER_H */
