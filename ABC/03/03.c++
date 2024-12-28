#include <iostream>
using namespace std;

int main(){
    // HとAの型を定義
    int H,A,hit;
    // AとBの標準入力を受け付ける
    cin >> H >> A;
    
    if (H%A==0){
        // ちょうど割り切れる場合
        cout << H / A << endl;
    }else{
        //ちょうど割り切れない場合
        cout << H / A + 1 << endl;
    }
}