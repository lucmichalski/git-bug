package cache

import (
	"github.com/MichaelMure/git-bug/identity"
)

var _ identity.Interface = &IdentityCache{}

// IdentityCache is a wrapper around an Identity for caching.
type IdentityCache struct {
	*identity.Identity
	repoCache *RepoCache
}

func NewIdentityCache(repoCache *RepoCache, id *identity.Identity) *IdentityCache {
	return &IdentityCache{
		Identity:  id,
		repoCache: repoCache,
	}
}

func (i *IdentityCache) notifyUpdated() error {
	return i.repoCache.identityUpdated(i.Identity.Id())
}

func (i *IdentityCache) Mutate(f func(identity.Mutator) identity.Mutator) error {
	i.Identity.Mutate(f)
	return i.notifyUpdated()
}

func (i *IdentityCache) Commit() error {
	err := i.Identity.Commit(i.repoCache.repo)
	if err != nil {
		return err
	}
	return i.notifyUpdated()
}

func (i *IdentityCache) CommitAsNeeded() error {
	err := i.Identity.CommitAsNeeded(i.repoCache.repo)
	if err != nil {
		return err
	}
	return i.notifyUpdated()
}
