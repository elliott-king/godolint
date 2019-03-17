package rules

import (
	"testing"
)

func TestDL3025Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM busybox
ENTRYPOINT s3cmd
`,
			file: "DL3025Check_Dockerfile",
			expectedRst: []string{
				"DL3025Check_Dockerfile:2 DL3025 Use arguments JSON notation for CMD and ENTRYPOINT arguments\n",
			},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM busybox
CMD my-service server
`,
			file: "DL3025Check_Dockerfile_2",
			expectedRst: []string{
				"DL3025Check_Dockerfile_2:2 DL3025 Use arguments JSON notation for CMD and ENTRYPOINT arguments\n",
			},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3025Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3025Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3025Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3025Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}