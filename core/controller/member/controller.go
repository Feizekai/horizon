package member

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	memberservice "github.com/horizoncd/horizon/pkg/member/service"
	"github.com/horizoncd/horizon/pkg/param"
	"github.com/horizoncd/horizon/pkg/util/errors"
)

type Controller interface {
	// CreateMember create a member of the group
	CreateMember(ctx context.Context, postMember *PostMember) (*Member, error)
	// UpdateMember update a member of the group
	UpdateMember(ctx context.Context, id uint, role string) (*Member, error)
	// RemoveMember leave group or remove a member of the group
	RemoveMember(ctx context.Context, id uint) error
	// ListMember list all the member of the group (and all the member from parent group)
	ListMember(ctx context.Context, resourceType string, resourceID uint) ([]Member, error)
	// GetMemberOfResource get the member of the group by user info in ctx
	GetMemberOfResource(ctx context.Context, resourceType string, resourceID uint) (*Member, error)
}

// NewController initializes a new group controller
func NewController(param *param.Param) Controller {
	return &controller{
		memberService: param.MemberService,
		convertHelper: New(param),
	}
}

type controller struct {
	memberService memberservice.Service
	convertHelper ConvertMemberHelp
}

func (c *controller) CreateMember(ctx context.Context, postMember *PostMember) (*Member, error) {
	const op = "group *controller: create group member"

	member, err := c.memberService.CreateMember(ctx, CovertPostMember(postMember))
	if err != nil {
		switch err {
		case memberservice.ErrMemberExist:
			return nil, errors.E(op, http.StatusBadRequest, err.Error())
		case memberservice.ErrGrantHighRole:
			return nil, errors.E(op, http.StatusBadRequest, err.Error())
		case memberservice.ErrNotPermitted:
			return nil, errors.E(op, http.StatusBadRequest, err.Error())
		default:
			return nil, errors.E(op, http.StatusInternalServerError, err.Error())
		}
	}
	retMember, err := c.convertHelper.ConvertMember(ctx, member)
	if err != nil {
		return nil, errors.E(op, http.StatusInternalServerError, err.Error())
	}
	return retMember, nil
}

func (c *controller) UpdateMember(ctx context.Context, id uint, role string) (*Member, error) {
	const op = "group *controller: update group member"

	member, err := c.memberService.UpdateMember(ctx, id, role)
	if err != nil {
		switch err {
		case memberservice.ErrMemberExist:
			return nil, errors.E(op, http.StatusBadRequest, err.Error())
		case memberservice.ErrGrantHighRole:
			fallthrough
		case memberservice.ErrNotPermitted:
			return nil, errors.E(op, http.StatusForbidden, err.Error())
		default:
			return nil, errors.E(op, http.StatusInternalServerError, err.Error())
		}
	}
	retMember, err := c.convertHelper.ConvertMember(ctx, member)
	if err != nil {
		return nil, errors.E(op, http.StatusInternalServerError, err.Error())
	}
	return retMember, nil
}

func (c *controller) RemoveMember(ctx context.Context, id uint) error {
	const op = "group controller: remove group member"
	err := c.memberService.RemoveMember(ctx, id)
	if err != nil {
		switch err {
		case memberservice.ErrMemberNotExist:
			return errors.E(op, http.StatusNotFound, err.Error())
		case memberservice.ErrNotPermitted:
			fallthrough
		case memberservice.ErrRemoveHighRole:
			return errors.E(op, http.StatusForbidden, err.Error())
		}
	}
	return nil
}

func (c *controller) ListMember(ctx context.Context, resourceType string, id uint) ([]Member, error) {
	op := errors.Op(fmt.Sprintf("group *controller: list %s member", resourceType))
	members, err := c.memberService.ListMember(ctx, resourceType, id)
	if err != nil {
		return nil, errors.E(op, http.StatusInternalServerError, err.Error())
	}
	if members == nil || len(members) < 1 {
		return nil, nil
	}
	retMembers, err := c.convertHelper.ConvertMembers(ctx, members)
	if err != nil {
		return nil, errors.E(op, http.StatusInternalServerError, err.Error())
	}
	return retMembers, nil
}

func (c *controller) GetMemberOfResource(ctx context.Context, resourceType string, resourceID uint) (*Member, error) {
	op := errors.Op(fmt.Sprintf("group *controller: get %s member", resourceType))
	strID := strconv.FormatUint(uint64(resourceID), 10)
	member, err := c.memberService.GetMemberOfResource(ctx, resourceType, strID)
	if err != nil {
		return nil, errors.E(op, http.StatusInternalServerError, err.Error())
	}
	if member == nil {
		return nil, nil
	}
	retMember, err := c.convertHelper.ConvertMember(ctx, member)
	if err != nil {
		return nil, errors.E(op, http.StatusInternalServerError, err.Error())
	}
	return retMember, nil
}
