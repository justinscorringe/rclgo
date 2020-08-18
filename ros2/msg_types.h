#ifndef MGS_TYPES_H
#define MSG_TYPES_H

#include <rcl/rcl.h>
#include <rosidl_generator_c/message_type_support_struct.h>
#include <rosidl_generator_c/string_functions.h>
#include <rcl/error_handling.h>

// MACROS for function headers
#define GET_MSG_TYPE_SUPPORT_HEADER(x,y,z) const rosidl_message_type_support_t* get_message_type_from_## x ##_## y ##_## z ();
//#define CREATE_MSG_INIT_HEADER(x,y,z) const x##__##y##__##z* init_## x ##_## y ##_## z ();
//#define CREATE_MSG_DESTROY_HEADER(x,y,z) void destroy_## x ##_## y ##_## z (x##__##y##__##z* msg);


// MACROS for function body
#define GET_MSG_TYPE_SUPPORT(x,y,z) const rosidl_message_type_support_t* get_message_type_from_## x ##_## y ##_## z (){ \
  return ROSIDL_GET_MSG_TYPE_SUPPORT(x,y,z); \
}

// #define CREATE_MSG_INIT(x,y,z) const x##__##y##__##z* init_## x ##_## y ##_## z (){ \
//   return  x##__##y##__##z##__create(); \
// }

// #define CREATE_MSG_DESTROY(x,y,z) void destroy_## x ##_## y ##_## z (x##__##y##__##z* msg){ \
//   return  x##__##y##__##z##__destroy(msg); \
//   }



#include <std_msgs/msg/string.h>
GET_MSG_TYPE_SUPPORT_HEADER(std_msgs,msg,String)
// CREATE_MSG_INIT_HEADER(std_msgs,msg,String)
// CREATE_MSG_DESTROY_HEADER(std_msgs,msg,String)

#endif