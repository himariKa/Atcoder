#include <iostream>
#include <string>
using namespace std;

//　判定関数
bool isAC(string S)
    {
    // Sの最初の文字列が大文字のAではない
    if(S[0] != 'A'){
        return false;
    }

    // Sの先頭2文字と末尾1文字を除いた'C'の個数を調べる
    int num_c = 0;
    for (int i = 2; i + 1 < S.size(); ++i){
        if (S[i] == 'C'){
            ++num_c;
        }
    }
    if(num_c != 1){
        return false;
    }

    //Sに含まれる大文字の個数を調べる
    int num_big = 0;
    for (auto c : S){
        if (isupper(c)){
            ++num_big;
        }
    }
    if(num_big != 2){
        return false;
    }

    //　条件が全て満たされている場合はtrue
    return true;
}

int main(){
    // 入力データの個数を表す変数Nへの入力
    string S;
    cin >> S;

    // 条件を判定して出力
    if (isAC(S)){
        cout << "AC" << endl;
    }else{
        cout << "WA" << endl;
    }
    
}