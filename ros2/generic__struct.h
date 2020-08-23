#ifndef GENERIC__STRUCT_H_
#define GENERIC__STRUCT_H_

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>


// Constants defined in the message

// Include directives for member types
// Member 'data'
#include "rosidl_generator_c/string.h"

typedef struct Generic
{
  rosidl_generator_c__String data;
} Generic;

typedef struct Generic__Sequence
{
  Generic * data;
  /// The number of valid items in data
  size_t size;
  /// The number of allocated items in data
  size_t capacity;
} Generic__Sequence;


#ifdef __cplusplus
}
#endif

#endif  // GENERIC__STRUCT_H_