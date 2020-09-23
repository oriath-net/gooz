#include "bitscan.h"

unsigned char _BitScanForward(unsigned long *index, unsigned long mask)
{
    *index = __builtin_ctz(mask);
    return mask != 0;
}

unsigned char _BitScanReverse(unsigned long *index, unsigned long mask)
{
    *index = 31 - __builtin_clz(mask);
    return mask != 0;
}
