#include <stdbool.h>
#include <stdint.h>

typedef short myint;
typedef short unusedTypedef;
int add(int a, int b);
int unusedFunction(void);
typedef int (*binop_t) (int, int);
int doCallback(int a, int b, binop_t cb);
typedef int * intPointer;
void store(int value, int *ptr);

# define CONST_INT 5
# define CONST_INT2 5llu
# define CONST_FLOAT 5.8
# define CONST_FLOAT2 5.8f
# define CONST_CHAR 'c'
# define CONST_STRING "defined string"

// this signature should not be included by CGo
void unusedFunction2(int x, __builtin_va_list args);

typedef struct collection {
	short         s;
	long          l;
	float         f;
	unsigned char c;
} collection_t;

struct point {
	int x;
	int y;
};

// linked list
typedef struct list_t {
	int           n;
	struct list_t *next;
} list_t;

typedef union joined {
	myint s;
	float f;
	short data[3];
} joined_t;
void unionSetShort(short s);
void unionSetFloat(float f);
void unionSetData(short f0, short f1, short f2);

// test globals and datatypes
extern int global;
extern int unusedGlobal;
extern bool globalBool;
extern bool globalBool2;
extern float globalFloat;
extern double globalDouble;
extern _Complex float globalComplexFloat;
extern _Complex double globalComplexDouble;
extern _Complex double globalComplexLongDouble;
extern char globalChar;
extern void *globalVoidPtrSet;
extern void *globalVoidPtrNull;
extern int64_t globalInt64;
extern collection_t globalStruct;
extern int globalStructSize;
extern short globalArray[3];
extern joined_t globalUnion;
extern int globalUnionSize;

// test duplicate definitions
int add(int a, int b);
extern int global;
