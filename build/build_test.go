package build

import (
	help "github.com/kohirens/stdlib/test"
	"testing"
)

var (
	tmpDir     = help.AbsPath("tmp")
	fixtureDir = "testdata"
)

func TestBuild_Artifacts(t *testing.T) {
	tests := []struct {
		name     string
		bundle   string
		execName string
		wantErr  bool
	}{
		{
			"multi-build",
			"repo-01",
			"ggl",
			false,
		},
	}

	for _, tt := range tests {
		src := help.SetupARepository(tt.bundle, tmpDir, fixtureDir, ps)

		t.Run(tt.name, func(t *testing.T) {
			_, gotErr := Artifacts(src, tt.execName, Platforms)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("Artifacts() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestBuild_buildExecutable(t *testing.T) {
	tests := []struct {
		name     string
		bundle   string
		execName string
		prefix   string
		platform *Platform
		wantErr  bool
		want     *Executable
	}{
		{"build windows",
			"repo-02",
			"ggl",
			"ggl-windows-amd64",
			&Platform{"windows", "amd64"},
			false,
			&Executable{
				Ext:  ".exe",
				Name: "ggl.exe",
			},
		},
		{"build_linux",
			"repo-02",
			"ggl",
			"ggl-linux-amd64",
			&Platform{"linux", "amd64"},
			false,
			&Executable{
				Ext:  "",
				Name: "ggl",
			},
		},
		{"build_mac",
			"repo-02",
			"ggl",
			"ggl-darwin-amd64",
			&Platform{"darwin", "amd64"},
			false,
			&Executable{
				Ext:  "",
				Name: "ggl",
			},
		},
	}

	for _, tt := range tests {
		src := help.SetupARepository(tt.bundle, tmpDir, fixtureDir, ps)

		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := buildExecutable(src, tt.execName, tt.prefix, tt.platform)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("buildExecutable() error = %v, wantErr %v", gotErr, tt.wantErr)
			}

			if got.Ext != tt.want.Ext {
				t.Errorf("got %v, want %v", got.Ext, tt.want.Ext)
			}

			if got.Name != tt.want.Name {
				t.Errorf("got %v, want %v", got.Name, tt.want.Name)
			}
		})
	}
}
