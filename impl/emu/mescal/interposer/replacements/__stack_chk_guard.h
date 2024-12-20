//
//  __stack_chk_guard.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2024-12-13.
//

#ifndef ___STACK_CHK_GUARD_H
#define ___STACK_CHK_GUARD_H

#ifdef __stack_chk_guard
#undef __stack_chk_guard
#endif	/* __stack_chk_guard */

__attribute__ ((noinline)) void __stack_chk_guard(void);

#endif	/* ___STACK_CHK_GUARD_H */
