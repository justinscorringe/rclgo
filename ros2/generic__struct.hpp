#ifndef GENERIC__STRUCT_HPP_
#define GENERIC__STRUCT_HPP_

#include <rosidl_generator_cpp/bounded_vector.hpp>
#include <rosidl_generator_cpp/message_initialization.hpp>
#include <algorithm>
#include <array>
#include <memory>
#include <string>
#include <vector>


// message struct
template<class ContainerAllocator>
struct Generic_
{
  using Type = Generic_<ContainerAllocator>;

  explicit Generic_(rosidl_generator_cpp::MessageInitialization _init = rosidl_generator_cpp::MessageInitialization::ALL)
  {
    if (rosidl_generator_cpp::MessageInitialization::ALL == _init ||
      rosidl_generator_cpp::MessageInitialization::ZERO == _init)
    {
      this->data = "";
    }
  }

  explicit Generic_(const ContainerAllocator & _alloc, rosidl_generator_cpp::MessageInitialization _init = rosidl_generator_cpp::MessageInitialization::ALL)
  : data(_alloc)
  {
    if (rosidl_generator_cpp::MessageInitialization::ALL == _init ||
      rosidl_generator_cpp::MessageInitialization::ZERO == _init)
    {
      this->data = "";
    }
  }

  // field types and members
  using _data_type =
    std::basic_string<char, std::char_traits<char>, typename ContainerAllocator::template rebind<char>::other>;
  _data_type data;

  // setters for named parameter idiom
  Type & set__data(
    const std::basic_string<char, std::char_traits<char>, typename ContainerAllocator::template rebind<char>::other> & _arg)
  {
    this->data = _arg;
    return *this;
  }

  // constant declarations

  // pointer types
  using RawPtr =
    Generic_<ContainerAllocator> *;
  using ConstRawPtr =
    const Generic_<ContainerAllocator> *;
  using SharedPtr =
    std::shared_ptr<Generic_<ContainerAllocator>>;
  using ConstSharedPtr =
    std::shared_ptr<Generic_<ContainerAllocator> const>;

  template<typename Deleter = std::default_delete<
      Generic_<ContainerAllocator>>>
  using UniquePtrWithDeleter =
    std::unique_ptr<Generic_<ContainerAllocator>, Deleter>;

  using UniquePtr = UniquePtrWithDeleter<>;

  template<typename Deleter = std::default_delete<
      Generic_<ContainerAllocator>>>
  using ConstUniquePtrWithDeleter =
    std::unique_ptr<Generic_<ContainerAllocator> const, Deleter>;
  using ConstUniquePtr = ConstUniquePtrWithDeleter<>;

  using WeakPtr =
    std::weak_ptr<Generic_<ContainerAllocator>>;
  using ConstWeakPtr =
    std::weak_ptr<Generic_<ContainerAllocator> const>;


  // comparison operators
  bool operator==(const Generic_ & other) const
  {
    if (this->data != other.data) {
      return false;
    }
    return true;
  }
  bool operator!=(const Generic_ & other) const
  {
    return !this->operator==(other);
  }
};  // struct String_

// alias to use template instance with default allocator
using Generic =
  Generic_<std::allocator<void>>;


#endif  // GENERIC__STRUCT_HPP_