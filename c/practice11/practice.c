#include <stdio.h>

int main()
{
  char data[5] = {1, 2, 3, 4, 5};
  int result = 0;
  char *ptr = data;
  for (int i = 0; i < 5; i++)
  {
    result += *ptr;
    ptr++;
  }

  printf("The sum of the array elements is: %d\n", result);
}