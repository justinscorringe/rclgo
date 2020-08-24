
#ifndef GENERIC__TYPE_SUPPORT_H_
#define GENERIC__TYPE_SUPPORT_H_

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "rosidl_generator_c/message_type_support_struct.h"

// ROSIDL STRING TYPE
#include "rosidl_generator_c/string.h"

// Forward declare the get type support functions for this type.
// Go member is a bridging introspection member type
typedef struct GoMember
{
  const char * name_;
  uint8_t type_id_;
  const rosidl_message_type_support_t * members_;
  bool is_array_;
  size_t array_size_;
} GoMember;


const rosidl_message_type_support_t
  new_generic_type(const char* message_namespace_, 
    const char* message_name_, 
    uint32_t member_count_,
    GoMember go_members_[]);

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



