package handler

// ethereum接続の話と，pbcのinterfaceを作る

type AuditLog interface {
    GetArtLogs() ([][]byte, err)
    Challen() (int, *pbc.Element, *pbc.Element, error)
    AuditVerify() (bool, error)
}

type ArtLog interface {
    GetArtLogs() ([][]byte, err)
}