// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stackrox/rox/central/metrics"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac/resources"
	"github.com/stackrox/rox/pkg/search"
	pgSearch "github.com/stackrox/rox/pkg/search/postgres"
	"gorm.io/gorm"
)

const (
	baseTable = "compliance_operator_benchmark_v2"
	storeName = "ComplianceOperatorBenchmarkV2"
)

var (
	log            = logging.LoggerForModule()
	schema         = pkgSchema.ComplianceOperatorBenchmarkV2Schema
	targetResource = resources.Compliance
)

type storeType = storage.ComplianceOperatorBenchmarkV2

// Store is the interface to interact with the storage for storage.ComplianceOperatorBenchmarkV2
type Store interface {
	Upsert(ctx context.Context, obj *storeType) error
	UpsertMany(ctx context.Context, objs []*storeType) error
	Delete(ctx context.Context, id string) error
	DeleteByQuery(ctx context.Context, q *v1.Query) ([]string, error)
	DeleteMany(ctx context.Context, identifiers []string) error

	Count(ctx context.Context, q *v1.Query) (int, error)
	Exists(ctx context.Context, id string) (bool, error)
	Search(ctx context.Context, q *v1.Query) ([]search.Result, error)

	Get(ctx context.Context, id string) (*storeType, bool, error)
	GetByQuery(ctx context.Context, query *v1.Query) ([]*storeType, error)
	GetMany(ctx context.Context, identifiers []string) ([]*storeType, []int, error)
	GetIDs(ctx context.Context) ([]string, error)

	Walk(ctx context.Context, fn func(obj *storeType) error) error
	WalkByQuery(ctx context.Context, query *v1.Query, fn func(obj *storeType) error) error
}

// New returns a new Store instance using the provided sql instance.
func New(db postgres.DB) Store {
	return pgSearch.NewGenericStore[storeType, *storeType](
		db,
		schema,
		pkGetter,
		insertIntoComplianceOperatorBenchmarkV2,
		copyFromComplianceOperatorBenchmarkV2,
		metricsSetAcquireDBConnDuration,
		metricsSetPostgresOperationDurationTime,
		pgSearch.GloballyScopedUpsertChecker[storeType, *storeType](targetResource),
		targetResource,
	)
}

// region Helper functions

func pkGetter(obj *storeType) string {
	return obj.GetId()
}

func metricsSetPostgresOperationDurationTime(start time.Time, op ops.Op) {
	metrics.SetPostgresOperationDurationTime(start, op, storeName)
}

func metricsSetAcquireDBConnDuration(start time.Time, op ops.Op) {
	metrics.SetAcquireDBConnDuration(start, op, storeName)
}

func insertIntoComplianceOperatorBenchmarkV2(batch *pgx.Batch, obj *storage.ComplianceOperatorBenchmarkV2) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(obj.GetId()),
		obj.GetName(),
		obj.GetShortName(),
		serialized,
	}

	finalStr := "INSERT INTO compliance_operator_benchmark_v2 (Id, Name, ShortName, serialized) VALUES($1, $2, $3, $4) ON CONFLICT(Id) DO UPDATE SET Id = EXCLUDED.Id, Name = EXCLUDED.Name, ShortName = EXCLUDED.ShortName, serialized = EXCLUDED.serialized"
	batch.Queue(finalStr, values...)

	var query string

	for childIndex, child := range obj.GetProfiles() {
		if err := insertIntoComplianceOperatorBenchmarkV2Profiles(batch, child, obj.GetId(), childIndex); err != nil {
			return err
		}
	}

	query = "delete from compliance_operator_benchmark_v2_profiles where compliance_operator_benchmark_v2_Id = $1 AND idx >= $2"
	batch.Queue(query, pgutils.NilOrUUID(obj.GetId()), len(obj.GetProfiles()))
	return nil
}

func insertIntoComplianceOperatorBenchmarkV2Profiles(batch *pgx.Batch, obj *storage.ComplianceOperatorBenchmarkV2_Profile, complianceOperatorBenchmarkV2ID string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		pgutils.NilOrUUID(complianceOperatorBenchmarkV2ID),
		idx,
		obj.GetProfileName(),
		obj.GetProfileVersion(),
	}

	finalStr := "INSERT INTO compliance_operator_benchmark_v2_profiles (compliance_operator_benchmark_v2_Id, idx, ProfileName, ProfileVersion) VALUES($1, $2, $3, $4) ON CONFLICT(compliance_operator_benchmark_v2_Id, idx) DO UPDATE SET compliance_operator_benchmark_v2_Id = EXCLUDED.compliance_operator_benchmark_v2_Id, idx = EXCLUDED.idx, ProfileName = EXCLUDED.ProfileName, ProfileVersion = EXCLUDED.ProfileVersion"
	batch.Queue(finalStr, values...)

	return nil
}

func copyFromComplianceOperatorBenchmarkV2(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, objs ...*storage.ComplianceOperatorBenchmarkV2) error {
	batchSize := pgSearch.MaxBatchSize
	if len(objs) < batchSize {
		batchSize = len(objs)
	}
	inputRows := make([][]interface{}, 0, batchSize)

	// This is a copy so first we must delete the rows and re-add them
	// Which is essentially the desired behaviour of an upsert.
	deletes := make([]string, 0, batchSize)

	copyCols := []string{
		"id",
		"name",
		"shortname",
		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{
			pgutils.NilOrUUID(obj.GetId()),
			obj.GetName(),
			obj.GetShortName(),
			serialized,
		})

		// Add the ID to be deleted.
		deletes = append(deletes, obj.GetId())

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if err := s.DeleteMany(ctx, deletes); err != nil {
				return err
			}
			// clear the inserts and vals for the next batch
			deletes = deletes[:0]

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"compliance_operator_benchmark_v2"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for idx, obj := range objs {
		_ = idx // idx may or may not be used depending on how nested we are, so avoid compile-time errors.

		if err := copyFromComplianceOperatorBenchmarkV2Profiles(ctx, s, tx, obj.GetId(), obj.GetProfiles()...); err != nil {
			return err
		}
	}

	return nil
}

func copyFromComplianceOperatorBenchmarkV2Profiles(ctx context.Context, s pgSearch.Deleter, tx *postgres.Tx, complianceOperatorBenchmarkV2ID string, objs ...*storage.ComplianceOperatorBenchmarkV2_Profile) error {
	batchSize := pgSearch.MaxBatchSize
	if len(objs) < batchSize {
		batchSize = len(objs)
	}
	inputRows := make([][]interface{}, 0, batchSize)

	copyCols := []string{
		"compliance_operator_benchmark_v2_id",
		"idx",
		"profilename",
		"profileversion",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj "+
			"in the loop is not used as it only consists of the parent ID and the index.  Putting this here as a stop gap "+
			"to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{
			pgutils.NilOrUUID(complianceOperatorBenchmarkV2ID),
			idx,
			obj.GetProfileName(),
			obj.GetProfileVersion(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			if _, err := tx.CopyFrom(ctx, pgx.Identifier{"compliance_operator_benchmark_v2_profiles"}, copyCols, pgx.CopyFromRows(inputRows)); err != nil {
				return err
			}
			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return nil
}

// endregion Helper functions

// region Used for testing

// CreateTableAndNewStore returns a new Store instance for testing.
func CreateTableAndNewStore(ctx context.Context, db postgres.DB, gormDB *gorm.DB) Store {
	pkgSchema.ApplySchemaForTable(ctx, gormDB, baseTable)
	return New(db)
}

// Destroy drops the tables associated with the target object type.
func Destroy(ctx context.Context, db postgres.DB) {
	dropTableComplianceOperatorBenchmarkV2(ctx, db)
}

func dropTableComplianceOperatorBenchmarkV2(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS compliance_operator_benchmark_v2 CASCADE")
	dropTableComplianceOperatorBenchmarkV2Profiles(ctx, db)

}

func dropTableComplianceOperatorBenchmarkV2Profiles(ctx context.Context, db postgres.DB) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS compliance_operator_benchmark_v2_profiles CASCADE")

}

// endregion Used for testing
