package triangles

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/giantswarm/micrologger"
)

func Test_command(t *testing.T) {
	testCases := []struct {
		name     string
		c1       string
		c2       string
		c3       string
		c4       string
		filename string
		height   int
		width    int
	}{
		{
			name:     "case 0: default pattern",
			filename: "triangles",
		},
		{
			name:     "case 1: tr color pattern",
			filename: "triangles-ca",
			c1:       "#bdc3c7",
			c2:       "#f6b93b",
			c3:       "#f39c12",
			c4:       "#95a5a6",
		},
		{
			name:     "case 2: banner resize",
			filename: "triangles-wide",
			width:    2000,
			height:   400,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			var err error

			var logger micrologger.Logger
			{
				c := micrologger.Config{}

				logger, err = micrologger.New(c)
				if err != nil {
					t.Fatal("expected", nil, "got", err)
				}
			}

			c := Config{
				Logger: logger,
			}

			testCommand, err := New(c)
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}

			args := []string{}
			if tc.filename != "" {
				args = append(args, []string{"--filename", fmt.Sprintf("../../../samples/%s.svg", tc.filename)}...)
			}
			if tc.c1 != "" {
				args = append(args, []string{"--color1", tc.c1}...)
			}
			if tc.c2 != "" {
				args = append(args, []string{"--color2", tc.c2}...)
			}
			if tc.c3 != "" {
				args = append(args, []string{"--color3", tc.c3}...)
			}
			if tc.c4 != "" {
				args = append(args, []string{"--color3", tc.c4}...)
			}
			if tc.height != 0 {
				args = append(args, []string{"--height", strconv.Itoa(tc.height)}...)
			}
			if tc.width != 0 {
				args = append(args, []string{"--width", strconv.Itoa(tc.width)}...)
			}

			testCommand.SetArgs(args)

			err = testCommand.Execute()
			if err != nil {
				t.Fatal("expected", nil, "got", err)
			}
		})
	}
}
