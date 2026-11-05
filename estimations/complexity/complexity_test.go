package complexity

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClassifyOS(t *testing.T) {
	cases := []struct {
		name, osName string
		want         Score
	}{
		{"RHEL 4 is Medium", "Red Hat Enterprise Linux 4 (64-bit)", 2},
		{"RHEL 5 is Medium", "Red Hat Enterprise Linux 5 (64-bit)", 2},
		{"RHEL 6 is score 1", "Red Hat Enterprise Linux 6 (64-bit)", 1},
		{"RHEL 8 is score 1", "Red Hat Enterprise Linux 8 (64-bit)", 1},
		{"RHEL 9 is score 1", "Red Hat Enterprise Linux 9 (64-bit)", 1},
		{"RHEL 10 is score 1", "Red Hat Enterprise Linux 10 (64-bit)", 1},
		{"Red Hat Fedora is score 1", "Red Hat Fedora (64-bit)", 1},
		{"Rocky Linux 9 is score 1", "Rocky Linux 9 (64-bit)", 1},
		{"CentOS 6 is score 1", "CentOS 6 (64-bit)", 1},
		{"CentOS 7 is score 1", "CentOS 7 (64-bit)", 1},
		{"Oracle Linux 6 is score 1", "Oracle Linux 6 (64-bit)", 1},
		{"Oracle Linux 8 is score 1", "Oracle Linux 8 (64-bit)", 1},
		{"CentOS 4 is Medium", "CentOS 4/5 (32-bit)", 2},
		{"CentOS 5 is Medium", "CentOS 5 (64-bit)", 2},
		{"Oracle Linux 4 is Medium", "Oracle Linux 4 (64-bit)", 2},
		{"Oracle Linux 5 is Medium", "Oracle Linux 5 (64-bit)", 2},
		{"SUSE 12 is Medium", "SUSE Linux Enterprise 12 (64-bit)", 2},
		{"SUSE 15 is Medium", "SUSE Linux Enterprise 15 (64-bit)", 2},
		{"Windows Server 2019 is Medium", "Microsoft Windows Server 2019 (64-bit)", 2},
		{"Windows 10 is Medium", "Microsoft Windows 10 (64-bit)", 2},
		{"Windows 11 is Medium", "Microsoft Windows 11 (64-bit)", 2},
		{"SUSE 11 is Hard", "SUSE Linux Enterprise 11 (64-bit)", 3},
		{"Ubuntu Linux is Hard", "Ubuntu Linux (64-bit)", 3},
		{"Ubuntu Linux 22.04 is Hard", "Ubuntu Linux 22.04 (64-bit)", 3},
		{"Debian GNU/Linux 11 is Hard", "Debian GNU/Linux 11 (64-bit)", 3},
		{"AlmaLinux 9 is Hard", "AlmaLinux 9 (64-bit)", 3},
		{"FreeBSD is Hard", "FreeBSD (64-bit)", 3},
		{"VMware Photon OS is Hard", "VMware Photon OS (64-bit)", 3},
		{"Amazon Linux 2 is Hard", "Amazon Linux 2 (64-bit)", 3},
		{"CoreOS Linux is Hard", "CoreOS Linux (64-bit)", 3},
		{"Apple macOS is Hard", "Apple macOS (64-bit)", 3},
		{"Oracle Solaris 11 is Hard", "Oracle Solaris 11 (64-bit)", 3},
		{"Windows Server 2008 is Hard", "Microsoft Windows Server 2008 (64-bit)", 3},
		{"Windows Server 2012 is Hard", "Microsoft Windows Server 2012 (64-bit)", 3},
		{"Windows 7 is Hard", "Microsoft Windows 7 (64-bit)", 3},
		{"Microsoft SQL exact", "Microsoft SQL", 4},
		{"Microsoft SQL full name", "Microsoft SQL Server 2019", 4},
		{"empty string", "", 0},
		{"generic other", "Other (64-bit)", 0},
		{"Other Linux", "Other Linux (64-bit)", 0},
		{"lowercase RHEL 9", "red hat enterprise linux 9 (64-bit)", 1},
		{"uppercase CENTOS 7", "CENTOS 7 (64-BIT)", 1},
		{"mixed case Windows 2019", "microsoft WINDOWS SERVER 2019 (64-bit)", 2},
		{"lowercase SUSE 15", "suse linux enterprise 15 (64-bit)", 2},
		{"lowercase Oracle Linux 8", "oracle linux 8 (64-bit)", 1},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ClassifyOS(tc.osName); got != tc.want {
				t.Errorf("ClassifyOS(%q) = %d, want %d", tc.osName, got, tc.want)
			}
		})
	}
}

// Real-world VMware OS strings collected from an RVTools output.
func TestClassifyOSInventoryNames(t *testing.T) {
	cases := []struct {
		name, osName string
		want         Score
	}{
		{"CentOS 4/5 32-bit", "CentOS 4/5 (32-bit)", 2},
		{"CentOS 4/5/6 64-bit", "CentOS 4/5/6 (64-bit)", 2},
		{"CentOS 6 64-bit", "CentOS 6 (64-bit)", 1},
		{"CentOS 7 64-bit", "CentOS 7 (64-bit)", 1},
		{"Debian GNU/Linux 11 64-bit", "Debian GNU/Linux 11 (64-bit)", 3},
		{"FreeBSD 64-bit", "FreeBSD (64-bit)", 3},
		{"Appgate SDP Linux", "Linux 6.5.0-45-generic Appgate SDP 1.0 Appgate SDP 1.0", 0},
		{"Windows 10 64-bit", "Microsoft Windows 10 (64-bit)", 2},
		{"Windows Server 2008 64-bit", "Microsoft Windows Server 2008 (64-bit)", 3},
		{"Windows Server 2008 R2 64-bit", "Microsoft Windows Server 2008 R2 (64-bit)", 3},
		{"Windows Server 2012 64-bit", "Microsoft Windows Server 2012 (64-bit)", 3},
		{"Windows Server 2016 64-bit", "Microsoft Windows Server 2016 (64-bit)", 2},
		{"Windows Server 2016 or later 64-bit", "Microsoft Windows Server 2016 or later (64-bit)", 2},
		{"Windows Server 2019 64-bit", "Microsoft Windows Server 2019 (64-bit)", 2},
		{"Windows Server 2022 64-bit", "Microsoft Windows Server 2022 (64-bit)", 2},
		{"Other 32-bit", "Other (32-bit)", 0},
		{"Other 64-bit", "Other (64-bit)", 0},
		{"Other 2.6.x Linux 64-bit", "Other 2.6.x Linux (64-bit)", 0},
		{"Other 3.x Linux 64-bit", "Other 3.x Linux (64-bit)", 0},
		{"Other 3.x or later Linux 64-bit", "Other 3.x or later Linux (64-bit)", 0},
		{"Other 4.x or later Linux 64-bit", "Other 4.x or later Linux (64-bit)", 0},
		{"Other 5.x or later Linux 64-bit", "Other 5.x or later Linux (64-bit)", 0},
		{"Other Linux 64-bit", "Other Linux (64-bit)", 0},
		{"RHEL 5 64-bit", "Red Hat Enterprise Linux 5 (64-bit)", 2},
		{"RHEL 6 64-bit", "Red Hat Enterprise Linux 6 (64-bit)", 1},
		{"RHEL 7 64-bit", "Red Hat Enterprise Linux 7 (64-bit)", 1},
		{"RHEL 8 64-bit", "Red Hat Enterprise Linux 8 (64-bit)", 1},
		{"RHEL 9 64-bit", "Red Hat Enterprise Linux 9 (64-bit)", 1},
		{"Red Hat Fedora 64-bit", "Red Hat Fedora (64-bit)", 1},
		{"SUSE Linux Enterprise 11 64-bit", "SUSE Linux Enterprise 11 (64-bit)", 3},
		{"SUSE Linux Enterprise 12 64-bit", "SUSE Linux Enterprise 12 (64-bit)", 2},
		{"Ubuntu Linux 64-bit", "Ubuntu Linux (64-bit)", 3},
		{"VMware Photon OS 64-bit", "VMware Photon OS (64-bit)", 3},
		{"RHEL 10 64-bit", "Red Hat Enterprise Linux 10 (64-bit)", 1},
		{"Oracle Linux 4 64-bit", "Oracle Linux 4 (64-bit)", 2},
		{"Oracle Linux 8 64-bit", "Oracle Linux 8 (64-bit)", 1},
		{"Oracle Solaris 11 64-bit", "Oracle Solaris 11 (64-bit)", 3},
		{"AlmaLinux 8 64-bit", "AlmaLinux 8 (64-bit)", 3},
		{"AlmaLinux 9 64-bit", "AlmaLinux 9 (64-bit)", 3},
		{"Rocky Linux 9 64-bit", "Rocky Linux 9 (64-bit)", 1},
		{"Windows 11 64-bit", "Microsoft Windows 11 (64-bit)", 2},
		{"Windows Server 2025 64-bit", "Microsoft Windows Server 2025 (64-bit)", 2},
		{"Windows XP 32-bit", "Microsoft Windows XP Professional (32-bit)", 3},
		{"Windows 7 64-bit", "Microsoft Windows 7 (64-bit)", 3},
		{"Amazon Linux 2 64-bit", "Amazon Linux 2 (64-bit)", 3},
		{"Windows Server 2022 x64 with Microsoft prefix", "Microsoft Windows Server 2022 x64", 2},
		{"Windows Server 2022 x64", "Windows Server 2022 x64", 2},
		{"Windows Server 2022 without Microsoft prefix", "Windows Server 2022 (64-bit)", 2},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ClassifyOS(tc.osName); got != tc.want {
				t.Errorf("ClassifyOS(%q) = %d, want %d", tc.osName, got, tc.want)
			}
		})
	}
}

func TestScoreDiskTierLabel(t *testing.T) {
	cases := []struct {
		name, label string
		want        Score
	}{
		{"legacy score 1", "Easy (0-10TB)", 1},
		{"legacy score 2", "Medium (10-20TB)", 2},
		{"legacy score 3", "Hard (20-50TB)", 3},
		{"legacy score 4", "White Glove (>50TB)", 4},
		{"new score 1", "0-10TiB", 1},
		{"new score 2", "10-20TiB", 2},
		{"new score 3", "20-50TiB", 3},
		{"new score 4", "50+TiB", 4},
		{"unknown label", "Unknown Tier", 0},
		{"empty label", "", 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ScoreDiskTierLabel(tc.label); got != tc.want {
				t.Errorf("ScoreDiskTierLabel(%q) = %d, want %d", tc.label, got, tc.want)
			}
		})
	}
}

func TestOSDiskEntryFields(t *testing.T) {
	entry := OSDiskEntry{Score: 2, VMCount: 10, TotalSizeTB: 5.5}
	if entry.Score != 2 || entry.VMCount != 10 || entry.TotalSizeTB != 5.5 {
		t.Errorf("OSDiskEntry fields = %+v, want Score=2 VMCount=10 TotalSizeTB=5.5", entry)
	}
}

func TestOSBreakdown(t *testing.T) {
	cases := []struct {
		name    string
		entries []VMOsEntry
		want    []OSDifficultyEntry
	}{
		{
			name: "empty input includes all scores with zero counts",
			want: []OSDifficultyEntry{{Score: 0}, {Score: 1}, {Score: 2}, {Score: 3}, {Score: 4}},
		},
		{
			name: "accumulates VM counts by score",
			entries: []VMOsEntry{
				{Name: "Red Hat Enterprise Linux 8 (64-bit)", Count: 100},
				{Name: "Red Hat Enterprise Linux 9 (64-bit)", Count: 50},
				{Name: "CentOS 7 (64-bit)", Count: 30},
				{Name: "FreeBSD (64-bit)", Count: 10},
			},
			want: []OSDifficultyEntry{{Score: 0}, {Score: 1, VMCount: 180}, {Score: 2}, {Score: 3, VMCount: 10}, {Score: 4}},
		},
		{
			name: "unrecognised names have score zero",
			entries: []VMOsEntry{
				{Name: "Other (64-bit)", Count: 5},
				{Name: "Other Linux (64-bit)", Count: 3},
			},
			want: []OSDifficultyEntry{{Score: 0, VMCount: 8}, {Score: 1}, {Score: 2}, {Score: 3}, {Score: 4}},
		},
		{
			name: "all five tiers",
			entries: []VMOsEntry{
				{Name: "Other (64-bit)", Count: 1},
				{Name: "Red Hat Enterprise Linux 9 (64-bit)", Count: 10},
				{Name: "Microsoft Windows Server 2022 (64-bit)", Count: 5},
				{Name: "Ubuntu Linux (64-bit)", Count: 3},
				{Name: "Microsoft SQL Server 2019", Count: 2},
			},
			want: []OSDifficultyEntry{{Score: 0, VMCount: 1}, {Score: 1, VMCount: 10}, {Score: 2, VMCount: 5}, {Score: 3, VMCount: 3}, {Score: 4, VMCount: 2}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := OSBreakdown(tc.entries); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("OSBreakdown() = %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestDiskSizeRangeRatings(t *testing.T) {
	want := map[string]Score{"0-10TiB": 1, "10-20TiB": 2, "20-50TiB": 3, "50+TiB": 4}
	first := DiskSizeRangeRatings()
	second := DiskSizeRangeRatings()
	if !reflect.DeepEqual(first, want) {
		t.Errorf("DiskSizeRangeRatings() = %v, want %v", first, want)
	}
	first["0-10TiB"] = 99
	first["0-10TB"] = 99
	if !reflect.DeepEqual(second, want) {
		t.Errorf("mutating one result changed another: got %v, want %v", second, want)
	}
	if got := DiskSizeScores["Easy (0-10TB)"]; got != 1 {
		t.Errorf("mutating returned map changed legacy DiskSizeScores: got %d, want 1", got)
	}
}

func TestOSRatings(t *testing.T) {
	cases := []struct {
		name    string
		entries []VMOsEntry
		want    map[string]Score
	}{
		{"empty input", nil, map[string]Score{}},
		{
			name:    "known OS",
			entries: []VMOsEntry{{Name: "Red Hat Enterprise Linux 8 (64-bit)", Count: 10}},
			want:    map[string]Score{"Red Hat Enterprise Linux 8 (64-bit)": 1},
		},
		{
			name:    "unknown OS",
			entries: []VMOsEntry{{Name: "Other (64-bit)", Count: 3}},
			want:    map[string]Score{"Other (64-bit)": 0},
		},
		{
			name: "multiple names",
			entries: []VMOsEntry{
				{Name: "Red Hat Enterprise Linux 9 (64-bit)", Count: 100},
				{Name: "CentOS 7 (64-bit)", Count: 20},
				{Name: "Ubuntu Linux (64-bit)", Count: 5},
				{Name: "Oracle Linux 8 (64-bit)", Count: 2},
				{Name: "FreeBSD (64-bit)", Count: 1},
			},
			want: map[string]Score{
				"Red Hat Enterprise Linux 9 (64-bit)": 1,
				"CentOS 7 (64-bit)":                   1,
				"Ubuntu Linux (64-bit)":               3,
				"Oracle Linux 8 (64-bit)":             1,
				"FreeBSD (64-bit)":                    3,
			},
		},
		{
			name: "VM count does not affect number of entries",
			entries: []VMOsEntry{
				{Name: "Red Hat Enterprise Linux 8 (64-bit)", Count: 500},
				{Name: "Red Hat Enterprise Linux 9 (64-bit)", Count: 300},
			},
			want: map[string]Score{
				"Red Hat Enterprise Linux 8 (64-bit)": 1,
				"Red Hat Enterprise Linux 9 (64-bit)": 1,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := OSRatings(tc.entries); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("OSRatings() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestOSNameBreakdown(t *testing.T) {
	cases := []struct {
		name    string
		entries []VMOsEntry
		want    []OSNameEntry
	}{
		{"empty input returns non-nil slice", nil, []OSNameEntry{}},
		{
			name: "one row per input in score then name order",
			entries: []VMOsEntry{
				{Name: "Red Hat Enterprise Linux 9 (64-bit)", Count: 100},
				{Name: "CentOS 7 (64-bit)", Count: 20},
				{Name: "FreeBSD (64-bit)", Count: 5},
			},
			want: []OSNameEntry{
				{Name: "CentOS 7 (64-bit)", Score: 1, VMCount: 20},
				{Name: "Red Hat Enterprise Linux 9 (64-bit)", Score: 1, VMCount: 100},
				{Name: "FreeBSD (64-bit)", Score: 3, VMCount: 5},
			},
		},
		{
			name:    "unknown OS",
			entries: []VMOsEntry{{Name: "Other Linux (64-bit)", Count: 3}},
			want:    []OSNameEntry{{Name: "Other Linux (64-bit)", Score: 0, VMCount: 3}},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := OSNameBreakdown(tc.entries); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("OSNameBreakdown() = %+v, want %+v", got, tc.want)
			}
		})
	}
	t.Run("duplicate names stay separate", func(t *testing.T) {
		entries := []VMOsEntry{
			{Name: "Red Hat Enterprise Linux 8 (64-bit)", Count: 100},
			{Name: "Red Hat Enterprise Linux 8 (64-bit)", Count: 50},
		}
		got := OSNameBreakdown(entries)
		if len(got) != 2 {
			t.Fatalf("OSNameBreakdown() has %d entries, want 2", len(got))
		}
		if got[0].Name != entries[0].Name || got[1].Name != entries[0].Name || got[0].Score != 1 || got[1].Score != 1 {
			t.Errorf("duplicate entries have unexpected names or scores: %+v", got)
		}
		if got[0].VMCount+got[1].VMCount != 150 {
			t.Errorf("duplicate entries have total VMCount %d, want 150", got[0].VMCount+got[1].VMCount)
		}
	})
}

func TestOSNameBreakdownMatchesOSBreakdown(t *testing.T) {
	entries := []VMOsEntry{
		{Name: "Red Hat Enterprise Linux 9 (64-bit)", Count: 100},
		{Name: "Red Hat Enterprise Linux 8 (64-bit)", Count: 50},
		{Name: "CentOS 7 (64-bit)", Count: 20},
		{Name: "Ubuntu Linux (64-bit)", Count: 8},
		{Name: "Oracle Linux 8 (64-bit)", Count: 3},
		{Name: "FreeBSD (64-bit)", Count: 5},
	}
	sumByScore := map[Score]int{}
	for _, entry := range OSNameBreakdown(entries) {
		sumByScore[entry.Score] += entry.VMCount
	}
	for _, entry := range OSBreakdown(entries) {
		if got := sumByScore[entry.Score]; got != entry.VMCount {
			t.Errorf("score %d: OSNameBreakdown count = %d, OSBreakdown count = %d", entry.Score, got, entry.VMCount)
		}
	}
}

func TestCombineComplexity(t *testing.T) {
	want := [5][4]Score{
		{0, 0, 0, 0},
		{1, 1, 3, 3},
		{2, 2, 3, 3},
		{2, 2, 3, 3},
		{4, 4, 4, 4},
	}
	for osScore, row := range want {
		for diskIndex, score := range row {
			diskScore := diskIndex + 1
			t.Run(fmt.Sprintf("os=%d/disk=%d", osScore, diskScore), func(t *testing.T) {
				if got := CombineComplexity(osScore, diskScore); got != score {
					t.Errorf("CombineComplexity(%d, %d) = %d, want %d", osScore, diskScore, got, score)
				}
			})
		}
	}
	if got := CombineComplexity(99, 1); got != 0 {
		t.Errorf("CombineComplexity(99, 1) = %d, want 0", got)
	}
}

func TestComplexityMatrixCoversAllScores(t *testing.T) {
	for _, osScore := range OSScores {
		row, ok := ComplexityMatrix[osScore]
		if !ok {
			t.Errorf("ComplexityMatrix missing OS score %d", osScore)
			continue
		}
		for _, diskScore := range DiskScores {
			matrixScore, ok := row[diskScore]
			if !ok {
				t.Errorf("ComplexityMatrix[%d] missing disk score %d", osScore, diskScore)
				continue
			}
			if got := CombineComplexity(osScore, diskScore); got != matrixScore {
				t.Errorf("CombineComplexity(%d, %d) = %d, matrix has %d", osScore, diskScore, got, matrixScore)
			}
		}
	}
}

func TestDiskBreakdown(t *testing.T) {
	cases := []struct {
		name  string
		tiers []DiskTierInput
		want  []DiskComplexityEntry
	}{
		{
			name: "empty input includes all scores with zero values",
			want: []DiskComplexityEntry{{Score: 1}, {Score: 2}, {Score: 3}, {Score: 4}},
		},
		{
			name: "known labels",
			tiers: []DiskTierInput{
				{Label: "Easy (0-10TB)", VMCount: 360, TotalSizeTB: 42.48},
				{Label: "Hard (20-50TB)", VMCount: 5, TotalSizeTB: 25.1},
			},
			want: []DiskComplexityEntry{{Score: 1, VMCount: 360, TotalSizeTB: 42.48}, {Score: 2}, {Score: 3, VMCount: 5, TotalSizeTB: 25.1}, {Score: 4}},
		},
		{
			name: "unknown label ignored",
			tiers: []DiskTierInput{
				{Label: "Easy (0-10TB)", VMCount: 10, TotalSizeTB: 5.0},
				{Label: "UnknownTier", VMCount: 99, TotalSizeTB: 999.0},
			},
			want: []DiskComplexityEntry{{Score: 1, VMCount: 10, TotalSizeTB: 5.0}, {Score: 2}, {Score: 3}, {Score: 4}},
		},
		{
			name: "all four tiers",
			tiers: []DiskTierInput{
				{Label: "Easy (0-10TB)", VMCount: 100, TotalSizeTB: 5.0},
				{Label: "Medium (10-20TB)", VMCount: 20, TotalSizeTB: 15.0},
				{Label: "Hard (20-50TB)", VMCount: 5, TotalSizeTB: 30.0},
				{Label: "White Glove (>50TB)", VMCount: 1, TotalSizeTB: 75.0},
			},
			want: []DiskComplexityEntry{
				{Score: 1, VMCount: 100, TotalSizeTB: 5.0},
				{Score: 2, VMCount: 20, TotalSizeTB: 15.0},
				{Score: 3, VMCount: 5, TotalSizeTB: 30.0},
				{Score: 4, VMCount: 1, TotalSizeTB: 75.0},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := DiskBreakdown(tc.tiers); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("DiskBreakdown() = %+v, want %+v", got, tc.want)
			}
		})
	}
}
