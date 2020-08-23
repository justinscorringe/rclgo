#include <stddef.h>
#include "rosidl_typesupport_introspection_c/field_types.h"
#include "rosidl_typesupport_introspection_c/identifier.h"
#include "rosidl_typesupport_introspection_c/message_introspection.h"
#include "generic_type.h"
#include "generic__struct.h"

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
get_generic_type() 
{
  if (!Generic__rosidl_typesupport_introspection_c__Generic_message_type_support_handle.typesupport_identifier) {
   Generic__rosidl_typesupport_introspection_c__Generic_message_type_support_handle.typesupport_identifier =
     rosidl_typesupport_introspection_c__identifier;
 }return &Generic__rosidl_typesupport_introspection_c__Generic_message_type_support_handle;
}
