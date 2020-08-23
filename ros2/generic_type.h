
#ifndef GENERIC__TYPE_SUPPORT_H_
#define GENERIC__TYPE_SUPPORT_H_

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "rosidl_generator_c/message_type_support_struct.h"

// ROSIDL STRING TYPE
#include "rosidl_generator_c/string.h"

// Forward declare the get type support functions for this type.

const rosidl_message_type_support_t *
  get_generic_type();


// Struct definition of message
typedef struct Generic
{
  rosidl_generator_c__String data;
} Generic;

// Message functions

// Init
bool Generic__init(Generic * msg);

// Fini
void Generic__fini(Generic * msg);

// Create
Generic * Generic__create();

// Destroy
void Generic__destroy(Generic * msg);



#endif  // GENERIC__TYPE_SUPPORT_H_



