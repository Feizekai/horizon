package group

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"g.hz.netease.com/horizon/lib/q"
	"g.hz.netease.com/horizon/pkg/group/dao"
	"g.hz.netease.com/horizon/pkg/group/models"
)

var (
	// Mgr is the global group manager
	Mgr = New()

	ErrHasChildren = errors.New("children exist, cannot be deleted")
)

type Manager interface {
	Create(ctx context.Context, group *models.Group) (uint, error)
	Delete(ctx context.Context, id uint) (int64, error)
	GetByID(ctx context.Context, id uint) (*models.Group, error)
	GetByIDs(ctx context.Context, ids []uint) ([]*models.Group, error)
	GetByIDsOrderByIDDesc(ctx context.Context, ids []uint) ([]*models.Group, error)
	GetByPaths(ctx context.Context, paths []string) ([]*models.Group, error)
	GetByNameFuzzily(ctx context.Context, name string) ([]*models.Group, error)
	UpdateBasic(ctx context.Context, group *models.Group) error
	ListWithoutPage(ctx context.Context, query *q.Query) ([]*models.Group, error)
	List(ctx context.Context, query *q.Query) ([]*models.Group, int64, error)
	Transfer(ctx context.Context, id, newParentID uint) error
}

type manager struct {
	dao dao.DAO
}

func (m manager) Transfer(ctx context.Context, id, newParentID uint) error {
	return m.dao.Transfer(ctx, id, newParentID)
}

func (m manager) GetByIDsOrderByIDDesc(ctx context.Context, ids []uint) ([]*models.Group, error) {
	return m.dao.GetByIDsOrderByIDDesc(ctx, ids)
}

func (m manager) GetByPaths(ctx context.Context, paths []string) ([]*models.Group, error) {
	return m.dao.GetByPaths(ctx, paths)
}

func (m manager) GetByIDs(ctx context.Context, ids []uint) ([]*models.Group, error) {
	return m.dao.GetByIDs(ctx, ids)
}

func (m manager) GetByNameFuzzily(ctx context.Context, name string) ([]*models.Group, error) {
	return m.dao.GetByNameFuzzily(ctx, name)
}

func (m manager) Create(ctx context.Context, group *models.Group) (uint, error) {
	var pGroup *models.Group
	var err error
	// check if parent exists
	if group.ParentID > 0 {
		pGroup, err = m.dao.GetByID(ctx, group.ParentID)
		if err != nil {
			return 0, err
		}
	}

	// check if there's a record with the same parentID and name
	err = m.dao.CheckNameUnique(ctx, group)
	if err != nil {
		return 0, err
	}
	// check if there's a record with the same parentID and path
	err = m.dao.CheckPathUnique(ctx, group)
	if err != nil {
		return 0, err
	}

	id, err := m.dao.Create(ctx, group)
	if err != nil {
		return 0, err
	}

	// update traversal_ids, like 1; 1,2,3
	var traversalIDs string
	if pGroup == nil {
		traversalIDs = strconv.Itoa(int(id))
	} else {
		traversalIDs = fmt.Sprintf("%s,%d", pGroup.TraversalIDs, id)
	}
	err = m.dao.UpdateTraversalIDs(ctx, id, traversalIDs)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m manager) Delete(ctx context.Context, id uint) (int64, error) {
	count, err := m.dao.CountByParentID(ctx, id)
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, ErrHasChildren
	}
	// todo check application children exist

	return m.dao.Delete(ctx, id)
}

func (m manager) GetByID(ctx context.Context, id uint) (*models.Group, error) {
	return m.dao.GetByID(ctx, id)
}

func (m manager) UpdateBasic(ctx context.Context, group *models.Group) error {
	// check record exist
	_, err := m.dao.GetByID(ctx, group.ID)
	if err != nil {
		return err
	}

	// check if there's record with the same parentID and name
	err = m.dao.CheckNameUnique(ctx, group)
	if err != nil {
		return err
	}
	// check if there's a record with the same parentID and path
	err = m.dao.CheckPathUnique(ctx, group)
	if err != nil {
		return err
	}

	return m.dao.UpdateBasic(ctx, group)
}

func (m manager) ListWithoutPage(ctx context.Context, query *q.Query) ([]*models.Group, error) {
	return m.dao.ListWithoutPage(ctx, query)
}

func (m manager) List(ctx context.Context, query *q.Query) ([]*models.Group, int64, error) {
	return m.dao.List(ctx, query)
}

func New() Manager {
	return &manager{dao: dao.New()}
}
