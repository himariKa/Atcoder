#include <iostream>
#include <string>
using namespace std;

int main(){
    // 入力データの個数を表す変数Nへの入力
    int A,B,C,X;
    cin >> A >> B >> C >> X;

    int counter = 0;
    // 500円何枚か
    for (A = 0; A <= 50; A++)
    {
        // 100円が何枚か
       for (B = 0; B <= 50; B++)
       {
            // 50円が何枚か
            for (C = 0; C <= 50; C++)
            {
                if(X==500*A+100*B+50*C){
                    counter++;
                    return;
                }

            }
       }
       
    }
    cout << counter << endl;
    
}