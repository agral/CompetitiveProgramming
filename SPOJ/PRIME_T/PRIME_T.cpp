#include <iostream>
#include <cmath>

using namespace std;

bool isPrime(long n)
{
  int threshold = sqrt(n);
  for(int i=2;i<=threshold;i++)
  {
    if(n%i==0)
    {
      return false;
    }
  }
  return true;
}

int main()
{
  long num;
  long n;
  cin>>n;
  for(long i=0;i<n;i++)
  {
    cin>>num;
    if(num<2)
    {
      cout<<"NIE"<<endl;
    }
    else
    {
      if(isPrime(num))
      {
        cout<<"TAK"<<endl;
      }
      else
      {
        cout<<"NIE"<<endl;
      };
    };
  }
  return 0;
}
