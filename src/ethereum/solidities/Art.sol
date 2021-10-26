// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract IndexTable {

  struct ArtLog {
    mapping(address => bool) owner; // ユーザが所持しているか　
    bytes[] hashedFile;
    bool isRegistered;
    uint[] logIds; // 検証前or検証済みのLogをstore　
  }
  mapping(uint256 => ArtLog) public ArtLogTable;

  address sp;
  address tpa;

  // LogTableの作成
  struct Log {
    bool result;
    string chal;
    string k1;
    string k2;
    string myu;
    string gamma;
  }
  mapping(uint256 => Log) public Logs;

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
  function registerOriginalData(bytes[] memory _fileData, address _userAddress) public returns(uint256){
    //トランザクション送信者がSPでないと動かない
    require(sp == msg.sender, "You are not SP");

    // ハッシュファイルからfitIDを生成
    uint256 fitId = uint256(keccak256(_fileData[0]));

    // 登録ユーザがすでに登録されているか
    require(
      ArtLogTable[fitId].owner[_userAddress] == false,
      "This user has already registered as the owner"
    );
    // すでに登録済みのデータであれば不適切
    require(
      ArtLogTable[fitId].isRegistered == false,
      "This user has be already registered as the owner"
    );

    ArtLogTable[fitId].hashedFile = _fileData;
    ArtLogTable[fitId].owner[_userAddress] = true;

    // ArtLogTable[fitId].firstUser = msg.sender;
    // ↑固有ユーザの公開鍵を取得できるようにしたい→paraコントラクトからid検索で取得？

    return fitId;
  }
  
  //重複データの登録
  function registerDedUpData(uint256 _fitId, address _userAddress) public {
    //トランザクション送信者がSPでないと動かない
    require(sp == msg.sender, "You are not SP");

    //既に登録されているユーザがいなかったら固有データなのでだめ．
    require(
      ArtLogTable[_fitId].isRegistered,
      "This data has not be registered"
    );
    ArtLogTable[_fitId].owner[_userAddress] = true;
  }

  // dev用 FITのownerを削除
  function deleteData(bytes[] memory _fileData, address _userAddress) public {

    // ハッシュファイルからfitIDを生成
    uint256 fitId = uint256(keccak256(_fileData[0]));

    // 登録ユーザでは無くす
    ArtLogTable[fitId].owner[_userAddress] = false;

    // 登録されていないデータとして扱う
    ArtLogTable[fitId].isRegistered = false;

    ArtLogTable[fitId].hashedFile = _fileData;
  }

  //IndexTableから指定したIDのものを取得
  function getHashedFile(uint _fitId) public view returns(bytes[] memory){
    return ArtLogTable[_fitId].hashedFile;
  }

  //Logの値からkeyとしてIDを生成，Logsに追加した後IDをreturn
  function registerLog( bool _result, string memory _chal, string memory _k1, string memory _k2, string memory _myu, string memory _gamma, uint256 _fitId) public returns(uint256){
    
    //トランザクション送信者がTPAでないと動かない
    require(
      tpa == msg.sender,
      "You are not TPA"
    );

    // logIdを入力データから生成
    uint256 logId = uint256(keccak256(abi.encodePacked(_result, _chal, _k1, _k2, _myu, _gamma, _fitId)));

    // LogTableに代入
    Logs[logId].result = _result;
    Logs[logId].chal = _chal;
    Logs[logId].k1 = _k1;
    Logs[logId].k2 = _k2;
    Logs[logId].myu = _myu;
    Logs[logId].gamma = _gamma;

    //TODO logIdをArtLogTable[_fitId].logIdsにlogId追加するコントラクトをFIT側で定義→ここで実行

    return logId;
  }

  //ID配列を入力
  //該当するLogを配列にしてreturn
  function getLog(uint256[] memory _logId) public view returns(Log[] memory){
    Log[] memory LogMemory = new Log[](_logId.length);
    for(uint i = 0; i < _logId.length; i++) {
        LogMemory[i] = Logs[_logId[i]];
    }
    return LogMemory;
  }

  // 検証済みのLog情報の処理 => FIT.[_fitId].logから該当ログを削除
  function verifyLog(uint256[] memory _logId, uint256 _fitId) public {
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
