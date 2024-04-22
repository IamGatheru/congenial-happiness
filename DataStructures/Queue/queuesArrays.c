#include <stdlib.h>
#include <stdio.h>
#include <string.h>

int queue[256];
int count = 0;

int removeFromQueue()
{
  int i;
  for (i = 0; i < count - 1; i++)
    queue[i] = queue[i + 1];
  count--;
  //return queue;
}

void addToQueue(int x)
{
  queue[count] = x;
  count++;
}

int main(char *argv[], int argc)
{
  addToQueue(90);
  addToQueue(45);
  addToQueue(88);
  addToQueue(87);
  addToQueue(97);

  int j = 0;
  for (; j < count - 1; j++)
    printf("%d\n", queue[j]);

  return (EXIT_SUCCESS);
}
