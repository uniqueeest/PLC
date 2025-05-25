#include <stdio.h>

void test(short *ptr)
{
  short soft = 0;
  soft = *ptr;
  printf("soft =  %d", soft);
  *ptr = 100;
}

void swap(short *a, short *b)
{
  short temp;
  temp = *a;
  *a = *b;
  *b = temp;
}

int main()
{
  /* 1번 */
  short birthday;
  short *ptr;
  ptr = &birthday;

  printf("birthday 변수의 주소는 %p입니다 \n", ptr);

  /* 2번 */
  short tip = 3;
  test(&tip);
  printf("\ntip 변수의 값은 %d입니다 \n", tip);

  /* 3번 */
  short a = 10, b = 20;
  printf("a = %d, b = %d\n", a, b);
  swap(&a, &b);
  printf("a = %d, b = %d\n", a, b);

  /* 4번 */
  int data = 0x12345678, i;
  char *p = (char *)&data;
  for (i = 0; i < sizeof(data); i++)
  {
    printf("%X, ", *(p + i));
  }

  return 0;
}