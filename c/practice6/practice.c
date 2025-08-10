/* 조건문 연습 */
#include <stdio.h>

int main()
{
  int score = 92;
  char grade;
  if (score >= 90)
  {
    grade = 'A';
  }
  else
  {
    grade = 'F';
  }

  printf("점수는 %d이고, 학점은 %c입니다.\n", score, grade);
  return 0;
}