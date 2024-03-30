#include <stdio.h>
#include <stdlib.h>
#include <string.h>

//Global variables
int stack[256];//array initialisation
int count = 0;//counter initialisation

void push(int x)
{
  //adds an element to topmost index of the stack
  if (count == 256)
  {
    fprintf(stderr, "Stack Overflow\n");
    exit(1);
  }
  stack[count] = x;
  count++;
}

int pop()
{
  if (count == 0)
  {
    fprintf(stderr, "Empty stack\n");
    return(-1);
  }
  int res = stack[count - 1];
  count--;
  return res;
}

int main(char *argv[], int argc)
{
    push(24);
    push(36);
    push(27);
    push(34);
    push(9);

    int i = 0;
    for (; i < 5; i++)
      printf("%d\n", pop());

    return EXIT_SUCCESS;
}
