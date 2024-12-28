#include <iostream>
using namespace std;

int main(){
    // KとAとBの型を定義
    int K,A,B;
    // KとAとBの標準入力を受け付ける
    cin >> K >> A >> B;

    // デフォルトは存在していない
    bool exist = false;
    // A以上B以下を一つずつ調べる
    for (int i=A;i<=B;i++){
        // 存在していた場合の処理
        if(i % K == 0){
            exist = true;
        }
    }

    if (exist){
        cout << "OK" << endl;
    }else{
        cout << "NG" << endl;
    }
}