package engine

import (
	"strings"
	"testing"
)

func TestBuildDropColumnUpSQLContainsDependencySections(t *testing.T) {
	t.Parallel()

	plan := dropColumnDependencyPlan{
		Constraints: []dropConstraintTarget{
			{Table: "dbo.TableA", Name: "CK_TableA__amount_nonneg"},
		},
		Indexes: []dropIndexTarget{
			{Table: "dbo.TableA", Name: "IX_TableA__amount"},
		},
	}

	sql := strings.ToUpper(buildDropColumnUpSQL("dbo.TableA", []string{"amount"}, plan))
	required := []string{
		"ALTER TABLE DBO.TABLEA DROP CONSTRAINT [CK_TABLEA__AMOUNT_NONNEG];",
		"FROM SYS.INDEXES",
		"DROP INDEX IF EXISTS [IX_TABLEA__AMOUNT] ON DBO.TABLEA;",
		"ALTER TABLE DBO.TABLEA DROP COLUMN [AMOUNT]",
	}

	for _, token := range required {
		if !strings.Contains(sql, strings.ToUpper(token)) {
			t.Fatalf("generated SQL missing token: %s", token)
		}
	}

	forbidden := []string{
		"SP_EXECUTESQL",
	}
	for _, token := range forbidden {
		if strings.Contains(sql, token) {
			t.Fatalf("generated SQL must not contain forbidden runtime discovery token: %s", token)
		}
	}
}
