package models

type Relacion struct {
	UsuarioID         string `bson:"usuario_id" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuario_relacion_id" json:"usuarioRelacionId"`
}
