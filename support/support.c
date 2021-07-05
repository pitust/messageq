#include <stdint.h>

void outb(uint16_t port, uint8_t val)
{
    asm volatile("outb %0, %1"
                 :
                 : "a"(val), "Nd"(port));
}
// todo: maybe just make a semihosting api here and an lldb script?
void putchar(char c)
{
    outb(0xe9, c);
}
void puts(const char* s) {
    while (*s) putchar(*s++);
    putchar('\n');
}
void abort()
{
    puts("abort()");
    while (1)
    {
    }
}
static char* envi_internal = (char*)0;
char** environ = &envi_internal;


void* memcpy(uint8_t* restrict dst, const uint8_t* restrict src, uintptr_t n) {
    for (uintptr_t i = 0;i < n;i++) dst[i] = src[i];
    return dst;
}
void* memmove(uint8_t* restrict dst, const uint8_t* restrict src, uintptr_t n) {
    if (dst < src) {
        src += n - 1; dst += n - 1;
        for (uintptr_t i = 0; i < n; i++) dst[n - i] = src[n - i];
        return dst;
    }
    for (uintptr_t i = 0;i < n;i++) dst[i] = src[i];
    return dst;
}

/// memset - fill memory with a constant byte
///
/// The memset() function fills the first len bytes of the memory 
/// area pointed to by mem with the constant byte data.
///
/// The memset() function returns a pointer to the memory area mem.
uint8_t* memset(uint8_t* mem, int data, uintptr_t len) {
    for (uintptr_t i = 0; i < len; i++)
        mem[i] = (uint8_t)data;
    return mem;
}

uint64_t strlen(const char* s) {
    uint64_t i = 0;
    while (*s) {
        i++;
        s++;
    }
    return i;
}

int main(int argc, const char** argv);

void do_call_onstack(uint8_t* stack, void(*fn)(uint64_t), uint64_t arg);

extern char _stack_top[];

struct timespec {
	uint64_t tv_sec;  // time_t: follows the platform bitness
	uint64_t tv_nsec; // long: on Linux and macOS, follows the platform bitness
};
void clock_gettime(int i, struct timespec* tv) {
    tv->tv_sec = 0;
    tv->tv_nsec = 0;
}

void usleep() {}

void kinit(uint64_t v) {
    const char* arr[2] = {"messageq-kernel", 0};
    main(1, arr);
    puts("End of code!");
    abort();
}