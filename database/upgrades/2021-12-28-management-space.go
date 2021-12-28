package upgrades

import (
	"database/sql"
)

func init() {
	upgrades[32] = upgrade{"Store management space in user table", func(tx *sql.Tx, ctx context) error {
		_, err := tx.Exec(`ALTER TABLE user ADD COLUMN management_space TEXT`)
		return err
	}}
}
