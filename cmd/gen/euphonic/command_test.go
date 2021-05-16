package euphonic

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
		c5       string
		c6       string
		filename string
		height   int
		width    int
	}{
		{
			name:     "case 0: default pattern",
			filename: "euphonic",
			width:    500,
			height:   500,
		},
		{
			name:     "case 1: se color pattern",
			filename: "euphonic-cv",
			c1:       "#2980b9",
			c2:       "#2c3e50",
			c3:       "#e58e26",
			c4:       "#d35400",
			c5:       "#f6b93b",
			c6:       "#0c2461",
		},
		{
			name:     "case 2: banner resize",
			filename: "euphonic-wide",
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
				args = append(args, []string{"--color4", tc.c4}...)
			}
			if tc.c3 != "" {
				args = append(args, []string{"--color5", tc.c5}...)
			}
			if tc.c4 != "" {
				args = append(args, []string{"--color6", tc.c6}...)
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
