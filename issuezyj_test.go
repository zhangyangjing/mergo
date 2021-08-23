package mergo_test

import (
	"fmt"
	"github.com/zhangyangjing/mergo"
	"testing"
)

type CommonAppConfig struct {
	Backup *bool
}

type BackupConfig struct {
	DataRetentionMaxDays *uint32
}

type Config struct {
	Drive    *CommonAppConfig
	Mail     *CommonAppConfig
	Calendar *CommonAppConfig
	Contacts *CommonAppConfig
	Site     *CommonAppConfig
	Backup   *BackupConfig
}

func TestIssueZyj(t *testing.T) {
	T := true
	F := false

	days := uint32(365)
	days2 := uint32(3)

	d1 := Config{
		Drive: &CommonAppConfig{
			Backup: &T,
		},
		Mail: &CommonAppConfig{
			Backup: &T,
		},
		Backup: &BackupConfig{
			DataRetentionMaxDays: &days,
		},
	}

	d2 := Config{
		Mail: &CommonAppConfig{
			Backup: &F,
		},
		Site: &CommonAppConfig{
			Backup: &F,
		},
		Backup: &BackupConfig{
			DataRetentionMaxDays: &days2,
		},
	}

	err := mergo.Merge(&d1, d2, mergo.WithOverride)
	if err != nil {
		t.Errorf("Error while merging %s", err)
	}

	fmt.Printf("result: drive: %t, mail: %t, site: %t, days: %d\n", *d1.Drive.Backup, *d1.Mail.Backup, *d1.Site.Backup, *d1.Backup.DataRetentionMaxDays)
}
