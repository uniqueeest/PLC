#include <stdio.h>
#include <string.h>

int GetStringLength(char data[])
{
  int i = 0;
  while (data[i])
  {
    i++;
  }
  return i;
}

int main()
{
  /* 직접 구현 함수 사용 */
  int data_length;
  char data[] = "Happy";
  data_length = GetStringLength(data);
  printf("문자열의 길이는 %d입니다. 함수를 사용했습니다.\n", data_length);

  /* strlen 메서드 사용 */
  int data_length2;
  char data2[] = "Happy";
  data_length2 = strlen(data2);
  printf("문자열의 길이는 %d입니다. 메서드를 사용했습니다.\n", data_length2);

  /* strcopy, strcat 사용 */
  char data3[] = "Happy";
  char data4[20];

  strcpy(data4, data3);
  strcat(data4, " Birthday");
  printf("%s + \"Birthday\" = %s.\n", data3, data4);
}