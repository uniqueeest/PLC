## 함수

### main 함수

- c언어로 작성한 프로그램은 함수의 집합체고 그 함수들끼리 서로 호출하는 형태로 진행
- 컴파일러는 어떤 함수가 최초로 시작되는 함수인지 알아야 번역 가능
- c 언어는 main 이라는 함수를 '프로그램 시작 함수'로 약속함

### 함수의 원형

- 두 함수가 서로 호출하는 경우 피호출자가 항상 호출자의 위쪽에 놓을 수 없기 때문에 함수 원형(Function Prototype)을 사용해서 해결
- 함수의 원형은 함수 이름, 매개변수, 반환 자료형을 포함하는 표현식

```c
int Sum (int value1, int value2); /* 함수의 원형 선언 */

void main()
{
  int s = Sum(a, b)
}

int Sum(int value1, int value2)
{
  int result = value1 + value2;
  return result;
}
```
