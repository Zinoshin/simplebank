package token

import "time"

// Maker는 토큰을 관리하기 위한 인터페이스
type Maker interface {
	// CreateToken은 특정한 사용자 이름과 기간을 위한 새로운 토큰을 생성
	CreateToken(username string, duration time.Duration) (string, error)

	// VertifyToken은 토큰이 유효한지 확인
	VerifyToken(token string) (*Payload, error)
}
