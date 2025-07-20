package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/saeedjhn/go-otp-auth/internal/models"

	"github.com/saeedjhn/go-otp-auth/internal/types"

	mysqlrepo "github.com/saeedjhn/go-otp-auth/internal/repository/mysql"

	"github.com/saeedjhn/go-otp-auth/pkg/persistance/db/mysql"
	"github.com/saeedjhn/go-otp-auth/pkg/richerror"

	"github.com/saeedjhn/go-otp-auth/pkg/msg"
)

type DB struct {
	conn *mysql.DB
}

func New(conn *mysql.DB) DB {
	return DB{conn: conn}
}

func (r DB) Create(ctx context.Context, u models.User) (models.User, error) {
	query := "INSERT INTO users (name, mobile, email, password) values(?, ?, ?, ?)"

	stmt, err := r.conn.PrepareStatement( //nolint:sqlclosecheck // nothing
		ctx, uint(mysqlrepo.StatementKeyUserCreate), query,
	)
	if err != nil {
		return models.User{}, richerror.New(_opMysqlUserCreate).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	res, err := stmt.ExecContext(ctx, u.Name, u.Mobile, u.Email, u.Password)
	// res, err := r.conn.Conn().Exec(query, u.Name, u.Mobile, u.Email, u.Password)
	if err != nil {
		return models.User{},
			richerror.New(_opMysqlUserCreate).WithErr(err).
				WithMessage(msg.ErrorMsg500InternalServerError).
				WithKind(richerror.KindStatusInternalServerError)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return models.User{}, richerror.New(_opMysqlUserCreate).WithErr(err).
			WithMessage(msg.ErrorMsg500InternalServerError).
			WithKind(richerror.KindStatusInternalServerError)
	}

	u.ID = types.ID(id) // #nosec G115 // integer overflow conversion int64 -> uint64

	return u, nil
}

func (r DB) IsExistsByMobile(ctx context.Context, mobile string) (bool, error) { //nolint:dupl // nothing
	var exists bool

	query := "select exists(select 1 from users where mobile = ?)"

	stmt, err := r.conn.PrepareStatement( //nolint:sqlclosecheck // nothing
		ctx, uint(mysqlrepo.StatementKeyUserIsExistsByMobile), query,
	)
	if err != nil {
		return false, richerror.New(_opMysqlUserIsExistsByMobile).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	err = stmt.QueryRowContext(ctx, mobile).Scan(&exists)
	// err := r.conn.Conn().QueryRow(query, mobile).Scan(&exists)
	if err != nil {
		return false,
			richerror.New(_opMysqlUserIsExistsByMobile).WithErr(err).
				WithMessage(msg.ErrorMsg500InternalServerError).
				WithKind(richerror.KindStatusInternalServerError)
	}

	if exists {
		return true, nil
	}

	return false, nil
}

func (r DB) IsExistsByEmail(ctx context.Context, email string) (bool, error) { //nolint:dupl // nothing
	var exists bool

	query := "select exists(select 1 from users where email = ?)"

	stmt, err := r.conn.PrepareStatement( //nolint:sqlclosecheck // nothing
		ctx, uint(mysqlrepo.StatementKeyUserIsExistsByEmail), query,
	)
	if err != nil {
		return false, richerror.New(_opMysqlUserIsExistsByEmail).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	err = stmt.QueryRowContext(ctx, email).Scan(&exists)
	// err := r.conn.Conn().QueryRow(query, mobile).Scan(&exists)
	if err != nil {
		return false,
			richerror.New(_opMysqlUserIsExistsByEmail).WithErr(err).
				WithMessage(msg.ErrorMsg500InternalServerError).
				WithKind(richerror.KindStatusInternalServerError)
	}

	if exists {
		return true, nil
	}

	return false, nil
}

func (r DB) GetByMobile(ctx context.Context, mobile string) (models.User, error) {
	query := "SELECT * FROM users WHERE mobile = ?"

	// row := r.conn.Conn().QueryRow(query, mobile)
	stmt, err := r.conn.PrepareStatement( //nolint:sqlclosecheck // nothing
		ctx, uint(mysqlrepo.StatementKeyUserGetByMobile), query,
	)
	if err != nil {
		return models.User{}, richerror.New(_opMysqlUserGetByMobile).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	row := stmt.QueryRowContext(ctx, mobile)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, richerror.New(_opMysqlUserGetByMobile).WithErr(err).
				WithMessage(errMsgDBRecordNotFound).
				WithKind(richerror.KindStatusNotFound)
		}

		return models.User{}, richerror.New(_opMysqlUserGetByMobile).WithErr(err).
			WithMessage(errMsgDBCantScanQueryResult).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return user, nil
}

func (r DB) GetByID(ctx context.Context, id uint64) (models.User, error) {
	query := "SELECT * FROM users WHERE id = ?"

	// row := r.conn.Conn().QueryRow(query, id)
	stmt, err := r.conn.PrepareStatement( //nolint:sqlclosecheck // nothing
		ctx, uint(mysqlrepo.StatementKeyUserGetByID), query,
	)
	if err != nil {
		return models.User{}, richerror.New(_opMysqlUserGetByID).WithErr(err).
			WithMessage(msg.ErrMsgCantPrepareStatement).WithKind(richerror.KindStatusInternalServerError)
	}

	row := stmt.QueryRowContext(ctx, id)
	user, err := scanUser(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, richerror.New(_opMysqlUserGetByID).WithErr(err).
				WithMessage(errMsgDBRecordNotFound).
				WithKind(richerror.KindStatusNotFound)
		}

		return models.User{}, richerror.New(_opMysqlUserGetByID).WithErr(err).
			WithMessage(errMsgDBCantScanQueryResult).
			WithKind(richerror.KindStatusInternalServerError)
	}

	return user, nil
}

func scanUser(scanner Scanner) (models.User, error) {
	var user models.User

	err := scanner.Scan(&user.ID, &user.Name, &user.Mobile, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	// Convert something...

	return user, err
}
