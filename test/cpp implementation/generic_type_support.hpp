#ifndef GENERIC_TYPE_SUPPORT_CPP_HPP_
#define GENERIC_TYPE_SUPPORT_CPP_HPP_

#include "rosidl_generator_c/message_type_support_struct.h"
#include "rosidl_typesupport_interface/macros.h"
#include "rosidl_typesupport_introspection_cpp/visibility_control.h"
#include "rcutils/allocator.h"
#include "generic__struct.hpp"


#ifdef __cplusplus
extern "C"
{
#endif

ROSIDL_TYPESUPPORT_INTROSPECTION_CPP_PUBLIC
const rosidl_message_type_support_t * get_generic_type();



#ifdef __cplusplus
}
#endif


extern "C" {
    void Generic();
}



#endif // GENERIC_TYPE_SUPPORT_CPP_HPP_