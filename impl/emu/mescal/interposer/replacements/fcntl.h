//
//  fcntl.h
//  SAPSigner
//
//  Created by Pedro Tôrres on 2025-06-28.
//

#ifndef _FCNTL_H
#define _FCNTL_H

#ifdef fcntl
#undef fcntl
#endif	/* fcntl */

__attribute__ ((noinline)) int fcntl(int fildes, int cmd, ...);

#endif	/* _FCNTL_H */
