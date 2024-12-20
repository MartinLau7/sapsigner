# libmescal.dylib

## Variants

### Only Supports Regular Session

1. Mac App Store Update for OS X Mountain Lion

> `CommerceKit` from
> https://swcdn.apple.com/content/downloads/11/62/041-88157-A_9MV302JSGJ/nut08bz4n2byymwkltjsxcsans8pwndovm/AppStoreUpdate.pkg

### Also Supports Prime Session

1. OS X Mavericks 10.9.1 Update

> `CommerceKit` from
> https://updates.cdn-apple.com/2019/cert/041-88140-20191011-250758ef-d633-428d-afa8-7334f518e125/OSXUpd10.9.1.dmg

2. Mac App Store Update for OS X El Capitan

> `com.apple.CommerceKit.TransactionService` from
> https://swcdn.apple.com/content/downloads/15/08/041-89046-A_F9IVP0KUUH/mmgksihi7nj85zzrdse8l5wpw5ruf1usv8/MacAppStoreUpdateforElCapitan.pkg
>
> `CoreFP` and `CoreFP.icxs` from
> https://swcdn.apple.com/content/downloads/30/28/041-88379-A_SDNTMOXIL3/g4380v161gox8pqqcyrpqijk1s79zejl9d/CoreFP.pkg

## Symbols

| Original Name            | Mangled Name   |
|--------------------------|----------------|
| `FairPlayDisposeStorage` | `jEHf8Xzsv8K`  |
| `FairPlaySAPExchange`    | `Mib5yocT`     |
| `FairPlaySAPInit`        | `cp2g1b9ro`    |
| `FairPlaySAPPrime`       | `jfkdDAjba3jd` |
| `FairPlaySAPSign`        | `Fc3vhtJDvr`   |
| `FairPlaySAPTeardown`    | `IPaI1oem5iL`  |
| `FairPlaySAPVerify`      | `gLg1CWr7p`    |

## Header

```c
struct FPSAPContextOpaque_;

struct FairPlayHWInfo_ {
    unsigned int IDLength;
    unsigned char ID[20];
};

enum FairPlaySAPExchangeVersion_ {
    FairPlaySAPExchangeVersion_Regular = 200,
    FairPlaySAPExchangeVersion_Prime = 210,
};

enum FairPlaySAPPrime$_ {
    FairPlaySAPPrime$_$ = 100,
};

#define FairPlaySAPInit cp2g1b9ro
extern int FairPlaySAPInit(struct FPSAPContextOpaque_ **ctx, const struct FairPlayHWInfo_ *hw_info);

#define FairPlaySAPExchange Mib5yocT
extern int FairPlaySAPExchange(enum FairPlaySAPExchangeVersion_ version, const struct FairPlayHWInfo_ *hw_info, struct FPSAPContextOpaque_ *ctx, const char *in_buf, unsigned int in_len, char **out_buf, unsigned int *out_len, int *return_code);

#define FairPlayDisposeStorage jEHf8Xzsv8K
extern int FairPlayDisposeStorage(void *ptr);

#define FairPlaySAPSign Fc3vhtJDvr
extern int FairPlaySAPSign(struct FPSAPContextOpaque_ *ctx, const char *in_buf, unsigned int in_len, char **out_buf, unsigned int *out_len);

#define FairPlaySAPPrime jfkdDAjba3jd
extern int FairPlaySAPPrime(struct FPSAPContextOpaque_ *ctx, enum FairPlaySAPPrime$_ $, const char *in_buf, unsigned int in_len, char **out_buf, unsigned int *out_len);

#define FairPlaySAPVerify gLg1CWr7p
extern int FairPlaySAPVerify(struct FPSAPContextOpaque_ *ctx, const char *in_buf, unsigned int in_len, char **out_buf, unsigned int *out_len);

#define FairPlaySAPTeardown IPaI1oem5iL
extern int FairPlaySAPTeardown(struct FPSAPContextOpaque_ *ctx);
```
