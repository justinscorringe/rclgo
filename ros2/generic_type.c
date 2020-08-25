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

const rosidl_typesupport_introspection_c__MessageMember * new_generic_members(GoMember go_members_[], uint32_t member_count_) {
  rosidl_typesupport_introspection_c__MessageMember* new_member_array = malloc(member_count_ * sizeof *new_member_array);
  
  // Create full member spec above from go members
  for (uint32_t i = 0; i < member_count_; i++) {

    GoMember goMember = go_members_[i];

     // Start with primitive non-arrays
    if ((goMember.is_array_ == false) && (goMember.type_id_ != rosidl_typesupport_introspection_c__ROS_TYPE_MESSAGE)) {
      new_member_array[i].name_ = goMember.name_; // name
      new_member_array[i].type_id_ =  goMember.type_id_;  // type
      new_member_array[i].string_upper_bound_ = 0;  // upper bound of string
      new_member_array[i].members_ = NULL;  // members of sub message
      new_member_array[i].is_array_ = false;  // is array
      new_member_array[i].array_size_ = 0;  // array size
      new_member_array[i].is_upper_bound_ = false;  // is upper bound
      new_member_array[i].offset_ = goMember.member_offset_;  // bytes offset in struct
      new_member_array[i].default_value_ = NULL;  // default value
      // new_member_array[i].(* size_function)(const void *) = NULL;  // size() function pointer
      // new_member_array[i].(*get_const_function)(const void *, size_t index) = NULL;  // get_const(index) function pointer
      // new_member_array[i].(*get_function)(void *, size_t index) = NULL;  // get(index) function pointer
      // new_member_array[i].(* resize_function)(void *, size_t size) = NULL;  // resize(index) function pointer
       }
    }
  // Create size, get_const, get, and resize  function interfaces also.

  return new_member_array;
}

const GoMembers new_generic_struct(GoMember go_members_[], size_t member_count_) {
  // Generate size of struct and offset values with padding

  int32_t largest_member_ = 0;
  // Find size of largest member (built in non array :( )
  for (uint32_t i = 0; i < member_count_; i++) {

    GoMember member = go_members_[i];
    if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_INT8) {
      if (sizeof(int8_t) > largest_member_) {
        largest_member_ = sizeof(int8_t);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_INT16) {
      if (sizeof(int16_t) > largest_member_) {
        largest_member_ = sizeof(int16_t);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_INT32) {
      if (sizeof(int32_t) > largest_member_) {
        largest_member_ = sizeof(int32_t);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_INT64) {
      if (sizeof(int64_t) > largest_member_) {
        largest_member_ = sizeof(int64_t);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_UINT8) {
      if (sizeof(uint8_t) > largest_member_) {
        largest_member_ = sizeof(uint8_t);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_UINT16) {
      if (sizeof(uint16_t) > largest_member_) {
        largest_member_ = sizeof(uint16_t);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_UINT32) {
      if (sizeof(uint32_t) > largest_member_) {
        largest_member_ = sizeof(uint32_t);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_UINT64) {
      if (sizeof(uint64_t) > largest_member_) {
        largest_member_ = sizeof(uint64_t);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_FLOAT32) {
      if (sizeof(float) > largest_member_) {
        largest_member_ = sizeof(float);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_FLOAT64) {
      if (sizeof(float) > largest_member_) {
        largest_member_ = sizeof(float);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_STRING) {
      if (sizeof(rosidl_generator_c__String) > largest_member_) {
        largest_member_ = sizeof(rosidl_generator_c__String);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_BOOL) {
      if (sizeof(bool) > largest_member_) {
        largest_member_ = sizeof(bool);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_CHAR) {
      if (sizeof(char) > largest_member_) {
        largest_member_ = sizeof(char);
      }
    } else if (member.type_id_ == rosidl_typesupport_introspection_c__ROS_TYPE_BYTE) {
      if (sizeof(char) > largest_member_) {
        largest_member_ = sizeof(char);
      }
    }
    

  }

  // Calculate offset of each value
  for (uint32_t i = 0; i < member_count_; i++) {

    GoMember member = go_members_[i];

    member.member_offset_ = (i * largest_member_);
  }

  // Calculate size 
  size_t struct_size_ = largest_member_ * member_count_;

  GoMembers members;
  
  members.member_array = go_members_;
  members.struct_size_ = struct_size_;

  return members;

}

static const rosidl_typesupport_introspection_c__MessageMembers Generic__rosidl_typesupport_introspection_c__Generic_message_members = {
  "std_msgs__msg",  // message namespace
  "String",  // message name
  1,  // number of fields
  sizeof(Generic),
  Generic__rosidl_typesupport_introspection_c__Generic_message_member_array  // message members
};

// Create a generic type support
const rosidl_message_type_support_t new_generic_type(
  const char* message_namespace_, 
  const char* message_name_, 
  uint32_t member_count_,
  GoMember go_members_[])
  {

  // Create struct size and offsets   
  const GoMembers members_ = new_generic_struct(go_members_, member_count_);

  // Create introspection message members from go members
  const rosidl_typesupport_introspection_c__MessageMember * member_array_ = new_generic_members(members_.member_array, member_count_);
 
  // Create introspection members object
    const rosidl_typesupport_introspection_c__MessageMembers members = {
  message_namespace_,
  message_name_,
  member_count_,
  members_.struct_size_,
  member_array_
  //void (* init_function)(void *, enum rosidl_runtime_c_message_initialization);
  //void (* fini_function)(void *);
  };

  const rosidl_message_type_support_t generic_type = {
  rosidl_typesupport_introspection_c__identifier,
  &members,
  get_message_typesupport_handle_function
  };

  return generic_type;
}

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

// Functions for Generic type

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
  // data
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

