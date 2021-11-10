// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract IndexTable {

  struct ArtLog {
    mapping(address => bool) owner; // ユーザが所持しているか　
    bytes[] hashedFile;
    bool isRegistered;
    bytes[] logIds; // 検証前or検証済みのLogをstore　
  }
  mapping(bytes => ArtLog) public ArtLogTable;

  address sp;
  address tpa;

  // LogTableの作成
  struct Log {
    bool result;
    uint chal;
    bytes k1;
    bytes k2;
    bytes myu;
    bytes gamma;
  }
  mapping(string => Log) public Logs;

  // 公開鍵構造体
  struct PubKey {
      string pubkey;
  }

  struct Para {
      string pairing;
      bytes u;
      bytes g;
  }

  //　Para各値の定義
  mapping(address => PubKey) public PubKeys;

  Para para;

  constructor(address _spa, address _tpa, string memory _pairing, bytes memory _g, bytes memory _u) {
    sp = _spa;
    tpa = _tpa; 
    para.pairing = _pairing;
    para.g = _g;
    para.u = _u;
  }

  //固有データの登録　=>　fitのIDを返す
  function registerArt(bytes[] memory _fileData, bytes memory _artId, address _userAddress) public {
    //トランザクション送信者がSPでないと動かない
    require(sp == msg.sender, "You are not SP");

    // 作品がすでに登録されているか
    // require(
    //   ArtLogTable[_artId].owner[_userAddress] == false,
    //   "This user has already registered as the owner"
    // );
    // すでに登録済みのデータであれば不適切
    require(
      ArtLogTable[_artId].isRegistered == false,
      "This user has be already registered as the owner"
    );

    ArtLogTable[_artId].hashedFile = _fileData;
    ArtLogTable[_artId].owner[_userAddress] = true;
  }

  //artIdを用いてhashedFileを取得
  function getHashedFile(bytes memory _artId) public view returns(bytes[] memory){
    return ArtLogTable[_artId].hashedFile;
  }

  //Logの値からkeyとしてIDを生成，Logsに追加した後IDをreturn
  function registerLog( bool _result, uint _chal, bytes memory _k1, bytes memory _k2, bytes memory _myu, bytes memory _gamma, bytes memory _artId, string memory _logId) public {
    
    //トランザクション送信者がTPAでないと動かない
    require(
      tpa == msg.sender,
      "You are not TPA"
    );

    // LogTableに代入
    Logs[_logId].result = _result;
    Logs[_logId].chal = _chal;
    Logs[_logId].k1 = _k1;
    Logs[_logId].k2 = _k2;
    Logs[_logId].myu = _myu;
    Logs[_logId].gamma = _gamma;

    ArtLogTable[_artId].logIds.push(_artId);
  }

  //ID配列を入力
  //該当するLogを配列にしてreturn
  function getLog(string[] memory _logId) public view returns(Log[] memory){
    Log[] memory LogMemory = new Log[](_logId.length);
    for(uint i = 0; i < _logId.length; i++) {
        LogMemory[i] = Logs[_logId[i]];
    }
    return LogMemory;
  }

  // 検証済みのLog情報の処理 => FIT.[_artId].logから該当ログを削除
  function verifyLog(uint256[] memory _logId, uint256 _artId) public {
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
