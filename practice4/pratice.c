#include <stdio.h>

int sum(int a, int b)
{
  int result = a + b;
  return result;
}

int main()
{
  int a = 5;
  int b = 10;
  int result = sum(a, b);
  printf("The sum of %d and %d is %d\n", a, b, result);
  return 0;
}