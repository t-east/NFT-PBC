// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;
contract Params{

    // 公開鍵構造体
    struct PubKey {
        string pubkey;
    }

    struct Para {
        string pairing;
        string u;
        string g;
    }

    //　Para各値の定義
    mapping(address => PubKey) public PubKeys;
    Para para;

    address paraRegisterOwner;

    // コントラクトデプロイ時に，アドレスを登録
    constructor(address _systemManager){
        paraRegisterOwner = _systemManager;
    }

    // Para登録
    function RegisterParams(string memory _pairing,string memory _g,string memory _u) public {
        require(paraRegisterOwner == msg.sender, "You are not a system manager");
        para.pairing = _pairing;
        para.g = _g;
        para.u = _u;
        // etc...
    }

    // 公開鍵登録
    function RegisterPubKey(string memory _pubkey) public {
        PubKeys[msg.sender].pubkey = _pubkey;
    }

    //パラメータ取得
    function GetParam() public view returns(Para memory){
        return para;
    }

    //公開鍵取得
    function GetPubKey(address _owner) public view returns(string memory){
        return PubKeys[_owner].pubkey;
    }


}