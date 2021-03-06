
#ifndef GENERIC__TYPE_SUPPORT_H_
#define GENERIC__TYPE_SUPPORT_H_

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "rosidl_generator_c/message_type_support_struct.h"

// ROSIDL STRING TYPE
#include "rosidl_generator_c/string.h"

// Go member is a bridging introspection member type
typedef struct GoMember
{
  const char * name_;
  uint8_t type_id_;
  const rosidl_message_type_support_t * members_;
  size_t member_offset_;
  bool is_array_;
  size_t array_size_;
} GoMember;

typedef struct GoMembers
{
  size_t struct_size_;
  GoMember * member_array;
} GoMembers;

// Dynamic type

const rosidl_message_type_support_t
  new_generic_type(const char* message_namespace_, 
    const char* message_name_, 
    uint32_t member_count_,
    GoMembers * go_members_);


// Static type

// const rosidl_message_type_support_t *
//   get_generic_type();

// Static struct definition of message (std_msgs/String)
typedef struct Generic
{
  rosidl_generator_c__String data;
} Generic;

// Message functions

// Init
bool Generic__init(Generic * msg);

// Fini
void Generic__fini(Generic * msg);

// Create static
//Generic * Generic__create();

// Create dynamic
Generic * Generic__create(GoMembers * go_members_, uint32_t member_count_);

// Destroy
void Generic__destroy(Generic * msg);



#endif  // GENERIC__TYPE_SUPPORT_H_



