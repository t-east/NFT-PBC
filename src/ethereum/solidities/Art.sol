// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract IndexTable {

  struct ArtLog {
    address Owner; // ユーザが所持しているか　
    uint32 NftId;
    bytes[] HashedFile;
    address Provider;
    bytes LogId; // 検証前or検証済みのLogをstore　
  }
  mapping(bytes => ArtLog) public ArtLogs;

  // LogTableの作成
  struct AuditLog {
    bool Result;
    uint32 Chal;
    bytes K1;
    bytes K2;
    bytes Myu;
    bytes Gamma;
  }
  mapping(bytes => AuditLog) public AuditLogs;

  struct ArtLogTable {
    bytes Id;
    ArtLog Log;
    bytes PublicKey;
  }

  struct AuditLogTable {
    AuditLog Log;
    bytes Id;
  }

  address Sp;
  address Tpa;
  address User;

  struct Para {
    string Pairing;
    bytes U;
    bytes G;
  }

  //　Para各値の定義
  mapping(address => bytes) public PubKeys;

  bytes[] ArtIds;
  Para para;

  constructor(address _user, address _spa, address _tpa, string memory _pairing, bytes memory _g, bytes memory _u) {
    User = _user;
    Sp = _spa;
    Tpa = _tpa; 
    para.Pairing = _pairing;
    para.G = _g;
    para.U = _u;
  }

  //固有データの登録　=>　fitのIDを返す
  function SetArtLog(bytes[] memory _fileData, bytes memory _artId, address _userAddress) public {
    //トランザクション送信者がSPでないと動かない
    require(Sp == msg.sender, "You are not SP");
    ArtLogs[_artId].Provider = msg.sender;
    ArtLogs[_artId].HashedFile = _fileData;
    ArtLogs[_artId].Owner = _userAddress;
    ArtIds.push(_artId);
  }

  //artIdを用いてArtLog取得
  function GetArtLogs() public view returns(ArtLogTable[] memory){
    ArtLogTable[] memory ArtLogMemory = new ArtLogTable[](ArtIds.length);
    for(uint i = 0; i < ArtIds.length; i++) {
      ArtLogMemory[i].Log = ArtLogs[ArtIds[i]];
      ArtLogMemory[i].Id = ArtIds[i];
      ArtLogMemory[i].PublicKey = PubKeys[ArtLogs[ArtIds[i]].Owner];
    }
    return ArtLogMemory;
  }

   //artIdを用いてArtLog取得
  function GetArtLog(bytes memory _artId) public view returns(ArtLogTable memory){
    ArtLogTable memory artLog;
    artLog.Log = ArtLogs[_artId];
    artLog.Id = _artId;
    artLog.PublicKey = PubKeys[ArtLogs[_artId].Owner];
    return artLog;
  }

  function GetAuditLogs() public view returns(AuditLogTable[] memory){
    AuditLogTable[] memory AuditLogMemory = new AuditLogTable[](ArtIds.length);
    for(uint i = 0; i < ArtIds.length; i++) {
      AuditLogMemory[i].Log = AuditLogs[ArtLogs[ArtIds[i]].LogId];
      AuditLogMemory[i].Id = ArtLogs[ArtIds[i]].LogId;
    }
    return AuditLogMemory;
  }

  function GetAuditLog(bytes memory _auditId ) public view returns(AuditLogTable memory){
    AuditLogTable memory AuditLogMemory;
    AuditLogMemory.Log = AuditLogs[_auditId];
    AuditLogMemory.Id = _auditId;
    return AuditLogMemory;
  }

  function SetAuditChal( uint32 _chal, bytes memory _k1, bytes memory _k2, bytes memory _logId, bytes memory _artId) public {
    //トランザクション送信者がTPAでないと動かない
    require(
      Tpa == msg.sender,
      "You are not TPA"
    );

    // LogTableに代入
    AuditLogs[_logId].Chal = _chal;
    AuditLogs[_logId].K1 = _k1;
    AuditLogs[_logId].K2 = _k2;
    ArtLogs[_artId].LogId = _logId;
  }

  //Logの値からkeyとしてIDを生成，Logsに追加した後IDをreturn
  function SetAuditProof( bytes memory _myu, bytes memory _gamma, bytes memory _artId) public {
    
    //トランザクション送信者がTPAでないと動かない
    require(
      ArtLogs[_artId].Provider == msg.sender,
      "You cannot register Proof"
    );

    // LogTableに証明データを記録
    AuditLogs[ArtLogs[_artId].LogId].Myu = _myu;
    AuditLogs[ArtLogs[_artId].LogId].Gamma = _gamma;
  }

    //Logの値からkeyとしてIDを生成，Logsに追加した後IDをreturn
  function SetAuditResult( bool _result, bytes memory _logId) public {
    
    //トランザクション送信者がTPAでないと動かない
    require(
      Tpa == msg.sender,
      "You are not TPA"
    );

    // LogTableに証明データを記録
    AuditLogs[_logId].Result = _result;
  }

  // 公開鍵登録
  function RegisterPubKey(bytes memory _pubkey) public {
    require(
      User == msg.sender,
      "You are not User"
    );
      PubKeys[msg.sender] = _pubkey;
  }

  //パラメータ取得
  function GetPubkey(address _address) public view returns(bytes memory){
      return PubKeys[_address];
  }

  //パラメータ取得
  function GetParam() public view returns(Para memory){
      return para;
  }
}
