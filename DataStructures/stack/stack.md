The stack data structure:
  1. Linear data structure.
  2. The first element in comes out last and vice versa.
  3. Has applications in function calls:
    - If a function(called) function calls other functions, the last function to be called is executed first and the same process goes on till the first function to be called is executed.
  4. Has three major operations:
    1. pop() - Deletes the topmost element.
    2. peek() - Returns the value of the topmost element.
    3. push() - Adds an element to the top of the stack.

    The Implementation of the stack data structure takes two major ways:
    1. Array implementation.
    2. Linked list implementation.

    Array implementation.
    1. pop():
    check if the stack is NULL(i.e if the top element exists)
    if it doesn't: return "empty stack" and exit
    else:
      set val = stack[top]
      set top = top - 1
      end

    2. peek():
    1. Check if stack[max-1] is null
    if null;  return "Stack is empty"
    else:
    return stack[top - 1]
    end

    3. push():
    if stack is full(top = max - 1)
    exit
    else:
    set top = top + 1
    set stack[top] = val
    end/exit

