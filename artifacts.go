package conveyor

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

// Artifact represents an image that was successfully created from a build.
type Artifact struct {
	// Unique identifier for this artifact.
	ID string `db:"id"`
	// Autogenerated sequence id.
	Seq int64 `db:"seq"`
	// The build that this artifact was a result of.
	BuildID string `db:"build_id"`
	// The name of the image that was produced.
	Image string `db:"image"`
	// The repository that this artifact relates to.
	Repository string `db:"repository"`
	// The sha that this artifact relates to.
	Sha string `db:"sha"`
}

// artifactsCreate creates a new artifact linked to the build.
func artifactsCreate(tx *sqlx.Tx, a *Artifact) error {
	const createArtifactSQL = `INSERT INTO artifacts (build_id, image, repository, sha)
(
	SELECT :build_id, :image, repository, sha
	FROM builds
	WHERE id = :build_id
)
RETURNING id, repository, sha`
	return insert(tx, createArtifactSQL, a, &a.ID, &a.Repository, &a.Sha)
}

// artifactsFindByID finds an artifact by ID.
func artifactsFindByID(tx *sqlx.Tx, artifactID string) (*Artifact, error) {
	var sql = `SELECT * FROM artifacts WHERE id = ? LIMIT 1`
	var a Artifact
	err := tx.Get(&a, tx.Rebind(sql), artifactID)
	return &a, err
}

// artifactsFindByRepoSha finds an artifact by image.
func artifactsFindByRepoSha(tx *sqlx.Tx, repoSha string) (*Artifact, error) {
	parts := strings.Split(repoSha, "@")
	var sql = `SELECT * FROM artifacts
WHERE repository = ?
AND sha = ?
ORDER BY seq desc
LIMIT 1`
	var a Artifact
	err := tx.Get(&a, tx.Rebind(sql), parts[0], parts[1])
	return &a, err
}
