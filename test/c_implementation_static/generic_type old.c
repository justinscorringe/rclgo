#include <stddef.h>
#include <assert.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>
#include "rosidl_typesupport_introspection_c/field_types.h"
#include "rosidl_typesupport_introspection_c/identifier.h"
#include "rosidl_typesupport_introspection_c/message_introspection.h"
#include "generic_type.h"

// Include directives for member types
// Member `data`
#include "rosidl_generator_c/string_functions.h"

static rosidl_typesupport_introspection_c__MessageMember Generic__rosidl_typesupport_introspection_c__Generic_message_member_array[1] = {
  {
    "data",  // name
    rosidl_typesupport_introspection_c__ROS_TYPE_STRING,  // type
    0,  // upper bound of string
    NULL,  // members of sub message
    false,  // is array
    0,  // array size
    false,  // is upper bound
    offsetof(Generic, data),  // bytes offset in struct
    NULL,  // default value
    NULL,  // size() function pointer
    NULL,  // get_const(index) function pointer
    NULL,  // get(index) function pointer
    NULL  // resize(index) function pointer
  }
};

static const rosidl_typesupport_introspection_c__MessageMembers Generic__rosidl_typesupport_introspection_c__Generic_message_members = {
  "std_msgs__msg",  // message namespace
  "String",  // message name
  1,  // number of fields
  sizeof(Generic),
  Generic__rosidl_typesupport_introspection_c__Generic_message_member_array  // message members
};

// this is not const since it must be initialized on first access
// since C does not allow non-integral compile-time constants
static rosidl_message_type_support_t Generic__rosidl_typesupport_introspection_c__Generic_message_type_support_handle = {
  0,
  &Generic__rosidl_typesupport_introspection_c__Generic_message_members,
  get_message_typesupport_handle_function,
};

// Create generic type
const rosidl_message_type_support_t *
get_generic_type() {
  if (!Generic__rosidl_typesupport_introspection_c__Generic_message_type_support_handle.typesupport_identifier) {
   Generic__rosidl_typesupport_introspection_c__Generic_message_type_support_handle.typesupport_identifier =
     rosidl_typesupport_introspection_c__identifier;
 }
 return &Generic__rosidl_typesupport_introspection_c__Generic_message_type_support_handle;
}

////////////////////////////////
// Functions for Generic type //
////////////////////////////////

bool Generic__init(Generic * msg)
{
  if (!msg) {
    return false;
   }
  // data
  if (!rosidl_generator_c__String__init(&msg->data)) {
    Generic__fini(msg);
    return false;
  }
  return true;
}

void Generic__fini(Generic * msg)
{
  if (!msg) {
    return;
  }
  //data
  rosidl_generator_c__String__fini(&msg->data);
}

Generic * Generic__create()
{
  Generic * msg = (Generic *)malloc(sizeof(Generic));
  if (!msg) {
    return NULL;
  }
  memset(msg, 0, sizeof(Generic));
  bool success = Generic__init(msg);
  if (!success) {
    free(msg);
    return NULL;
  }
  return msg;
}

void Generic__destroy(Generic * msg)
{
  if (msg) {
    Generic__fini(msg);
  }
  free(msg);
}

