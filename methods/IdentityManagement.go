package methods

import (
	"context"
	"log"
	"time"

	"hestia/api/pb"
	"hestia/api/utils/db"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IdentityManagementServer struct {
	pb.UnimplementedIdentityManagementServiceServer
}

func (s *IdentityManagementServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing email")
	}
	if in.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing password")
	}

	//Check if user exists in the database
	db, err := db.GetDbPoolConn()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	var userId uuid.UUID
	var name string
	err = db.QueryRow(ctx, "SELECT id, name FROM users.users WHERE email = $1 AND password = $2", in.GetEmail(), in.GetPassword()).Scan(&userId, &name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Error(codes.NotFound, "Wrong email or password")
		} else {
			log.Println(err)
			return nil, status.Error(codes.Internal, "Database error")
		}
	}

	// Generate a token
	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}
	defer tx.Rollback(ctx)

	// Insert token into the database
	token := uuid.New()
	_, err = tx.Exec(ctx, "INSERT INTO users.users_session (id, user_id, expiry_date) VALUES ($1, $2, $3)", token, userId, time.Now().Add(time.Hour*72))
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	return &pb.LoginResponse{Token: token.String(), Name: name, Email: in.GetEmail(), Companies: []*pb.Company{}}, nil
}

func (s *IdentityManagementServer) Alive(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	// Get metadata X-AUTH-TOKEN
	metadata, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	token := metadata.Get("X-AUTH-TOKEN")
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	db, err := db.GetDbPoolConn()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Check if token exists in the database
	var expiry_date time.Time
	err = db.QueryRow(ctx, "SELECT expiry_date FROM users.users_session WHERE id = $1", token[0]).Scan(&expiry_date)
	if err != nil {
		if err == pgx.ErrNoRows {
			// For security reasons, we don't want to give the user any information about the token
			return nil, status.Error(codes.Unauthenticated, "Token expired")
		} else {
			// Handle other errors
			log.Println(err)
			return nil, status.Error(codes.Internal, "Database error")
		}
	}

	// Check if token is expired
	if expiry_date.Before(time.Now()) {
		return nil, status.Error(codes.Unauthenticated, "Token expired")
	}

	return &emptypb.Empty{}, nil
}
