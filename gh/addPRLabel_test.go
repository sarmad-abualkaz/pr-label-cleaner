package gh

import (
	"context"
    "fmt"
    "net/http"
    "testing"
)

func TestAddLable(t *testing.T){

    type endpoint struct{
        owner    string
        prNumber int
        repo     string        
    }
    
    type args struct{
		label    string
        owner    string
        prNumber int
        repo     string
	}
	
	tests := []struct {
		name       string
		args       args
        endpoint   endpoint
        wantedErr  error
	}{
		{
            name: "should succefully remove label",
            args: args{
                label:    "test",
                owner:    "testowner",
                prNumber: 1,
                repo:     "testrepo",
            },
            endpoint: endpoint{
                owner:    "testowner",
                prNumber: 1,
                repo:     "testrepo",
            },
            wantedErr:  nil,
        },
		{
            name: "should fail to find PR",
            args: args{
                label:    "test",
                owner:    "testowner",
                prNumber: 1,
                repo:     "testrepo",
            },
            endpoint: endpoint{
                owner:    "testowner",
                prNumber: 2,
                repo:     "testrepo",
            },
            wantedErr:  fmt.Errorf("PR not found"),
        },
    }

    for _, tt := range tests {

        // var apiurl string
        
        t.Run(tt.name, func(t *testing.T) {

            client, mux, _, teardown := setup()
            defer teardown()

			ctx := context.Background()

                
            apiurl := fmt.Sprintf("/repos/%s/%s/issues/%d/labels", tt.endpoint.owner, tt.endpoint.repo, tt.endpoint.prNumber)

            mux.HandleFunc(apiurl, func(w http.ResponseWriter, r *http.Request) {
                testMethod(t, r, "POST")
            })
            
            err := AddLable(ctx, client, tt.args.label, tt.args.owner, tt.args.prNumber, tt.args.repo)

            if tt.wantedErr == nil {
                
                if err != nil {
                    t.Errorf("AddLable returned error: %v expect %v", err, tt.wantedErr)
                }

            } else if err.Error() != tt.wantedErr.Error() {
                
                t.Errorf("AddLable returned error: %v expect %v", err, tt.wantedErr)
            
            }

        })
    }
}
