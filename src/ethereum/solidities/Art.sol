// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract IndexTable {

  struct ArtLog {
    address owner; // ユーザが所持しているか　
    uint32 nftId;
    bytes[] hashedFile;
    address provider;
    bool isRegistered;
    bytes[] logIds; // 検証前or検証済みのLogをstore　
  }
  mapping(bytes => ArtLog) public ArtLogTable;

  address sp;
  address tpa;

  // LogTableの作成
  struct Log {
    bool result;
    uint32 chal;
    bytes k1;
    bytes k2;
    bytes myu;
    bytes gamma;
    bytes artId;
  }
  mapping(bytes => Log) public Logs;

  struct LogTable {
    Log Log;
    bytes LogId;
  }

  // 公開鍵構造体
  struct PubKey {
      bytes pubkey;
  }

  struct Para {
      string pairing;
      bytes u;
      bytes g;
  }

  //　Para各値の定義
  mapping(address => PubKey) public PubKeys;

  bytes[] ArtIds;
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

    ArtLogTable[_artId].provider = msg.sender;
    ArtLogTable[_artId].hashedFile = _fileData;
    ArtLogTable[_artId].owner = _userAddress;
    ArtIds.push(_artId);
  }

  //artIdを用いてhashedFileを取得
  function GetHashedFile(bytes memory _artId) public view returns(bytes[] memory){
    return ArtLogTable[_artId].hashedFile;
  }

  function SetLogId( uint32 _chal, bytes memory _k1, bytes memory _k2, bytes memory _logId, bytes memory _artId) public {
    //トランザクション送信者がTPAでないと動かない
    require(
      tpa == msg.sender,
      "You are not TPA"
    );

    // LogTableに代入
    Logs[_logId].chal = _chal;
    Logs[_logId].k1 = _k1;
    Logs[_logId].k2 = _k2;
    Logs[_logId].artId = _artId;

    ArtLogTable[_artId].logIds.push(_logId);
  }
  //Logの値からkeyとしてIDを生成，Logsに追加した後IDをreturn
  function registerAuditProof( bytes memory _myu, bytes memory _gamma, bytes memory _artId, bytes memory _logId) public {
    
    //トランザクション送信者がTPAでないと動かない
    require(
      ArtLogTable[_artId].provider == msg.sender,
      "You cannot register Proof"
    );

    // LogTableに証明データを記録
    Logs[_logId].myu = _myu;
    Logs[_logId].gamma = _gamma;
  }

    //Logの値からkeyとしてIDを生成，Logsに追加した後IDをreturn
  function SetAuditResult( bool _result, bytes memory _logId) public {
    
    //トランザクション送信者がTPAでないと動かない
    require(
      tpa == msg.sender,
      "You are not TPA"
    );

    // LogTableに証明データを記録
    Logs[_logId].result = _result;
  }

  // 検証済みのLog情報の処理 => FIT.[_artId].logから該当ログを削除
  function verifyLog(uint256[] memory _logId, uint256 _artId) public {
  }

  // 公開鍵登録
  function RegisterPubKey(bytes memory _pubkey) public {
      PubKeys[msg.sender].pubkey = _pubkey;
  }

  function GetArtIds() public view returns(bytes[] memory){
      return ArtIds;
  }

  function GetLog() public view returns(LogTable[] memory){
    LogTable[] memory LogMemory = new LogTable[](ArtIds.length);
    for(uint i = 0; i < ArtIds.length; i++) {
      for(uint j = 0; j < ArtLogTable[ArtIds[i]].logIds.length; j++) {
        bytes memory logId;
        LogTable memory log;
        logId = ArtLogTable[ArtIds[i]].logIds[j];
        if (Logs[logId].result == false){
          log.LogId = logId;
          log.Log = Logs[logId];
          LogMemory[i] = log;
        }
      }
    }
    return LogMemory;
  }

  //パラメータ取得
  function GetParam() public view returns(Para memory){
      return para;
  }

  //公開鍵取得
  function GetPubKey(bytes memory _artId) public view returns(bytes memory){
    address artOwner = ArtLogTable[_artId].owner;
    return PubKeys[artOwner].pubkey;
  }
}
