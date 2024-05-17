package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sachatarba/course-db/internal/entity"
)

type AuthorizationService struct {
	// Repo
	sessionRepo ISessionRepository

	// Services
	clientService IClientService
}

func NewAuthorizationService(sessionRepo ISessionRepository, clientService IClientService) IAuthorizationService {
	return &AuthorizationService{
		sessionRepo:   sessionRepo,
		clientService: clientService,
	}
}

func (a *AuthorizationService) IsAuthorize(ctx context.Context, sessionID uuid.UUID) (bool, error) {
	_, err := a.sessionRepo.GetSessionBySessionID(ctx, sessionID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (a *AuthorizationService) Authorize(ctx context.Context, login string, password string) (entity.Session, error) {
	client, err := a.clientService.GetClientByLogin(ctx, login)
	if err != nil {
		return entity.Session{}, err
	}

	// "кеш функцию знаешь? я закешировал, понятно, да?"
	if client.Password != password {
		return entity.Session{}, ErrWrongPassword
	}

	session := entity.Session{
		ClientID:  client.ID,
		SessionID: uuid.New(),
		TTL:       time.Now().Add(10 * time.Hour),
	}

	err = a.sessionRepo.CreateNewSession(ctx, session)
	if err != nil {
		return entity.Session{}, err
	}

	return session, nil
}

func (a *AuthorizationService) Register(ctx context.Context, client entity.Client) (entity.Session, error) {
	err := a.clientService.RegisterNewClient(ctx, client)
	if err != nil {
		return entity.Session{}, err
	}

	session := entity.Session{
		ClientID:  client.ID,
		SessionID: uuid.New(),
		TTL:       time.Now().Add(10 * time.Hour),
	}

	err = a.sessionRepo.CreateNewSession(ctx, session)
		if err != nil {
		return entity.Session{}, err
	}

	return session, nil
}

func (a *AuthorizationService) Logout(ctx context.Context, sessionID uuid.UUID) (entity.Session, error) {
	session, err := a.sessionRepo.GetSessionBySessionID(ctx, sessionID)
	if err != nil {
		return entity.Session{}, err
	}

	err = a.sessionRepo.DeleteSession(ctx, sessionID)
	if err != nil {
		return entity.Session{}, err
	}

	session = entity.Session{
		ClientID:  sessionID,
		SessionID: session.SessionID,
		TTL:       time.Now().Add(-10 * time.Hour),
	}

	return session, nil
}


// Пока особой логики нет, надо подумать над удалением сессий (надо ли оно)
func (a *AuthorizationService) DeleteClient(ctx context.Context, clientID uuid.UUID) (entity.Session, error) {
	err := a.clientService.DeleteClient(ctx, clientID)
	if err != nil {
		return entity.Session{}, err
	}
	// В принципе сессии можно не удалять, 
	// так как в редисе они сами удалятся по истечению TTL,
	// а удалять все сессии может быть достаточно долго

	// sessions, err := a.sessionRepo.GetSessionsByClientID(ctx, clientID)
	// if err != nil {
	// 	return entity.Session{}, err
	// }

	// for i := 0; i < len(sessions) && err == nil; i++ {
	// 	err = a.sessionRepo.DeleteSession(ctx, sessions[i].SessionID)
	// }
	// if err != nil {
	// 	return entity.Session{}, err
	// }

	// session := entity.Session{
	// 	ClientID:  clientID,
	// 	SessionID: sessionID,
	// 	TTL:       time.Now().Add(-10 * time.Hour),
	// }

	session := entity.Session{}

	return session, nil
}
